package describer

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/opengovern/og-describer-linode/provider/model"
	"golang.org/x/time/rate"
	"net/http"
	"net/url"
	"strconv"
	"time"
)

type LinodeAPIHandler struct {
	Client       *http.Client
	Token        string
	RateLimiter  *rate.Limiter
	Semaphore    chan struct{}
	MaxRetries   int
	RetryBackoff time.Duration
}

func NewLinodeAPIHandler(token string, rateLimit rate.Limit, burst int, maxConcurrency int, maxRetries int, retryBackoff time.Duration) *LinodeAPIHandler {
	return &LinodeAPIHandler{
		Client:       http.DefaultClient,
		Token:        token,
		RateLimiter:  rate.NewLimiter(rateLimit, burst),
		Semaphore:    make(chan struct{}, maxConcurrency),
		MaxRetries:   maxRetries,
		RetryBackoff: retryBackoff,
	}
}

func getLinodeInstances(ctx context.Context, handler *LinodeAPIHandler) ([]model.LinodeDescription, error) {
	var linodeInstances []model.LinodeDescription
	var linodeListResponse *model.LinodeListResponse
	var resp *http.Response
	baseURL := "https://api.linode.com/v4/linode/instances"
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
			if e = json.NewDecoder(resp.Body).Decode(&linodeListResponse); e != nil {
				return nil, e
			}
			linodeInstances = append(linodeInstances, linodeListResponse.Data...)
			if linodeListResponse.Page == linodeListResponse.Pages {
				break
			}
			page += 1
		}
		return resp, e
	}
	err := handler.DoRequest(ctx, &http.Request{}, requestFunc)
	if err != nil {
		return nil, err
	}
	return linodeInstances, nil
}

// DoRequest executes the linode API request with rate limiting, retries, and concurrency control.
func (h *LinodeAPIHandler) DoRequest(ctx context.Context, req *http.Request, requestFunc func(req *http.Request) (*http.Response, error)) error {
	h.Semaphore <- struct{}{}
	defer func() { <-h.Semaphore }()
	var resp *http.Response
	var err error
	for attempt := 0; attempt <= h.MaxRetries; attempt++ {
		// Wait based on rate limiter
		if err = h.RateLimiter.Wait(ctx); err != nil {
			return err
		}
		// Set request headers
		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", h.Token))
		// Execute the request function
		resp, err = requestFunc(req)
		if err == nil {
			return nil
		}
		// Set rate limiter new value
		retryAfter := resp.Header.Get("Retry-After")
		var resetDuration int
		if retryAfter != "" {
			resetDuration, _ = strconv.Atoi(retryAfter)
		}
		var remainRequests int
		remainRequestsStr := resp.Header.Get("X-RateLimit-Remaining")
		if remainRequestsStr != "" {
			remainRequests, err = strconv.Atoi(remainRequestsStr)
			if err == nil && resetDuration > 0 {
				h.RateLimiter = rate.NewLimiter(rate.Every(time.Duration(resetDuration)/time.Duration(remainRequests)), 1)
			}
		}
		// Handle rate limit errors
		if resp != nil && resp.StatusCode == http.StatusTooManyRequests {
			if resetDuration > 0 {
				time.Sleep(time.Duration(resetDuration))
				continue
			}
			// Exponential backoff if headers are missing
			backoff := h.RetryBackoff * (1 << attempt)
			time.Sleep(backoff)
			continue
		}
		// Handle temporary network errors
		if isTemporary(err) {
			backoff := h.RetryBackoff * (1 << attempt)
			time.Sleep(backoff)
			continue
		}
		break
	}
	return err
}

// isTemporary checks if an error is temporary.
func isTemporary(err error) bool {
	if err == nil {
		return false
	}
	var netErr interface{ Temporary() bool }
	if errors.As(err, &netErr) {
		return netErr.Temporary()
	}
	return false
}
