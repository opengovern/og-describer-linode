package provider

import (
	"errors"
	"github.com/opengovern/og-describer-linode/discovery/pkg/models"
	"github.com/opengovern/og-util/pkg/describe/enums"
	"golang.org/x/net/context"
	"golang.org/x/time/rate"
	"time"
)

// DescribeListByLinode A wrapper to pass linode authorization to describer functions
func DescribeListByLinode(describe func(context.Context, *LinodeAPIHandler, *models.StreamSender) ([]models.Resource, error)) models.ResourceDescriber {
	return func(ctx context.Context, cfg models.IntegrationCredentials, triggerType enums.DescribeTriggerType, additionalParameters map[string]string, stream *models.StreamSender) ([]models.Resource, error) {
		ctx = WithTriggerType(ctx, triggerType)

		var err error
		// Check for the token
		if cfg.Token == "" {
			return nil, errors.New("token must be configured")
		}

		linodeAPIHandler := NewLinodeAPIHandler(cfg.Token, rate.Every(time.Minute/200), 1, 10, 5, 5*time.Minute)

		// Get values from describer
		var values []models.Resource
		result, err := describe(ctx, linodeAPIHandler, stream)
		if err != nil {
			return nil, err
		}
		values = append(values, result...)
		return values, nil
	}
}

// DescribeSingleByLinode A wrapper to pass linode authorization to describer functions
func DescribeSingleByLinode(describe func(context.Context, *LinodeAPIHandler, string) (*models.Resource, error)) models.SingleResourceDescriber {
	return func(ctx context.Context, cfg models.IntegrationCredentials, triggerType enums.DescribeTriggerType, additionalParameters map[string]string, resourceID string, stream *models.StreamSender) (*models.Resource, error) {
		ctx = WithTriggerType(ctx, triggerType)

		var err error
		// Check for the token
		if cfg.Token == "" {
			return nil, errors.New("token must be configured")
		}

		linodeAPIHandler := NewLinodeAPIHandler(cfg.Token, rate.Every(time.Minute/200), 1, 10, 5, 5*time.Minute)

		// Get value from describer
		value, err := describe(ctx, linodeAPIHandler, resourceID)
		if err != nil {
			return nil, err
		}
		return value, nil
	}
}
