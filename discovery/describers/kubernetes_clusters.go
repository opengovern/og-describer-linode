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

func ListKubernetesClusters(ctx context.Context, handler *provider.LinodeAPIHandler, stream *models.StreamSender) ([]models.Resource, error) {
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
		if err := processKubernetesClusters(ctx, handler, account.EUUID, linodeChan, &wg); err != nil {
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

func GetKubernetesCluster(ctx context.Context, handler *provider.LinodeAPIHandler, resourceID string) (*models.Resource, error) {
	cluster, err := processKubernetesCluster(ctx, handler, resourceID)
	if err != nil {
		return nil, err
	}
	account, err := provider.GetAccount(ctx, handler)
	if err != nil {
		return nil, err
	}
	value := models.Resource{
		ID:   strconv.Itoa(cluster.ID),
		Name: cluster.Label,
		Description: provider.KubernetesClusterDescription{
			ID:         cluster.ID,
			Label:      cluster.Label,
			Region:     cluster.Region,
			Created:    cluster.Created,
			Updated:    cluster.Updated,
			Status:     cluster.Status,
			K8sVersion: cluster.K8sVersion,
			Tags:       cluster.Tags,
			ControlPlane: provider.LKEClusterControlPlane{
				HighAvailability: cluster.ControlPlane.HighAvailability,
			},
			Account: account.EUUID,
		},
	}
	return &value, nil
}

func processKubernetesClusters(ctx context.Context, handler *provider.LinodeAPIHandler, account string, openaiChan chan<- models.Resource, wg *sync.WaitGroup) error {
	var clusters []provider.KubernetesClusterResp
	var clusterListResponse provider.KubernetesClusterListResponse
	var resp *http.Response
	baseURL := "https://api.linode.com/v4/lke/clusters"
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

			if e = json.NewDecoder(resp.Body).Decode(&clusterListResponse); e != nil {
				return nil, fmt.Errorf("failed to decode response: %w", e)
			}
			clusters = append(clusters, clusterListResponse.Data...)
			return resp, nil
		}

		err = handler.DoRequest(ctx, req, requestFunc)
		if err != nil {
			return fmt.Errorf("error during request handling: %w", err)
		}

		if clusterListResponse.Page == clusterListResponse.Pages {
			break
		}
		page++
	}

	for _, cluster := range clusters {
		wg.Add(1)
		go func(cluster provider.KubernetesClusterResp) {
			defer wg.Done()
			value := models.Resource{
				ID:   strconv.Itoa(cluster.ID),
				Name: cluster.Label,
				Description: provider.KubernetesClusterDescription{
					ID:         cluster.ID,
					Label:      cluster.Label,
					Region:     cluster.Region,
					Created:    cluster.Created,
					Updated:    cluster.Updated,
					Status:     cluster.Status,
					K8sVersion: cluster.K8sVersion,
					Tags:       cluster.Tags,
					ControlPlane: provider.LKEClusterControlPlane{
						HighAvailability: cluster.ControlPlane.HighAvailability,
					},
					Account: account,
				},
			}
			openaiChan <- value
		}(cluster)
	}
	return nil
}

func processKubernetesCluster(ctx context.Context, handler *provider.LinodeAPIHandler, resourceID string) (*provider.KubernetesClusterDescription, error) {
	var cluster provider.KubernetesClusterDescription
	var resp *http.Response
	baseURL := "https://api.linode.com/v4/lke/clusters/"

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

		if e = json.NewDecoder(resp.Body).Decode(&cluster); e != nil {
			return nil, fmt.Errorf("failed to decode response: %w", e)
		}
		return resp, e
	}

	err = handler.DoRequest(ctx, req, requestFunc)
	if err != nil {
		return nil, fmt.Errorf("error during request handling: %w", err)
	}
	return &cluster, nil
}
