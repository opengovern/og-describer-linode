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

func ListNodeBalancerNodes(ctx context.Context, handler *provider.LinodeAPIHandler, stream *models.StreamSender) ([]models.Resource, error) {
	var wg sync.WaitGroup
	linodeChan := make(chan models.Resource)
	errorChan := make(chan error, 1) // Buffered channel to capture errors
	nodeBalancers, err := ListNodeBalancers(ctx, handler, stream)
	if err != nil {
		return nil, err
	}

	go func() {
		defer close(linodeChan)
		defer close(errorChan)
		for _, nodeBalancer := range nodeBalancers {
			configs, err := provider.ListConfigs(ctx, handler, nodeBalancer.ID)
			if err != nil {
				errorChan <- err // Send error to the error channel
			}
			for _, config := range configs {
				if err = processNodeBalancerNodes(ctx, handler, nodeBalancer.ID, strconv.Itoa(config.ID), linodeChan, &wg); err != nil {
					errorChan <- err // Send error to the error channel
				}
			}
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

func processNodeBalancerNodes(ctx context.Context, handler *provider.LinodeAPIHandler, nodeBalancerID, configID string, openaiChan chan<- models.Resource, wg *sync.WaitGroup) error {
	var nodeBalancerNodes []provider.NodeRespJSON
	var nodeBalancerNodeListResponse provider.NodeBalancerNodeListResponse
	var resp *http.Response
	baseURL := "https://api.linode.com/v4/nodebalancers/"
	page := 1

	for {
		params := url.Values{}
		params.Set("page", strconv.Itoa(page))
		params.Set("page_size", "500")
		finalURL := fmt.Sprintf("%s%s/configs/%s/nodes?%s", baseURL, nodeBalancerID, configID, params.Encode())

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

			if e = json.NewDecoder(resp.Body).Decode(&nodeBalancerNodeListResponse); e != nil {
				return nil, fmt.Errorf("failed to decode response: %w", e)
			}
			nodeBalancerNodes = append(nodeBalancerNodes, nodeBalancerNodeListResponse.Data...)
			return resp, nil
		}

		err = handler.DoRequest(ctx, req, requestFunc)
		if err != nil {
			return fmt.Errorf("error during request handling: %w", err)
		}

		if nodeBalancerNodeListResponse.Page == nodeBalancerNodeListResponse.Pages {
			break
		}
		page++
	}

	for _, nodeBalancerNode := range nodeBalancerNodes {
		wg.Add(1)
		go func(nodeBalancerNode provider.NodeRespJSON) {
			defer wg.Done()
			value := models.Resource{
				ID:   strconv.Itoa(nodeBalancerNode.ID),
				Name: strconv.Itoa(nodeBalancerNode.ID),
				Description: provider.NodeDescription{
					Address:        nodeBalancerNode.Address,
					ConfigID:       nodeBalancerNode.ConfigID,
					ID:             nodeBalancerNode.ID,
					Label:          nodeBalancerNode.Label,
					Mode:           nodeBalancerNode.Mode,
					NodeBalancerID: nodeBalancerNode.NodeBalancerID,
					Status:         nodeBalancerNode.Status,
					Weight:         nodeBalancerNode.Weight,
				},
			}
			openaiChan <- value
		}(nodeBalancerNode)
	}
	return nil
}
