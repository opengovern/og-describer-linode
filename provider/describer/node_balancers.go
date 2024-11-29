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

func ListNodeBalancers(ctx context.Context, handler *LinodeAPIHandler, stream *models.StreamSender) ([]models.Resource, error) {
	var wg sync.WaitGroup
	linodeChan := make(chan models.Resource)
	errorChan := make(chan error, 1) // Buffered channel to capture errors

	go func() {
		defer close(linodeChan)
		defer close(errorChan)
		if err := processNodeBalancers(ctx, handler, linodeChan, &wg); err != nil {
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

func GetNodeBalancer(ctx context.Context, handler *LinodeAPIHandler, resourceID string) (*models.Resource, error) {
	nodeBalancer, err := processNodeBalancer(ctx, handler, resourceID)
	if err != nil {
		return nil, err
	}
	var name string
	if nodeBalancer.Label != nil {
		name = *nodeBalancer.Label
	}
	value := models.Resource{
		ID:   strconv.Itoa(nodeBalancer.ID),
		Name: name,
		Description: JSONAllFieldsMarshaller{
			Value: nodeBalancer,
		},
	}
	return &value, nil
}

func processNodeBalancers(ctx context.Context, handler *LinodeAPIHandler, openaiChan chan<- models.Resource, wg *sync.WaitGroup) error {
	var nodeBalancers []model.NodeBalancerDescription
	var nodeBalancerListResponse *model.NodeBalancerListResponse
	var resp *http.Response
	baseURL := "https://api.linode.com/v4/nodebalancers"
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

			if e = json.NewDecoder(resp.Body).Decode(&nodeBalancerListResponse); e != nil {
				return nil, fmt.Errorf("failed to decode response: %w", e)
			}
			nodeBalancers = append(nodeBalancers, nodeBalancerListResponse.Data...)
			return resp, nil
		}

		err = handler.DoRequest(ctx, req, requestFunc)
		if err != nil {
			return fmt.Errorf("error during request handling: %w", err)
		}

		if nodeBalancerListResponse.Page == nodeBalancerListResponse.Pages {
			break
		}
		page++
	}

	for _, nodeBalancer := range nodeBalancers {
		wg.Add(1)
		go func(nodeBalancer model.NodeBalancerDescription) {
			defer wg.Done()
			var name string
			if nodeBalancer.Label != nil {
				name = *nodeBalancer.Label
			}
			value := models.Resource{
				ID:   strconv.Itoa(nodeBalancer.ID),
				Name: name,
				Description: JSONAllFieldsMarshaller{
					Value: nodeBalancer,
				},
			}
			openaiChan <- value
		}(nodeBalancer)
	}
	return nil
}

func processNodeBalancer(ctx context.Context, handler *LinodeAPIHandler, resourceID string) (*model.NodeBalancerDescription, error) {
	var nodeBalancer *model.NodeBalancerDescription
	var resp *http.Response
	baseURL := "https://api.linode.com/v4/nodebalancers/"

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

		if e = json.NewDecoder(resp.Body).Decode(nodeBalancer); e != nil {
			return nil, fmt.Errorf("failed to decode response: %w", e)
		}
		return resp, e
	}

	err = handler.DoRequest(ctx, req, requestFunc)
	if err != nil {
		return nil, fmt.Errorf("error during request handling: %w", err)
	}
	return nodeBalancer, nil
}
