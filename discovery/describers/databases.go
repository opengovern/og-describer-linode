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

func ListDatabases(ctx context.Context, handler *provider.LinodeAPIHandler, stream *models.StreamSender) ([]models.Resource, error) {
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
		if err := processDatabases(ctx, handler, accounts[0].ID, linodeChan, &wg); err != nil {
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

func processDatabases(ctx context.Context, handler *provider.LinodeAPIHandler, account string, openaiChan chan<- models.Resource, wg *sync.WaitGroup) error {
	var databases []provider.DatabaseSingleResponse
	var databaseListResponse provider.DatabaseListResponse
	var resp *http.Response
	baseURL := "https://api.linode.com/v4/databases/instances"
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

			if e = json.NewDecoder(resp.Body).Decode(&databaseListResponse); e != nil {
				return nil, fmt.Errorf("failed to decode response: %w", e)
			}
			databases = append(databases, databaseListResponse.Data...)
			return resp, nil
		}

		err = handler.DoRequest(ctx, req, requestFunc)
		if err != nil {
			return fmt.Errorf("error during request handling: %w", err)
		}

		if databaseListResponse.Page == databaseListResponse.Pages {
			break
		}
		page++
	}

	for _, database := range databases {
		wg.Add(1)
		go func(database provider.DatabaseSingleResponse) {
			defer wg.Done()
			value := models.Resource{
				ID:   strconv.Itoa(database.ID),
				Name: database.Label,
				Description: provider.DatabaseDescription{
					ID:      database.ID,
					Label:   database.Label,
					Region:  database.Region,
					Type:    database.Type,
					Status:  database.Status,
					Created: database.Created,
					Updated: database.Updated,
					Hosts: provider.DatabaseHost{
						Primary:   database.Hosts.Primary,
						Secondary: database.Hosts.Secondary,
					},
					ClusterSize:     database.ClusterSize,
					ReplicationType: database.ReplicationType,
					SSLConnection:   database.SSLConnection,
					Encrypted:       database.Encrypted,
					AllowList:       database.AllowList,
					InstanceURI:     database.InstanceURI,
					Engine:          database.Engine,
					Version:         database.Version,
					Account:         account,
				},
			}
			openaiChan <- value
		}(database)
	}
	return nil
}
