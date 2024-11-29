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

func ListStackScripts(ctx context.Context, handler *LinodeAPIHandler, stream *models.StreamSender) ([]models.Resource, error) {
	var wg sync.WaitGroup
	linodeChan := make(chan models.Resource)
	go func() {
		processStackScripts(ctx, handler, linodeChan, &wg)
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

func GetStackScript(ctx context.Context, handler *LinodeAPIHandler, resourceID string) (*models.Resource, error) {
	stackScript, err := processStackScript(ctx, handler, resourceID)
	if err != nil {
		return nil, err
	}
	value := models.Resource{
		ID:   strconv.Itoa(stackScript.ID),
		Name: stackScript.Label,
		Description: JSONAllFieldsMarshaller{
			Value: stackScript,
		},
	}
	return &value, nil
}

func processStackScripts(ctx context.Context, handler *LinodeAPIHandler, openaiChan chan<- models.Resource, wg *sync.WaitGroup) {
	var stackScripts []model.StackScriptDescription
	var stackScriptListResponse *model.StackScriptListResponse
	var resp *http.Response
	baseURL := "https://api.linode.com/v4/linode/stackscripts"
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
			if e = json.NewDecoder(resp.Body).Decode(&stackScriptListResponse); e != nil {
				return nil, e
			}
			stackScripts = append(stackScripts, stackScriptListResponse.Data...)
			if stackScriptListResponse.Page == stackScriptListResponse.Pages {
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
	for _, stackScript := range stackScripts {
		wg.Add(1)
		go func(stackScript model.StackScriptDescription) {
			defer wg.Done()
			value := models.Resource{
				ID:   strconv.Itoa(stackScript.ID),
				Name: stackScript.Label,
				Description: JSONAllFieldsMarshaller{
					Value: stackScript,
				},
			}
			openaiChan <- value
		}(stackScript)
	}
}

func processStackScript(ctx context.Context, handler *LinodeAPIHandler, resourceID string) (*model.StackScriptDescription, error) {
	var stackScript *model.StackScriptDescription
	var resp *http.Response
	baseURL := "https://api.linode.com/v4/linode/stackscripts/"
	requestFunc := func(req *http.Request) (*http.Response, error) {
		var e error
		finalURL := fmt.Sprintf("%s%s", baseURL, resourceID)
		req, e = http.NewRequest("GET", finalURL, nil)
		if e != nil {
			return nil, e
		}
		resp, e = handler.Client.Do(req)
		if e = json.NewDecoder(resp.Body).Decode(stackScript); e != nil {
			return nil, e
		}
		return resp, e
	}
	err := handler.DoRequest(ctx, &http.Request{}, requestFunc)
	if err != nil {
		return nil, err
	}
	return stackScript, nil
}
