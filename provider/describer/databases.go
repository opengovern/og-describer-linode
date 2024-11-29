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

func ListDatabases(ctx context.Context, handler *LinodeAPIHandler, stream *models.StreamSender) ([]models.Resource, error) {
	var wg sync.WaitGroup
	linodeChan := make(chan models.Resource)
	go func() {
		processDatabases(ctx, handler, linodeChan, &wg)
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

func processDatabases(ctx context.Context, handler *LinodeAPIHandler, openaiChan chan<- models.Resource, wg *sync.WaitGroup) {
	var databases []model.DatabaseDescription
	var databaseListResponse *model.DatabaseListResponse
	var resp *http.Response
	baseURL := "https://api.linode.com/v4/databases/instances"
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
			if e = json.NewDecoder(resp.Body).Decode(&databaseListResponse); e != nil {
				return nil, e
			}
			databases = append(databases, databaseListResponse.Data...)
			if databaseListResponse.Page == databaseListResponse.Pages {
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
	for _, database := range databases {
		wg.Add(1)
		go func(database model.DatabaseDescription) {
			defer wg.Done()
			value := models.Resource{
				ID:   strconv.Itoa(database.ID),
				Name: database.Label,
				Description: JSONAllFieldsMarshaller{
					Value: database,
				},
			}
			openaiChan <- value
		}(database)
	}
}
