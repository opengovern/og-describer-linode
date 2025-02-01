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

func ListVolumes(ctx context.Context, handler *provider.LinodeAPIHandler, stream *models.StreamSender) ([]models.Resource, error) {
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
		if err := processVolumes(ctx, handler, account.EUUID, linodeChan, &wg); err != nil {
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

func GetVolume(ctx context.Context, handler *provider.LinodeAPIHandler, resourceID string) (*models.Resource, error) {
	volume, err := processVolume(ctx, handler, resourceID)
	if err != nil {
		return nil, err
	}
	account, err := provider.GetAccount(ctx, handler)
	if err != nil {
		return nil, err
	}
	value := models.Resource{
		ID:   strconv.Itoa(volume.ID),
		Name: volume.Label,
		Description: provider.VolumeDescription{
			ID:             volume.ID,
			Label:          volume.Label,
			Status:         volume.Status,
			Region:         volume.Region,
			Size:           volume.Size,
			LinodeID:       volume.LinodeID,
			FilesystemPath: volume.FilesystemPath,
			Tags:           volume.Tags,
			Created:        volume.Created,
			Updated:        volume.Updated,
			Encryption:     volume.Encryption,
			Account:        account.EUUID,
		},
	}
	return &value, nil
}

func processVolumes(ctx context.Context, handler *provider.LinodeAPIHandler, account string, openaiChan chan<- models.Resource, wg *sync.WaitGroup) error {
	var volumes []provider.VolumeSingleResponse
	var volumeListResponse provider.VolumeListResponse
	var resp *http.Response
	baseURL := "https://api.linode.com/v4/volumes"
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

			if e = json.NewDecoder(resp.Body).Decode(&volumeListResponse); e != nil {
				return nil, fmt.Errorf("failed to decode response: %w", e)
			}
			volumes = append(volumes, volumeListResponse.Data...)
			return resp, nil
		}

		err = handler.DoRequest(ctx, req, requestFunc)
		if err != nil {
			return fmt.Errorf("error during request handling: %w", err)
		}

		if volumeListResponse.Page == volumeListResponse.Pages {
			break
		}
		page++
	}
	for _, volume := range volumes {
		wg.Add(1)
		go func(volume provider.VolumeSingleResponse) {
			defer wg.Done()
			value := models.Resource{
				ID:   strconv.Itoa(volume.ID),
				Name: volume.Label,
				Description: provider.VolumeDescription{
					ID:             volume.ID,
					Label:          volume.Label,
					Status:         volume.Status,
					Region:         volume.Region,
					Size:           volume.Size,
					LinodeID:       volume.LinodeID,
					FilesystemPath: volume.FilesystemPath,
					Tags:           volume.Tags,
					Created:        volume.Created,
					Updated:        volume.Updated,
					Encryption:     volume.Encryption,
					Account:        account,
				},
			}
			openaiChan <- value
		}(volume)
	}
	return nil
}

func processVolume(ctx context.Context, handler *provider.LinodeAPIHandler, resourceID string) (*provider.VolumeDescription, error) {
	var volume provider.VolumeDescription
	var resp *http.Response
	baseURL := "https://api.linode.com/v4/volumes/"

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

		if e = json.NewDecoder(resp.Body).Decode(&volume); e != nil {
			return nil, fmt.Errorf("failed to decode response: %w", e)
		}
		return resp, e
	}

	err = handler.DoRequest(ctx, req, requestFunc)
	if err != nil {
		return nil, fmt.Errorf("error during request handling: %w", err)
	}
	return &volume, nil
}
