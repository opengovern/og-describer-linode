package describer

import (
	"context"
	"encoding/json"
	"github.com/opengovern/og-describer-linode/pkg/sdk/models"
	"github.com/opengovern/og-describer-linode/provider/model"
	"net/http"
)

func GetObjectStorage(ctx context.Context, handler *LinodeAPIHandler, resourceID string) (*models.Resource, error) {
	objectStorage, err := processObjectStorage(ctx, handler)
	if err != nil {
		return nil, err
	}
	value := models.Resource{
		ID:   resourceID,
		Name: resourceID,
		Description: JSONAllFieldsMarshaller{
			Value: objectStorage,
		},
	}
	return &value, nil
}

func processObjectStorage(ctx context.Context, handler *LinodeAPIHandler) (*model.ObjectStorageDescription, error) {
	var objectStorage *model.ObjectStorageDescription
	var resp *http.Response
	baseURL := "https://api.linode.com/v4/object-storage/transfer"
	requestFunc := func(req *http.Request) (*http.Response, error) {
		var e error
		req, e = http.NewRequest("GET", baseURL, nil)
		if e != nil {
			return nil, e
		}
		resp, e = handler.Client.Do(req)
		if e = json.NewDecoder(resp.Body).Decode(objectStorage); e != nil {
			return nil, e
		}
		return resp, e
	}
	err := handler.DoRequest(ctx, &http.Request{}, requestFunc)
	if err != nil {
		return nil, err
	}
	return objectStorage, nil
}
