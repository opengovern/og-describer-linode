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

func ListLongViewClients(ctx context.Context, handler *provider.LinodeAPIHandler, stream *models.StreamSender) ([]models.Resource, error) {
	var wg sync.WaitGroup
	linodeChan := make(chan models.Resource)
	errorChan := make(chan error, 1) // Buffered channel to capture errors
	accounts, err := ListAccounts(ctx, handler, stream)
	if err != nil {
		return nil, err
	}

	go func() {
		defer close(linodeChan)
		defer close(errorChan)
		if err := processLongViewClients(ctx, handler, accounts[0].ID, linodeChan, &wg); err != nil {
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

func GetLongViewClient(ctx context.Context, handler *provider.LinodeAPIHandler, resourceID string) (*models.Resource, error) {
	client, err := processLongViewClient(ctx, handler, resourceID)
	if err != nil {
		return nil, err
	}
	accounts, err := ListAccounts(ctx, handler, nil)
	if err != nil {
		return nil, err
	}
	value := models.Resource{
		ID:   strconv.Itoa(client.ID),
		Name: client.Label,
		Description: provider.LongViewClientDescription{
			ID:          client.ID,
			APIKey:      client.APIKey,
			Created:     client.Created,
			InstallCode: client.InstallCode,
			Label:       client.Label,
			Updated:     client.Updated,
			Apps: struct {
				Apache any `json:"apache"`
				MySQL  any `json:"mysql"`
				NginX  any `json:"nginx"`
			}{
				Apache: client.Apps.Apache,
				MySQL:  client.Apps.MySQL,
				NginX:  client.Apps.NginX,
			},
			Account: accounts[0].ID,
		},
	}
	return &value, nil
}

func processLongViewClients(ctx context.Context, handler *provider.LinodeAPIHandler, account string, openaiChan chan<- models.Resource, wg *sync.WaitGroup) error {
	var clients []provider.LongViewClientDescription
	var clientListResponse provider.LongViewClientListResponse
	var resp *http.Response
	baseURL := "https://api.linode.com/v4/longview/clients"
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

			if e = json.NewDecoder(resp.Body).Decode(&clientListResponse); e != nil {
				return nil, fmt.Errorf("failed to decode response: %w", e)
			}
			clients = append(clients, clientListResponse.Data...)
			return resp, nil
		}

		err = handler.DoRequest(ctx, req, requestFunc)
		if err != nil {
			return fmt.Errorf("error during request handling: %w", err)
		}

		if clientListResponse.Page == clientListResponse.Pages {
			break
		}
		page++
	}

	for _, client := range clients {
		wg.Add(1)
		go func(client provider.LongViewClientDescription) {
			defer wg.Done()
			value := models.Resource{
				ID:   strconv.Itoa(client.ID),
				Name: client.Label,
				Description: provider.LongViewClientDescription{
					ID:          client.ID,
					APIKey:      client.APIKey,
					Created:     client.Created,
					InstallCode: client.InstallCode,
					Label:       client.Label,
					Updated:     client.Updated,
					Apps: struct {
						Apache any `json:"apache"`
						MySQL  any `json:"mysql"`
						NginX  any `json:"nginx"`
					}{
						Apache: client.Apps.Apache,
						MySQL:  client.Apps.MySQL,
						NginX:  client.Apps.NginX,
					},
					Account: account,
				},
			}
			openaiChan <- value
		}(client)
	}
	return nil
}

func processLongViewClient(ctx context.Context, handler *provider.LinodeAPIHandler, resourceID string) (*provider.LongViewClientDescription, error) {
	var client provider.LongViewClientDescription
	var resp *http.Response
	baseURL := "https://api.linode.com/v4/longview/clients/"

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

		if e = json.NewDecoder(resp.Body).Decode(&client); e != nil {
			return nil, fmt.Errorf("failed to decode response: %w", e)
		}
		return resp, e
	}

	err = handler.DoRequest(ctx, req, requestFunc)
	if err != nil {
		return nil, fmt.Errorf("error during request handling: %w", err)
	}
	return &client, nil
}
