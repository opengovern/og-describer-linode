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

func ListNodePools(ctx context.Context, handler *provider.LinodeAPIHandler, stream *models.StreamSender) ([]models.Resource, error) {
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
			if err := processNodePools(ctx, handler, account.EUUID, strconv.Itoa(cluster.ID), linodeChan, &wg); err != nil {
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

func processNodePools(ctx context.Context, handler *provider.LinodeAPIHandler, account, clusterID string, linodeChan chan<- models.Resource, wg *sync.WaitGroup) error {
	var nodePools []provider.NodePoolJSON
	var nodePoolListResponse provider.NodePoolListResponse
	var resp *http.Response
	baseURL := "https://api.linode.com/v4/lke/clusters/"
	page := 1

	for {
		params := url.Values{}
		params.Set("page", strconv.Itoa(page))
		params.Set("page_size", "500")
		finalURL := fmt.Sprintf("%s%s/pools?%s", baseURL, clusterID, params.Encode())

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

			if e = json.NewDecoder(resp.Body).Decode(&nodePoolListResponse); e != nil {
				return nil, fmt.Errorf("failed to decode response: %w", e)
			}
			nodePools = append(nodePools, nodePoolListResponse.Data...)
			return resp, nil
		}

		err = handler.DoRequest(ctx, req, requestFunc)
		if err != nil {
			return fmt.Errorf("error during request handling: %w", err)
		}

		if nodePoolListResponse.Page == nodePoolListResponse.Pages {
			break
		}
		page++
	}

	for _, nodePool := range nodePools {
		wg.Add(1)
		go func(nodePool provider.NodePoolJSON) {
			defer wg.Done()
			autoScaler := provider.Autoscaler{
				Enabled: nodePool.Autoscaler.Enabled,
				Max:     nodePool.Autoscaler.Max,
				Min:     nodePool.Autoscaler.Min,
			}
			var disks []provider.Disk
			for _, disk := range nodePool.Disks {
				disks = append(disks, provider.Disk{
					Size: disk.Size,
					Type: disk.Type,
				})
			}
			var nodes []provider.NodePoolNode
			for _, node := range nodePool.Nodes {
				nodes = append(nodes, provider.NodePoolNode{
					ID:         node.ID,
					InstanceID: node.InstanceID,
					Status:     node.Status,
				})
			}
			var taints []provider.Taint
			for _, taint := range nodePool.Taints {
				taints = append(taints, provider.Taint{
					Effect: taint.Effect,
					Key:    taint.Key,
					Value:  taint.Value,
				})
			}
			value := models.Resource{
				ID:   strconv.Itoa(nodePool.ID),
				Name: strconv.Itoa(nodePool.ID),
				Description: provider.NodePoolDescription{
					Autoscaler:     autoScaler,
					ClusterID:      clusterID,
					Count:          nodePool.Count,
					DiskEncryption: nodePool.DiskEncryption,
					Disks:          disks,
					ID:             nodePool.ID,
					Labels:         nodePool.Labels,
					Nodes:          nodes,
					Tags:           nodePool.Tags,
					Taints:         taints,
					Type:           nodePool.Type,
					Account:        account,
				},
			}
			linodeChan <- value
		}(nodePool)
	}
	return nil
}
