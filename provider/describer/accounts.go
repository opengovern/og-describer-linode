package describer

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/opengovern/og-describer-linode/pkg/sdk/models"
	"github.com/opengovern/og-describer-linode/provider/model"
	"net/http"
)

func ListAccounts(ctx context.Context, handler *LinodeAPIHandler, stream *models.StreamSender) ([]models.Resource, error) {
	account, err := processAccount(ctx, handler)
	if err != nil {
		return nil, err
	}
	var values []models.Resource
	value := models.Resource{
		ID:   account.Email,
		Name: account.Email,
		Description: model.AccountDescription{
			Email:   account.Email,
			City:    account.City,
			Company: account.Company,
			Country: account.Country,
			Euuid:   account.EUUID,
		},
	}
	if stream != nil {
		if err = (*stream)(value); err != nil {
			return nil, err
		}
	} else {
		values = append(values, value)
	}
	return values, nil
}

func processAccount(ctx context.Context, handler *LinodeAPIHandler) (*model.Account, error) {
	var account model.Account
	var resp *http.Response
	baseURL := "https://api.linode.com/v4/account"

	req, err := http.NewRequest("GET", baseURL, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	requestFunc := func(req *http.Request) (*http.Response, error) {
		var e error
		resp, e = handler.Client.Do(req)
		if e != nil {
			return nil, fmt.Errorf("request execution failed: %w", e)
		}
		defer resp.Body.Close()

		if e = json.NewDecoder(resp.Body).Decode(&account); e != nil {
			return nil, fmt.Errorf("failed to decode response: %w", e)
		}
		return resp, e
	}

	err = handler.DoRequest(ctx, req, requestFunc)
	if err != nil {
		return nil, fmt.Errorf("error during request handling: %w", err)
	}
	return &account, nil
}
