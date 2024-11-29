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

func ListEvents(ctx context.Context, handler *LinodeAPIHandler, stream *models.StreamSender) ([]models.Resource, error) {
	var wg sync.WaitGroup
	linodeChan := make(chan models.Resource)
	go func() {
		processEvents(ctx, handler, linodeChan, &wg)
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

func GetEvent(ctx context.Context, handler *LinodeAPIHandler, resourceID string) (*models.Resource, error) {
	event, err := processEvent(ctx, handler, resourceID)
	if err != nil {
		return nil, err
	}
	value := models.Resource{
		ID:   strconv.Itoa(event.ID),
		Name: event.Username,
		Description: JSONAllFieldsMarshaller{
			Value: event,
		},
	}
	return &value, nil
}

func processEvents(ctx context.Context, handler *LinodeAPIHandler, openaiChan chan<- models.Resource, wg *sync.WaitGroup) {
	var events []model.EventDescription
	var eventListResponse *model.EventListResponse
	var resp *http.Response
	baseURL := "https://api.linode.com/v4/account/events"
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
			if e = json.NewDecoder(resp.Body).Decode(&eventListResponse); e != nil {
				return nil, e
			}
			events = append(events, eventListResponse.Data...)
			if eventListResponse.Page == eventListResponse.Pages {
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
	for _, event := range events {
		wg.Add(1)
		go func(event model.EventDescription) {
			defer wg.Done()
			value := models.Resource{
				ID:   strconv.Itoa(event.ID),
				Name: event.Username,
				Description: JSONAllFieldsMarshaller{
					Value: event,
				},
			}
			openaiChan <- value
		}(event)
	}
}

func processEvent(ctx context.Context, handler *LinodeAPIHandler, resourceID string) (*model.EventDescription, error) {
	var event *model.EventDescription
	var resp *http.Response
	baseURL := "https://api.linode.com/v4/account/events/"
	requestFunc := func(req *http.Request) (*http.Response, error) {
		var e error
		finalURL := fmt.Sprintf("%s%s", baseURL, resourceID)
		req, e = http.NewRequest("GET", finalURL, nil)
		if e != nil {
			return nil, e
		}
		resp, e = handler.Client.Do(req)
		if e = json.NewDecoder(resp.Body).Decode(event); e != nil {
			return nil, e
		}
		return resp, e
	}
	err := handler.DoRequest(ctx, &http.Request{}, requestFunc)
	if err != nil {
		return nil, err
	}
	return event, nil
}
