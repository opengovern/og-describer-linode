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
	go func() {
		processNodeBalancers(ctx, handler, linodeChan, &wg)
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

func processNodeBalancers(ctx context.Context, handler *LinodeAPIHandler, openaiChan chan<- models.Resource, wg *sync.WaitGroup) {
	var nodeBalancers []model.NodeBalancerDescription
	var nodeBalancerListResponse *model.NodeBalancerListResponse
	var resp *http.Response
	baseURL := "https://api.linode.com/v4/nodebalancers"
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
			if e = json.NewDecoder(resp.Body).Decode(&nodeBalancerListResponse); e != nil {
				return nil, e
			}
			nodeBalancers = append(nodeBalancers, nodeBalancerListResponse.Data...)
			if nodeBalancerListResponse.Page == nodeBalancerListResponse.Pages {
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
}

func processNodeBalancer(ctx context.Context, handler *LinodeAPIHandler, resourceID string) (*model.NodeBalancerDescription, error) {
	var nodeBalancer *model.NodeBalancerDescription
	var resp *http.Response
	baseURL := "https://api.linode.com/v4/nodebalancers/"
	requestFunc := func(req *http.Request) (*http.Response, error) {
		var e error
		finalURL := fmt.Sprintf("%s%s", baseURL, resourceID)
		req, e = http.NewRequest("GET", finalURL, nil)
		if e != nil {
			return nil, e
		}
		resp, e = handler.Client.Do(req)
		if e = json.NewDecoder(resp.Body).Decode(nodeBalancer); e != nil {
			return nil, e
		}
		return resp, e
	}
	err := handler.DoRequest(ctx, &http.Request{}, requestFunc)
	if err != nil {
		return nil, err
	}
	return nodeBalancer, nil
}
