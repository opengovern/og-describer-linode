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

func ListNodeBalancerConfigs(ctx context.Context, handler *provider.LinodeAPIHandler, stream *models.StreamSender) ([]models.Resource, error) {
	var wg sync.WaitGroup
	linodeChan := make(chan models.Resource)
	errorChan := make(chan error, 1) // Buffered channel to capture errors
	accounts, err := ListAccounts(ctx, handler, stream)
	if err != nil {
		return nil, err
	}
	nodeBalancers, err := ListNodeBalancers(ctx, handler, stream)
	if err != nil {
		return nil, err
	}
	fmt.Println(len(nodeBalancers))

	go func() {
		defer close(linodeChan)
		defer close(errorChan)
		for _, nodeBalancer := range nodeBalancers {
			if err := processNodeBalancerConfigs(ctx, handler, accounts[0].ID, nodeBalancer.ID, linodeChan, &wg); err != nil {
				errorChan <- err // Send error to the error channel
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

func processNodeBalancerConfigs(ctx context.Context, handler *provider.LinodeAPIHandler, account, nodeBalancerID string, openaiChan chan<- models.Resource, wg *sync.WaitGroup) error {
	var nodeBalancerConfigs []provider.NodeBalancerConfigJSON
	var nodeBalancerConfigListResponse provider.NodeBalancerConfigListResponse
	fmt.Println("hello")
	var resp *http.Response
	baseURL := "https://api.linode.com/v4/nodebalancers/"
	page := 1

	for {
		params := url.Values{}
		params.Set("page", strconv.Itoa(page))
		params.Set("page_size", "500")
		finalURL := fmt.Sprintf("%s%s/configs?%s", baseURL, nodeBalancerID, params.Encode())

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
			fmt.Println(resp.StatusCode)

			if e = json.NewDecoder(resp.Body).Decode(&nodeBalancerConfigListResponse); e != nil {
				return nil, fmt.Errorf("failed to decode response: %w", e)
			}
			nodeBalancerConfigs = append(nodeBalancerConfigs, nodeBalancerConfigListResponse.Data...)
			return resp, nil
		}

		err = handler.DoRequest(ctx, req, requestFunc)
		if err != nil {
			return fmt.Errorf("error during request handling: %w", err)
		}

		if nodeBalancerConfigListResponse.Page == nodeBalancerConfigListResponse.Pages {
			break
		}
		page++
	}

	for _, nodeBalancerConfig := range nodeBalancerConfigs {
		wg.Add(1)
		go func(nodeBalancerConfig provider.NodeBalancerConfigJSON) {
			defer wg.Done()
			var nodes []provider.Node
			for _, node := range nodeBalancerConfig.Nodes {
				nodes = append(nodes, provider.Node{
					Address:        node.Address,
					ConfigID:       node.ConfigID,
					ID:             node.ID,
					Label:          node.Label,
					Mode:           node.Mode,
					NodeBalancerID: node.NodeBalancerID,
					Status:         node.Status,
					Weight:         node.Weight,
				})
			}
			nodeStatus := provider.NodesStatus{
				Down: nodeBalancerConfig.NodesStatus.Down,
				Up:   nodeBalancerConfig.NodesStatus.Up,
			}
			value := models.Resource{
				ID:   strconv.Itoa(nodeBalancerConfig.ID),
				Name: strconv.Itoa(nodeBalancerConfig.ID),
				Description: provider.NodeBalancerConfigDescription{
					Algorithm:      nodeBalancerConfig.Algorithm,
					Check:          nodeBalancerConfig.Check,
					CheckAttempts:  nodeBalancerConfig.CheckAttempts,
					CheckBody:      nodeBalancerConfig.CheckBody,
					CheckInterval:  nodeBalancerConfig.CheckInterval,
					CheckPassive:   nodeBalancerConfig.CheckPassive,
					CheckPath:      nodeBalancerConfig.CheckPath,
					CheckTimeout:   nodeBalancerConfig.CheckTimeout,
					CipherSuite:    nodeBalancerConfig.CipherSuite,
					ID:             nodeBalancerConfig.ID,
					NodeBalancerID: nodeBalancerConfig.NodeBalancerID,
					Nodes:          nodes,
					NodesStatus:    nodeStatus,
					Port:           nodeBalancerConfig.Port,
					Protocol:       nodeBalancerConfig.Protocol,
					ProxyProtocol:  nodeBalancerConfig.ProxyProtocol,
					SSLCert:        nodeBalancerConfig.SSLCert,
					SSLCommonName:  nodeBalancerConfig.SSLCommonName,
					SSLFingerprint: nodeBalancerConfig.SSLFingerprint,
					SSLKey:         nodeBalancerConfig.SSLKey,
					Stickiness:     nodeBalancerConfig.Stickiness,
					Account:        account,
				},
			}
			openaiChan <- value
		}(nodeBalancerConfig)
	}
	return nil
}
