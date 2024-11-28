package provider

import (
	"errors"
	"github.com/linode/linodego"
	model "github.com/opengovern/og-describer-linode/pkg/sdk/models"
	"github.com/opengovern/og-describer-linode/provider/configs"
	"github.com/opengovern/og-describer-linode/provider/describer"
	"github.com/opengovern/og-util/pkg/describe/enums"
	"golang.org/x/net/context"
	"golang.org/x/oauth2"
	"golang.org/x/time/rate"
	"net/http"
	"time"
)

// DescribeListByLinode A wrapper to pass linode authorization to describer functions
func DescribeListByLinode(describe func(context.Context, *describer.LinodeAPIHandler, *model.StreamSender) ([]model.Resource, error)) model.ResourceDescriber {
	return func(ctx context.Context, cfg configs.IntegrationCredentials, triggerType enums.DescribeTriggerType, additionalParameters map[string]string, stream *model.StreamSender) ([]model.Resource, error) {
		ctx = describer.WithTriggerType(ctx, triggerType)

		// Create linode client using token
		var client linodego.Client
		var err error
		// Check for the token
		if cfg.Token == "" {
			return nil, errors.New("token must be configured")
		}

		tokenSource := oauth2.StaticTokenSource(&oauth2.Token{AccessToken: cfg.Token})

		oauth2Client := &http.Client{
			Transport: &oauth2.Transport{
				Source: tokenSource,
			},
		}

		client = linodego.NewClient(oauth2Client)

		linodeAPIHandler := describer.NewLinodeAPIHandler(client, rate.Every(time.Second/4), 1, 10, 5, 5*time.Minute)

		// Get values from describer
		var values []model.Resource
		result, err := describe(ctx, linodeAPIHandler, stream)
		if err != nil {
			return nil, err
		}
		values = append(values, result...)
		return values, nil
	}
}

// DescribeSingleByLinode A wrapper to pass linode authorization to describer functions
func DescribeSingleByLinode(describe func(context.Context, *describer.LinodeAPIHandler, string) (*model.Resource, error)) model.SingleResourceDescriber {
	return func(ctx context.Context, cfg configs.IntegrationCredentials, triggerType enums.DescribeTriggerType, additionalParameters map[string]string, resourceID string) (*model.Resource, error) {
		ctx = describer.WithTriggerType(ctx, triggerType)

		// Create linode client using token
		var client linodego.Client
		var err error
		// Check for the token
		if cfg.Token == "" {
			return nil, errors.New("token must be configured")
		}

		tokenSource := oauth2.StaticTokenSource(&oauth2.Token{AccessToken: cfg.Token})

		oauth2Client := &http.Client{
			Transport: &oauth2.Transport{
				Source: tokenSource,
			},
		}

		client = linodego.NewClient(oauth2Client)

		linodeAPIHandler := describer.NewLinodeAPIHandler(client, rate.Every(time.Second/4), 1, 10, 5, 5*time.Minute)

		// Get value from describer
		value, err := describe(ctx, linodeAPIHandler, resourceID)
		if err != nil {
			return nil, err
		}
		return value, nil
	}
}
