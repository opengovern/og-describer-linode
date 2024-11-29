package describer

import (
	"context"
	"encoding/json"
	"fmt"
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
	var objectStorage model.ObjectStorageDescription
	var resp *http.Response
	baseURL := "https://api.linode.com/v4/object-storage/transfer"

	req, err := http.NewRequest("GET", baseURL, nil)
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

		if e = json.NewDecoder(resp.Body).Decode(&objectStorage); e != nil {
			return nil, fmt.Errorf("failed to decode response: %w", e)
		}
		return resp, e
	}

	err = handler.DoRequest(ctx, req, requestFunc)
	if err != nil {
		return nil, fmt.Errorf("error during request handling: %w", err)
	}
	return &objectStorage, nil
}
