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

func ListFirewalls(ctx context.Context, handler *LinodeAPIHandler, stream *models.StreamSender) ([]models.Resource, error) {
	var wg sync.WaitGroup
	linodeChan := make(chan models.Resource)
	linodeInstances, err := getLinodeInstances(ctx, handler)
	if err != nil {
		return nil, err
	}
	go func() {
		for _, linodeInstance := range linodeInstances {
			processFirewalls(ctx, handler, linodeInstance.ID, linodeChan, &wg)
		}
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

func processFirewalls(ctx context.Context, handler *LinodeAPIHandler, linodeInstanceID int, openaiChan chan<- models.Resource, wg *sync.WaitGroup) {
	var firewalls []model.FirewallDescription
	var firewallListResponse *model.FirewallListResponse
	var resp *http.Response
	baseURL := "https://api.linode.com/apiVersion/linode/instances/"
	requestFunc := func(req *http.Request) (*http.Response, error) {
		var e error
		page := 1
		for {
			params := url.Values{}
			params.Set("page", strconv.Itoa(page))
			params.Set("page_size", "500")
			finalURL := fmt.Sprintf("%s/%d/firewalls?%s", baseURL, linodeInstanceID, params.Encode())
			req, e = http.NewRequest("GET", finalURL, nil)
			if e != nil {
				return nil, e
			}
			resp, e = handler.Client.Do(req)
			if e = json.NewDecoder(resp.Body).Decode(&firewallListResponse); e != nil {
				return nil, e
			}
			firewalls = append(firewalls, firewallListResponse.Data...)
			if firewallListResponse.Page == firewallListResponse.Pages {
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
	for _, firewall := range firewalls {
		wg.Add(1)
		go func(firewall model.FirewallDescription) {
			defer wg.Done()
			value := models.Resource{
				ID:   strconv.Itoa(firewall.ID),
				Name: firewall.Label,
				Description: JSONAllFieldsMarshaller{
					Value: firewall,
				},
			}
			openaiChan <- value
		}(firewall)
	}
}
