package describer

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/opengovern/og-describer-linode/pkg/sdk/models"
	"github.com/opengovern/og-describer-linode/provider/model"
	"net/http"
)

func GetAccount(ctx context.Context, handler *LinodeAPIHandler, resourceID string) (*models.Resource, error) {
	account, err := processAccount(ctx, handler, resourceID)
	if err != nil {
		return nil, err
	}
	balance := fmt.Sprintf("%f", account.Balance)
	balanceUninvoiced := fmt.Sprintf("%f", account.BalanceUninvoiced)
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
	return &value, nil
}

func processAccount(ctx context.Context, handler *LinodeAPIHandler, resourceID string) (*model.Account, error) {
	var account *model.Account
	var resp *http.Response
	baseURL := "https://api.linode.com/v4/account"
	requestFunc := func(req *http.Request) (*http.Response, error) {
		var e error
		req, e = http.NewRequest("GET", baseURL, nil)
		if e != nil {
			return nil, e
		}
		resp, e = handler.Client.Do(req)
		if e = json.NewDecoder(resp.Body).Decode(account); e != nil {
			return nil, e
		}
		return resp, e
	}
	err := handler.DoRequest(ctx, &http.Request{}, requestFunc)
	if err != nil {
		return nil, err
	}
	return account, nil
}
