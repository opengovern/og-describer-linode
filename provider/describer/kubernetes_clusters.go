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

func ListKubernetesClusters(ctx context.Context, handler *LinodeAPIHandler, stream *models.StreamSender) ([]models.Resource, error) {
	var wg sync.WaitGroup
	linodeChan := make(chan models.Resource)
	go func() {
		processKubernetesClusters(ctx, handler, linodeChan, &wg)
		wg.Wait()
		close(linodeChan)
	}()
	var values []models.Resource
	for value := range linodeChan {
		if stream != nil {
			if err := (*stream)(value); err != nil {
				return nil, err
			}
		} else {
			values = append(values, value)
		}
	}
	return values, nil
}

func GetKubernetesCluster(ctx context.Context, handler *LinodeAPIHandler, resourceID string) (*models.Resource, error) {
	cluster, err := processKubernetesCluster(ctx, handler, resourceID)
	if err != nil {
		return nil, err
	}
	value := models.Resource{
		ID:   strconv.Itoa(cluster.ID),
		Name: cluster.Label,
		Description: JSONAllFieldsMarshaller{
			Value: cluster,
		},
	}
	return &value, nil
}

func processKubernetesClusters(ctx context.Context, handler *LinodeAPIHandler, openaiChan chan<- models.Resource, wg *sync.WaitGroup) {
	var clusters []model.KubernetesClusterDescription
	var clusterListResponse *model.KubernetesClusterListResponse
	var resp *http.Response
	baseURL := "https://api.linode.com/v4/lke/clusters"
	requestFunc := func(req *http.Request) (*http.Response, error) {
		var e error
		page := 1
		for {
			params := url.Values{}
			params.Set("page", strconv.Itoa(page))
			params.Set("page_size", "500")
			finalURL := fmt.Sprintf("%s?%s", baseURL, params.Encode())
			req, e = http.NewRequest("GET", finalURL, nil)
			if e != nil {
				return nil, e
			}
			resp, e = handler.Client.Do(req)
			if e = json.NewDecoder(resp.Body).Decode(&clusterListResponse); e != nil {
				return nil, e
			}
			clusters = append(clusters, clusterListResponse.Data...)
			if clusterListResponse.Page == clusterListResponse.Pages {
				break
			}
			page += 1
		}
		return resp, e
	}
	err := handler.DoRequest(ctx, &http.Request{}, requestFunc)
	if err != nil {
		return
	}
	for _, cluster := range clusters {
		wg.Add(1)
		go func(cluster model.KubernetesClusterDescription) {
			defer wg.Done()
			value := models.Resource{
				ID:   strconv.Itoa(cluster.ID),
				Name: cluster.Label,
				Description: JSONAllFieldsMarshaller{
					Value: cluster,
				},
			}
			openaiChan <- value
		}(cluster)
	}
}

func processKubernetesCluster(ctx context.Context, handler *LinodeAPIHandler, resourceID string) (*model.KubernetesClusterDescription, error) {
	var cluster *model.KubernetesClusterDescription
	var resp *http.Response
	baseURL := "https://api.linode.com/v4/lke/clusters/"
	requestFunc := func(req *http.Request) (*http.Response, error) {
		var e error
		finalURL := fmt.Sprintf("%s%s", baseURL, resourceID)
		req, e = http.NewRequest("GET", finalURL, nil)
		if e != nil {
			return nil, e
		}
		resp, e = handler.Client.Do(req)
		if e = json.NewDecoder(resp.Body).Decode(cluster); e != nil {
			return nil, e
		}
		return resp, e
	}
	err := handler.DoRequest(ctx, &http.Request{}, requestFunc)
	if err != nil {
		return nil, err
	}
	return cluster, nil
}
