package describers

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/opengovern/og-describer-linode/discovery/pkg/models"
	"github.com/opengovern/og-describer-linode/discovery/provider"
	"net/http"
	"strconv"
	"sync"
)

func ListNodes(ctx context.Context, handler *provider.LinodeAPIHandler, stream *models.StreamSender) ([]models.Resource, error) {
	var wg sync.WaitGroup
	linodeChan := make(chan models.Resource)
	errorChan := make(chan error, 1) // Buffered channel to capture errors
	account, err := provider.GetAccount(ctx, handler)
	if err != nil {
		return nil, err
	}
	clusters, err := provider.ListClusters(ctx, handler)
	if err != nil {
		return nil, err
	}

	go func() {
		defer close(linodeChan)
		defer close(errorChan)
		for _, cluster := range clusters {
			if err := processNodes(ctx, handler, account.EUUID, strconv.Itoa(cluster.ID), linodeChan, &wg); err != nil {
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

func GetNode(ctx context.Context, handler *provider.LinodeAPIHandler, resourceID string) (*models.Resource, error) {
	clusters, err := provider.ListClusters(ctx, handler)
	if err != nil {
		return nil, err
	}
	var clusterID string
	for _, cluster := range clusters {
		nodePools, err := provider.ListNodePools(ctx, handler, strconv.Itoa(cluster.ID))
		if err != nil {
			return nil, err
		}
		for _, nodePool := range nodePools {
			for _, node := range nodePool.Nodes {
				if node.ID == resourceID {
					clusterID = strconv.Itoa(cluster.ID)
					break
				}
			}
			if clusterID != "" {
				break
			}
		}
		if clusterID != "" {
			break
		}
	}
	account, err := provider.GetAccount(ctx, handler)
	if err != nil {
		return nil, err
	}
	node, err := processNode(ctx, handler, clusterID, resourceID)
	if err != nil {
		return nil, err
	}
	value := models.Resource{
		ID:   node.ID,
		Name: node.ID,
		Description: provider.ClusterNodeDescription{
			ID:         node.ID,
			InstanceID: node.InstanceID,
			ClusterID:  clusterID,
			Status:     node.Status,
			Account:    account.EUUID,
		},
	}
	return &value, nil
}

func processNodes(ctx context.Context, handler *provider.LinodeAPIHandler, account, clusterID string, linodeChan chan<- models.Resource, wg *sync.WaitGroup) error {
	nodePools, err := provider.ListNodePools(ctx, handler, clusterID)
	if err != nil {
		return err
	}

	for _, nodePool := range nodePools {
		for _, node := range nodePool.Nodes {
			wg.Add(1)
			go func(node provider.NodePoolNodeJSON) {
				defer wg.Done()
				value := models.Resource{
					ID:   node.ID,
					Name: node.ID,
					Description: provider.ClusterNodeDescription{
						ID:         node.ID,
						InstanceID: node.InstanceID,
						ClusterID:  clusterID,
						Status:     node.Status,
						Account:    account,
					},
				}
				linodeChan <- value
			}(node)
		}
	}
	return nil
}

func processNode(ctx context.Context, handler *provider.LinodeAPIHandler, clusterID, resourceID string) (*provider.ClusterNodeJSON, error) {
	var node provider.ClusterNodeJSON
	var resp *http.Response
	baseURL := "https://api.linode.com/v4/lke/clusters/"

	finalURL := fmt.Sprintf("%s%s/nodes/%s", baseURL, clusterID, resourceID)
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

		if e = json.NewDecoder(resp.Body).Decode(&node); e != nil {
			return nil, fmt.Errorf("failed to decode response: %w", e)
		}
		return resp, e
	}

	err = handler.DoRequest(ctx, req, requestFunc)
	if err != nil {
		return nil, fmt.Errorf("error during request handling: %w", err)
	}
	return &node, nil
}
