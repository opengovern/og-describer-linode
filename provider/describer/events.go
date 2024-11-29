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
	errorChan := make(chan error, 1) // Buffered channel to capture errors

	go func() {
		defer close(linodeChan)
		defer close(errorChan)
		if err := processEvents(ctx, handler, linodeChan, &wg); err != nil {
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

func processEvents(ctx context.Context, handler *LinodeAPIHandler, openaiChan chan<- models.Resource, wg *sync.WaitGroup) error {
	var events []model.EventDescription
	var eventListResponse *model.EventListResponse
	var resp *http.Response
	baseURL := "https://api.linode.com/v4/account/events"
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

			if e = json.NewDecoder(resp.Body).Decode(&eventListResponse); e != nil {
				return nil, fmt.Errorf("failed to decode response: %w", e)
			}
			events = append(events, eventListResponse.Data...)
			return resp, nil
		}

		err = handler.DoRequest(ctx, req, requestFunc)
		if err != nil {
			return fmt.Errorf("error during request handling: %w", err)
		}

		if eventListResponse.Page == eventListResponse.Pages {
			break
		}
		page++
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
	return nil
}

func processEvent(ctx context.Context, handler *LinodeAPIHandler, resourceID string) (*model.EventDescription, error) {
	var event *model.EventDescription
	var resp *http.Response
	baseURL := "https://api.linode.com/v4/account/events/"

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

		if e = json.NewDecoder(resp.Body).Decode(event); e != nil {
			return nil, fmt.Errorf("failed to decode response: %w", e)
		}
		return resp, e
	}

	err = handler.DoRequest(ctx, req, requestFunc)
	if err != nil {
		return nil, fmt.Errorf("error during request handling: %w", err)
	}
	return event, nil
}
