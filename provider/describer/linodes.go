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

func ListLinodeInstances(ctx context.Context, handler *LinodeAPIHandler, stream *models.StreamSender) ([]models.Resource, error) {
	var wg sync.WaitGroup
	linodeChan := make(chan models.Resource)
	go func() {
		processLinodeInstances(ctx, handler, linodeChan, &wg)
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

func GetLinodeInstance(ctx context.Context, handler *LinodeAPIHandler, resourceID string) (*models.Resource, error) {
	linode, err := processLinodeInstance(ctx, handler, resourceID)
	if err != nil {
		return nil, err
	}
	value := models.Resource{
		ID:   strconv.Itoa(linode.ID),
		Name: linode.Label,
		Description: JSONAllFieldsMarshaller{
			Value: linode,
		},
	}
	return &value, nil
}

func processLinodeInstances(ctx context.Context, handler *LinodeAPIHandler, openaiChan chan<- models.Resource, wg *sync.WaitGroup) {
	var linodeInstances []model.LinodeDescription
	var linodeListResponse *model.LinodeListResponse
	var resp *http.Response
	baseURL := "https://api.linode.com/v4/linode/instances"
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
			if e = json.NewDecoder(resp.Body).Decode(&linodeListResponse); e != nil {
				return nil, e
			}
			linodeInstances = append(linodeInstances, linodeListResponse.Data...)
			if linodeListResponse.Page == linodeListResponse.Pages {
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
	for _, linode := range linodeInstances {
		wg.Add(1)
		go func(linode model.LinodeDescription) {
			defer wg.Done()
			value := models.Resource{
				ID:   strconv.Itoa(linode.ID),
				Name: linode.Label,
				Description: JSONAllFieldsMarshaller{
					Value: linode,
				},
			}
			openaiChan <- value
		}(linode)
	}
}

func processLinodeInstance(ctx context.Context, handler *LinodeAPIHandler, resourceID string) (*model.LinodeDescription, error) {
	var linode *model.LinodeDescription
	var resp *http.Response
	baseURL := "https://api.linode.com/v4/linode/instances/"
	requestFunc := func(req *http.Request) (*http.Response, error) {
		var e error
		finalURL := fmt.Sprintf("%s%s", baseURL, resourceID)
		req, e = http.NewRequest("GET", finalURL, nil)
		if e != nil {
			return nil, e
		}
		resp, e = handler.Client.Do(req)
		if e = json.NewDecoder(resp.Body).Decode(linode); e != nil {
			return nil, e
		}
		return resp, e
	}
	err := handler.DoRequest(ctx, &http.Request{}, requestFunc)
	if err != nil {
		return nil, err
	}
	return linode, nil
}
