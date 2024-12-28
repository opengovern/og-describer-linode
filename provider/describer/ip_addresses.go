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

func ListIPAddresses(ctx context.Context, handler *LinodeAPIHandler, stream *models.StreamSender) ([]models.Resource, error) {
	var wg sync.WaitGroup
	linodeChan := make(chan models.Resource)
	errorChan := make(chan error, 1) // Buffered channel to capture errors

	go func() {
		defer close(linodeChan)
		defer close(errorChan)
		if err := processIPAddresses(ctx, handler, linodeChan, &wg); err != nil {
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

func GetIPAddress(ctx context.Context, handler *LinodeAPIHandler, resourceID string) (*models.Resource, error) {
	ipAddress, err := processIPAddress(ctx, handler, resourceID)
	if err != nil {
		return nil, err
	}
	value := models.Resource{
		ID:          ipAddress.Address,
		Name:        ipAddress.Address,
		Description: ipAddress,
	}
	return &value, nil
}

func processIPAddresses(ctx context.Context, handler *LinodeAPIHandler, openaiChan chan<- models.Resource, wg *sync.WaitGroup) error {
	var ipAddresses []model.IPAddressResp
	var ipAddressListResponse model.IPAddressListResponse
	var resp *http.Response
	baseURL := "https://api.linode.com/v4/networking/ips"
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

			if e = json.NewDecoder(resp.Body).Decode(&ipAddressListResponse); e != nil {
				return nil, fmt.Errorf("failed to decode response: %w", e)
			}
			ipAddresses = append(ipAddresses, ipAddressListResponse.Data...)
			return resp, nil
		}

		err = handler.DoRequest(ctx, req, requestFunc)
		if err != nil {
			return fmt.Errorf("error during request handling: %w", err)
		}

		if ipAddressListResponse.Page == ipAddressListResponse.Pages {
			break
		}
		page++
	}

	for _, ipAddress := range ipAddresses {
		wg.Add(1)
		go func(ipAddress model.IPAddressResp) {
			defer wg.Done()
			value := models.Resource{
				ID:   ipAddress.Address,
				Name: ipAddress.Address,
				Description: model.IPAddressDescription{
					Address:    ipAddress.Address,
					Gateway:    ipAddress.Gateway,
					SubnetMask: ipAddress.SubnetMask,
					Prefix:     ipAddress.Prefix,
					Type:       ipAddress.Type,
					Public:     ipAddress.Public,
					RDNS:       ipAddress.RDNS,
					LinodeID:   ipAddress.LinodeID,
					Region:     ipAddress.Region,
					VPCNAT1To1: ipAddress.VPCNAT1To1,
					Reserved:   ipAddress.Reserved,
				},
			}
			openaiChan <- value
		}(ipAddress)
	}
	return nil
}

func processIPAddress(ctx context.Context, handler *LinodeAPIHandler, resourceID string) (*model.IPAddressDescription, error) {
	var ipAddress model.IPAddressDescription
	var resp *http.Response
	baseURL := "https://api.linode.com/v4/networking/ips/"

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

		if e = json.NewDecoder(resp.Body).Decode(&ipAddress); e != nil {
			return nil, fmt.Errorf("failed to decode response: %w", e)
		}
		return resp, e
	}

	err = handler.DoRequest(ctx, req, requestFunc)
	if err != nil {
		return nil, fmt.Errorf("error during request handling: %w", err)
	}
	return &ipAddress, nil
}
