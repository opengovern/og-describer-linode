package describers

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/opengovern/og-describer-linode/discovery/pkg/models"
	"github.com/opengovern/og-describer-linode/discovery/provider"
	"net/http"
	"net/url"
	"strconv"
	"sync"
)

func ListStackScripts(ctx context.Context, handler *provider.LinodeAPIHandler, stream *models.StreamSender) ([]models.Resource, error) {
	var wg sync.WaitGroup
	linodeChan := make(chan models.Resource)
	errorChan := make(chan error, 1) // Buffered channel to capture errors
	account, err := provider.GetAccount(ctx, handler)
	if err != nil {
		return nil, err
	}

	go func() {
		defer close(linodeChan)
		defer close(errorChan)
		if err := processStackScripts(ctx, handler, account.EUUID, linodeChan, &wg); err != nil {
			errorChan <- err // Send error to the error channel
		}
		wg.Wait()
	}()

	var values []models.Resource
	for {
		select {
		case value, ok := <-linodeChan:
			if !ok {
				return values, nil
			}
			if stream != nil {
				if err := (*stream)(value); err != nil {
					return nil, err
				}
			} else {
				values = append(values, value)
			}
		case err := <-errorChan:
			return nil, err
		}
	}
}

func GetStackScript(ctx context.Context, handler *provider.LinodeAPIHandler, resourceID string) (*models.Resource, error) {
	stackScript, err := processStackScript(ctx, handler, resourceID)
	if err != nil {
		return nil, err
	}
	account, err := provider.GetAccount(ctx, handler)
	if err != nil {
		return nil, err
	}
	value := models.Resource{
		ID:   strconv.Itoa(stackScript.ID),
		Name: stackScript.Label,
		Description: provider.StackScriptDescription{
			ID:                stackScript.ID,
			Label:             stackScript.Label,
			Description:       stackScript.Description,
			Username:          stackScript.Username,
			Ordinal:           stackScript.Ordinal,
			LogoURL:           stackScript.LogoURL,
			Images:            stackScript.Images,
			DeploymentsTotal:  stackScript.DeploymentsTotal,
			DeploymentsActive: stackScript.DeploymentsActive,
			IsPublic:          stackScript.IsPublic,
			Mine:              stackScript.Mine,
			Created:           stackScript.Created,
			Updated:           stackScript.Updated,
			RevNote:           stackScript.RevNote,
			Script:            stackScript.Script,
			UserDefinedFields: stackScript.UserDefinedFields,
			UserGravatarID:    stackScript.UserGravatarID,
			Account:           account.EUUID,
		},
	}
	return &value, nil
}

func processStackScripts(ctx context.Context, handler *provider.LinodeAPIHandler, account string, openaiChan chan<- models.Resource, wg *sync.WaitGroup) error {
	var stackScripts []provider.StackScriptResp
	var stackScriptListResponse provider.StackScriptListResponse
	var resp *http.Response
	baseURL := "https://api.linode.com/v4/linode/stackscripts"
	page := 1

	for {
		params := url.Values{}
		params.Set("page", strconv.Itoa(page))
		params.Set("page_size", "500")
		finalURL := fmt.Sprintf("%s?%s", baseURL, params.Encode())

		req, err := http.NewRequest("GET", finalURL, nil)
		if err != nil {
			return fmt.Errorf("failed to create request: %w", err)
		}

		requestFunc := func(req *http.Request) (*http.Response, error) {
			var e error
			resp, e = handler.Client.Do(req)
			if e != nil {
				return nil, fmt.Errorf("request execution failed: %w", e)
			}
			defer resp.Body.Close()

			if e = json.NewDecoder(resp.Body).Decode(&stackScriptListResponse); e != nil {
				return nil, fmt.Errorf("failed to decode response: %w", e)
			}
			stackScripts = append(stackScripts, stackScriptListResponse.Data...)
			return resp, nil
		}

		err = handler.DoRequest(ctx, req, requestFunc)
		if err != nil {
			return fmt.Errorf("error during request handling: %w", err)
		}

		if stackScriptListResponse.Page == stackScriptListResponse.Pages {
			break
		}
		page++
	}

	for _, stackScript := range stackScripts {
		wg.Add(1)
		go func(stackScript provider.StackScriptResp) {
			defer wg.Done()
			value := models.Resource{
				ID:   strconv.Itoa(stackScript.ID),
				Name: stackScript.Label,
				Description: provider.StackScriptDescription{
					ID:                stackScript.ID,
					Label:             stackScript.Label,
					Description:       stackScript.Description,
					Username:          stackScript.Username,
					Ordinal:           stackScript.Ordinal,
					LogoURL:           stackScript.LogoURL,
					Images:            stackScript.Images,
					DeploymentsTotal:  stackScript.DeploymentsTotal,
					DeploymentsActive: stackScript.DeploymentsActive,
					IsPublic:          stackScript.IsPublic,
					Mine:              stackScript.Mine,
					Created:           stackScript.Created,
					Updated:           stackScript.Updated,
					RevNote:           stackScript.RevNote,
					Script:            stackScript.Script,
					UserDefinedFields: stackScript.UserDefinedFields,
					UserGravatarID:    stackScript.UserGravatarID,
					Account:           account,
				},
			}
			openaiChan <- value
		}(stackScript)
	}
	return nil
}

func processStackScript(ctx context.Context, handler *provider.LinodeAPIHandler, resourceID string) (*provider.StackScriptDescription, error) {
	var stackScript provider.StackScriptDescription
	var resp *http.Response
	baseURL := "https://api.linode.com/v4/linode/stackscripts/"

	finalURL := fmt.Sprintf("%s%s", baseURL, resourceID)
	req, err := http.NewRequest("GET", finalURL, nil)
	if err != nil {
		return nil, err
	}

	requestFunc := func(req *http.Request) (*http.Response, error) {
		var e error
		resp, e = handler.Client.Do(req)
		if e != nil {
			return nil, fmt.Errorf("request execution failed: %w", e)
		}
		defer resp.Body.Close()

		if e = json.NewDecoder(resp.Body).Decode(&stackScript); e != nil {
			return nil, fmt.Errorf("failed to decode response: %w", e)
		}
		return resp, e
	}

	err = handler.DoRequest(ctx, req, requestFunc)
	if err != nil {
		return nil, fmt.Errorf("error during request handling: %w", err)
	}
	return &stackScript, nil
}
