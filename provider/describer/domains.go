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

func ListDomains(ctx context.Context, handler *LinodeAPIHandler, stream *models.StreamSender) ([]models.Resource, error) {
	var wg sync.WaitGroup
	linodeChan := make(chan models.Resource)
	go func() {
		processDomains(ctx, handler, linodeChan, &wg)
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

func GetDomain(ctx context.Context, handler *LinodeAPIHandler, resourceID string) (*models.Resource, error) {
	domain, err := processDomain(ctx, handler, resourceID)
	if err != nil {
		return nil, err
	}
	value := models.Resource{
		ID:   strconv.Itoa(domain.ID),
		Name: domain.Domain,
		Description: JSONAllFieldsMarshaller{
			Value: domain,
		},
	}
	return &value, nil
}

func processDomains(ctx context.Context, handler *LinodeAPIHandler, openaiChan chan<- models.Resource, wg *sync.WaitGroup) {
	var domains []model.DomainDescription
	var domainListResponse *model.DomainListResponse
	var resp *http.Response
	baseURL := "https://api.linode.com/v4/domains"
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
			if e = json.NewDecoder(resp.Body).Decode(&domainListResponse); e != nil {
				return nil, e
			}
			domains = append(domains, domainListResponse.Data...)
			if domainListResponse.Page == domainListResponse.Pages {
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
	for _, domain := range domains {
		wg.Add(1)
		go func(domain model.DomainDescription) {
			defer wg.Done()
			value := models.Resource{
				ID:   strconv.Itoa(domain.ID),
				Name: domain.Domain,
				Description: JSONAllFieldsMarshaller{
					Value: domain,
				},
			}
			openaiChan <- value
		}(domain)
	}
}

func processDomain(ctx context.Context, handler *LinodeAPIHandler, resourceID string) (*model.DomainDescription, error) {
	var domain *model.DomainDescription
	var resp *http.Response
	baseURL := "https://api.linode.com/v4/domains/"
	requestFunc := func(req *http.Request) (*http.Response, error) {
		var e error
		finalURL := fmt.Sprintf("%s%s", baseURL, resourceID)
		req, e = http.NewRequest("GET", finalURL, nil)
		if e != nil {
			return nil, e
		}
		resp, e = handler.Client.Do(req)
		if e = json.NewDecoder(resp.Body).Decode(domain); e != nil {
			return nil, e
		}
		return resp, e
	}
	err := handler.DoRequest(ctx, &http.Request{}, requestFunc)
	if err != nil {
		return nil, err
	}
	return domain, nil
}
