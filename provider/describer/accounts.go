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
	balance := fmt.Sprintf("%f", account.Balance)
	balanceUninvoiced := fmt.Sprintf("%f", account.BalanceUninvoiced)
	var values []models.Resource
	value := models.Resource{
		ID:   account.Email,
		Name: fmt.Sprintf("%s %s", account.FirstName, account.LastName),
		Description: JSONAllFieldsMarshaller{
			Value: model.AccountDescription{
				Email:             account.Email,
				Address1:          account.Address1,
				Address2:          account.Address2,
				Balance:           balance,
				BalanceUninvoiced: balanceUninvoiced,
				City:              account.City,
				Company:           account.Company,
				Country:           account.Country,
				CreditCard:        account.CreditCard,
				FirstName:         account.FirstName,
				LastName:          account.LastName,
				Euuid:             account.Euuid,
				Phone:             account.Phone,
				State:             account.State,
				TaxID:             account.TaxID,
				Zip:               account.Zip,
			},
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
