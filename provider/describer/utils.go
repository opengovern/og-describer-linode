package describer

import (
	"context"
	"errors"
	"fmt"
	"golang.org/x/time/rate"
	"net/http"
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
		if resp == nil {
			return err
		}
		// Set rate limiter new value
		var resetDuration int
		if resp != nil {
			retryAfter := resp.Header.Get("Retry-After")
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
