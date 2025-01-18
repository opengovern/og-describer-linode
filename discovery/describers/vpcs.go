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

func ListVPCs(ctx context.Context, handler *provider.LinodeAPIHandler, stream *models.StreamSender) ([]models.Resource, error) {
	var wg sync.WaitGroup
	linodeChan := make(chan models.Resource)
	errorChan := make(chan error, 1) // Buffered channel to capture errors

	go func() {
		defer close(linodeChan)
		defer close(errorChan)
		if err := processVPCs(ctx, handler, linodeChan, &wg); err != nil {
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

func GetVPC(ctx context.Context, handler *provider.LinodeAPIHandler, resourceID string) (*models.Resource, error) {
	vpc, err := processVPC(ctx, handler, resourceID)
	if err != nil {
		return nil, err
	}
	value := models.Resource{
		ID:          strconv.Itoa(vpc.ID),
		Name:        vpc.Label,
		Description: vpc,
	}
	return &value, nil
}

func processVPCs(ctx context.Context, handler *provider.LinodeAPIHandler, openaiChan chan<- models.Resource, wg *sync.WaitGroup) error {
	var vpcs []provider.VPCDescription
	var vpcListResponse provider.VPCListResponse
	var resp *http.Response
	baseURL := "https://api.linode.com/v4/vpcs"
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

			if e = json.NewDecoder(resp.Body).Decode(&vpcListResponse); e != nil {
				return nil, fmt.Errorf("failed to decode response: %w", e)
			}
			vpcs = append(vpcs, vpcListResponse.Data...)
			return resp, nil
		}

		err = handler.DoRequest(ctx, req, requestFunc)
		if err != nil {
			return fmt.Errorf("error during request handling: %w", err)
		}

		if vpcListResponse.Page == vpcListResponse.Pages {
			break
		}
		page++
	}
	for _, vpc := range vpcs {
		wg.Add(1)
		go func(vpc provider.VPCDescription) {
			defer wg.Done()
			value := models.Resource{
				ID:          strconv.Itoa(vpc.ID),
				Name:        vpc.Label,
				Description: vpc,
			}
			openaiChan <- value
		}(vpc)
	}
	return nil
}

func processVPC(ctx context.Context, handler *provider.LinodeAPIHandler, resourceID string) (*provider.VPCDescription, error) {
	var vpc provider.VPCDescription
	var resp *http.Response
	baseURL := "https://api.linode.com/v4/vpcs/"

	finalURL := fmt.Sprintf("%s%s", baseURL, resourceID)
	req, err := http.NewRequest("GET", finalURL, nil)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	requestFunc := func(req *http.Request) (*http.Response, error) {
		var e error
		resp, e = handler.Client.Do(req)
		if e != nil {
			return nil, fmt.Errorf("request execution failed: %w", e)
		}

		if e = json.NewDecoder(resp.Body).Decode(&vpc); e != nil {
			return nil, fmt.Errorf("failed to decode response: %w", e)
		}
		return resp, e
	}

	err = handler.DoRequest(ctx, req, requestFunc)
	if err != nil {
		return nil, fmt.Errorf("error during request handling: %w", err)
	}
	return &vpc, nil
}
