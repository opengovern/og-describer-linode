package describer

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/opengovern/og-describer-linode/pkg/sdk/models"
	"github.com/opengovern/og-describer-linode/provider/model"
	"net/http"
	"net/url"
	"strconv"
	"sync"
)

func ListLinodeInstances(ctx context.Context, handler *LinodeAPIHandler, stream *models.StreamSender) ([]models.Resource, error) {
	var wg sync.WaitGroup
	linodeChan := make(chan models.Resource)
	errorChan := make(chan error, 1) // Buffered channel to capture errors

	go func() {
		defer close(linodeChan)
		defer close(errorChan)
		if err := processLinodeInstances(ctx, handler, linodeChan, &wg); err != nil {
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

func GetLinodeInstance(ctx context.Context, handler *LinodeAPIHandler, resourceID string) (*models.Resource, error) {
	linode, err := processLinodeInstance(ctx, handler, resourceID)
	if err != nil {
		return nil, err
	}
	value := models.Resource{
		ID:          strconv.Itoa(linode.ID),
		Name:        linode.Label,
		Description: linode,
	}
	return &value, nil
}

func processLinodeInstances(ctx context.Context, handler *LinodeAPIHandler, openaiChan chan<- models.Resource, wg *sync.WaitGroup) error {
	var linodeInstances []model.LinodeSingleResponse
	var linodeListResponse model.LinodeListResponse
	var resp *http.Response
	baseURL := "https://api.linode.com/v4/linode/instances"
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

			if e = json.NewDecoder(resp.Body).Decode(&linodeListResponse); e != nil {
				return nil, fmt.Errorf("failed to decode response: %w", e)
			}
			linodeInstances = append(linodeInstances, linodeListResponse.Data...)
			return resp, nil
		}

		err = handler.DoRequest(ctx, req, requestFunc)
		if err != nil {
			return fmt.Errorf("error during request handling: %w", err)
		}

		if linodeListResponse.Page == linodeListResponse.Pages {
			break
		}
		page++
	}

	for _, linode := range linodeInstances {
		wg.Add(1)
		go func(linode model.LinodeSingleResponse) {
			defer wg.Done()
			value := models.Resource{
				ID:   strconv.Itoa(linode.ID),
				Name: linode.Label,
				Description: model.InstanceDescription{
					ID:              linode.ID,
					Created:         linode.Created,
					Updated:         linode.Updated,
					Type:            linode.Type,
					Status:          linode.Status,
					Region:          linode.Region,
					IPv4:            linode.IPv4,
					IPv6:            linode.IPv6,
					Image:           linode.Image,
					Backups:         linode.Backups,
					Hypervisor:      linode.Hypervisor,
					Specs:           linode.Specs,
					Alerts:          linode.Alerts,
					Group:           linode.Group,
					HasUserData:     linode.HasUserData,
					HostUUID:        linode.HostUUID,
					WatchdogEnabled: linode.WatchdogEnabled,
					Tags:            linode.Tags,
					PlacementGroup:  linode.PlacementGroup,
					DiskEncryption:  linode.DiskEncryption,
					LKEClusterID:    linode.LKEClusterID,
					Capabilities:    linode.Capabilities,
				},
			}
			openaiChan <- value
		}(linode)
	}
	return nil
}

func processLinodeInstance(ctx context.Context, handler *LinodeAPIHandler, resourceID string) (*model.InstanceDescription, error) {
	var linode model.InstanceDescription
	var resp *http.Response
	baseURL := "https://api.linode.com/v4/linode/instances/"

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

		if e = json.NewDecoder(resp.Body).Decode(&linode); e != nil {
			return nil, fmt.Errorf("failed to decode response: %w", e)
		}
		return resp, e
	}

	err = handler.DoRequest(ctx, req, requestFunc)
	if err != nil {
		return nil, fmt.Errorf("error during request handling: %w", err)
	}
	return &linode, nil
}
