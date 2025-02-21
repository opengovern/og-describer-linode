package describers

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/opengovern/og-describer-linode/discovery/pkg/models"
	"github.com/opengovern/og-describer-linode/discovery/provider"
	"net/http"
	"net/url"
	"strconv"
	"sync"
)

func ListDomains(ctx context.Context, handler *provider.LinodeAPIHandler, stream *models.StreamSender) ([]models.Resource, error) {
	var wg sync.WaitGroup
	linodeChan := make(chan models.Resource)
	errorChan := make(chan error, 1) // Buffered channel to capture errors
	account, err := provider.GetAccount(ctx, handler)
	if err != nil {
		return nil, err
	}

	go func() {
		defer close(linodeChan)
		defer close(errorChan)
		if err := processDomains(ctx, handler, account.EUUID, linodeChan, &wg); err != nil {
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

func GetDomain(ctx context.Context, handler *provider.LinodeAPIHandler, resourceID string) (*models.Resource, error) {
	domain, err := processDomain(ctx, handler, resourceID)
	if err != nil {
		return nil, err
	}
	accounts, err := ListAccounts(ctx, handler, nil)
	if err != nil {
		return nil, err
	}
	value := models.Resource{
		ID:   strconv.Itoa(domain.ID),
		Name: domain.Domain,
		Description: provider.DomainDescription{
			ID:          domain.ID,
			Domain:      domain.Domain,
			Type:        domain.Type,
			Group:       domain.Group,
			Status:      domain.Status,
			Description: domain.Description,
			SOAEmail:    domain.SOAEmail,
			RetrySec:    domain.RetrySec,
			MasterIPs:   domain.MasterIPs,
			AXfrIPs:     domain.AXfrIPs,
			Tags:        domain.Tags,
			ExpireSec:   domain.ExpireSec,
			RefreshSec:  domain.RefreshSec,
			TTLSec:      domain.TTLSec,
			Account:     accounts[0].ID,
		},
	}
	return &value, nil
}

func processDomains(ctx context.Context, handler *provider.LinodeAPIHandler, account string, openaiChan chan<- models.Resource, wg *sync.WaitGroup) error {
	var domains []provider.DomainRecord
	var domainListResponse provider.DomainListResponse
	var resp *http.Response
	baseURL := "https://api.linode.com/v4/domains"
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

			if e = json.NewDecoder(resp.Body).Decode(&domainListResponse); e != nil {
				return nil, fmt.Errorf("failed to decode response: %w", e)
			}
			domains = append(domains, domainListResponse.Data...)
			return resp, nil
		}

		err = handler.DoRequest(ctx, req, requestFunc)
		if err != nil {
			return fmt.Errorf("error during request handling: %w", err)
		}

		if domainListResponse.Page == domainListResponse.Pages {
			break
		}
		page++
	}
	for _, domain := range domains {
		wg.Add(1)
		go func(domain provider.DomainRecord) {
			defer wg.Done()
			value := models.Resource{
				ID:   strconv.Itoa(domain.ID),
				Name: domain.Domain,
				Description: provider.DomainDescription{
					ID:          domain.ID,
					Domain:      domain.Domain,
					Type:        domain.Type,
					Group:       domain.Group,
					Status:      domain.Status,
					Description: domain.Description,
					SOAEmail:    domain.SOAEmail,
					RetrySec:    domain.RetrySec,
					MasterIPs:   domain.MasterIPs,
					AXfrIPs:     domain.AXfrIPs,
					Tags:        domain.Tags,
					ExpireSec:   domain.ExpireSec,
					RefreshSec:  domain.RefreshSec,
					TTLSec:      domain.TTLSec,
					Account:     account,
				},
			}
			openaiChan <- value
		}(domain)
	}
	return nil
}

func processDomain(ctx context.Context, handler *provider.LinodeAPIHandler, resourceID string) (*provider.DomainDescription, error) {
	var domain provider.DomainDescription
	var resp *http.Response
	baseURL := "https://api.linode.com/v4/domains/"

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
		defer resp.Body.Close()

		if e = json.NewDecoder(resp.Body).Decode(&domain); e != nil {
			return nil, fmt.Errorf("failed to decode response: %w", e)
		}
		return resp, e
	}

	err = handler.DoRequest(ctx, req, requestFunc)
	if err != nil {
		return nil, fmt.Errorf("error during request handling: %w", err)
	}
	return &domain, nil
}
