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

func ListImages(ctx context.Context, handler *LinodeAPIHandler, stream *models.StreamSender) ([]models.Resource, error) {
	var wg sync.WaitGroup
	linodeChan := make(chan models.Resource)
	go func() {
		processImages(ctx, handler, linodeChan, &wg)
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

func GetImage(ctx context.Context, handler *LinodeAPIHandler, resourceID string) (*models.Resource, error) {
	image, err := processImage(ctx, handler, resourceID)
	if err != nil {
		return nil, err
	}
	value := models.Resource{
		ID:   image.ID,
		Name: image.Label,
		Description: JSONAllFieldsMarshaller{
			Value: image,
		},
	}
	return &value, nil
}

func processImages(ctx context.Context, handler *LinodeAPIHandler, openaiChan chan<- models.Resource, wg *sync.WaitGroup) {
	var images []model.ImageDescription
	var imageListResponse *model.ImageListResponse
	var resp *http.Response
	baseURL := "https://api.linode.com/v4/images"
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
			if e = json.NewDecoder(resp.Body).Decode(&imageListResponse); e != nil {
				return nil, e
			}
			images = append(images, imageListResponse.Data...)
			if imageListResponse.Page == imageListResponse.Pages {
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
	for _, image := range images {
		wg.Add(1)
		go func(image model.ImageDescription) {
			defer wg.Done()
			value := models.Resource{
				ID:   image.ID,
				Name: image.Label,
				Description: JSONAllFieldsMarshaller{
					Value: image,
				},
			}
			openaiChan <- value
		}(image)
	}
}

func processImage(ctx context.Context, handler *LinodeAPIHandler, resourceID string) (*model.ImageDescription, error) {
	var image *model.ImageDescription
	var resp *http.Response
	baseURL := "https://api.linode.com/v4/images/"
	requestFunc := func(req *http.Request) (*http.Response, error) {
		var e error
		finalURL := fmt.Sprintf("%s%s", baseURL, resourceID)
		req, e = http.NewRequest("GET", finalURL, nil)
		if e != nil {
			return nil, e
		}
		resp, e = handler.Client.Do(req)
		if e = json.NewDecoder(resp.Body).Decode(image); e != nil {
			return nil, e
		}
		return resp, e
	}
	err := handler.DoRequest(ctx, &http.Request{}, requestFunc)
	if err != nil {
		return nil, err
	}
	return image, nil
}
