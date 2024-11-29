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

func ListLongViewClients(ctx context.Context, handler *LinodeAPIHandler, stream *models.StreamSender) ([]models.Resource, error) {
	var wg sync.WaitGroup
	linodeChan := make(chan models.Resource)
	go func() {
		processLongViewClients(ctx, handler, linodeChan, &wg)
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

func GetLongViewClient(ctx context.Context, handler *LinodeAPIHandler, resourceID string) (*models.Resource, error) {
	client, err := processLongViewClient(ctx, handler, resourceID)
	if err != nil {
		return nil, err
	}
	value := models.Resource{
		ID:   strconv.Itoa(client.ID),
		Name: client.Label,
		Description: JSONAllFieldsMarshaller{
			Value: client,
		},
	}
	return &value, nil
}

func processLongViewClients(ctx context.Context, handler *LinodeAPIHandler, openaiChan chan<- models.Resource, wg *sync.WaitGroup) {
	var clients []model.LongViewClientDescription
	var clientListResponse *model.LongViewClientListResponse
	var resp *http.Response
	baseURL := "https://api.linode.com/v4/longview/clients"
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
			if e = json.NewDecoder(resp.Body).Decode(&clientListResponse); e != nil {
				return nil, e
			}
			clients = append(clients, clientListResponse.Data...)
			if clientListResponse.Page == clientListResponse.Pages {
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
	for _, client := range clients {
		wg.Add(1)
		go func(client model.LongViewClientDescription) {
			defer wg.Done()
			value := models.Resource{
				ID:   strconv.Itoa(client.ID),
				Name: client.Label,
				Description: JSONAllFieldsMarshaller{
					Value: client,
				},
			}
			openaiChan <- value
		}(client)
	}
}

func processLongViewClient(ctx context.Context, handler *LinodeAPIHandler, resourceID string) (*model.LongViewClientDescription, error) {
	var client *model.LongViewClientDescription
	var resp *http.Response
	baseURL := "https://api.linode.com/v4/longview/clients/"
	requestFunc := func(req *http.Request) (*http.Response, error) {
		var e error
		finalURL := fmt.Sprintf("%s%s", baseURL, resourceID)
		req, e = http.NewRequest("GET", finalURL, nil)
		if e != nil {
			return nil, e
		}
		resp, e = handler.Client.Do(req)
		if e = json.NewDecoder(resp.Body).Decode(client); e != nil {
			return nil, e
		}
		return resp, e
	}
	err := handler.DoRequest(ctx, &http.Request{}, requestFunc)
	if err != nil {
		return nil, err
	}
	return client, nil
}
