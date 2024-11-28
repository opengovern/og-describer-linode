package describer

import (
	"github.com/linode/linodego"
	"golang.org/x/time/rate"
	"time"
)

type LinodeAPIHandler struct {
	Client       linodego.Client
	RateLimiter  *rate.Limiter
	Semaphore    chan struct{}
	MaxRetries   int
	RetryBackoff time.Duration
}

func NewLinodeAPIHandler(client linodego.Client, rateLimit rate.Limit, burst int, maxConcurrency int, maxRetries int, retryBackoff time.Duration) *LinodeAPIHandler {
	return &LinodeAPIHandler{
		Client:       client,
		RateLimiter:  rate.NewLimiter(rateLimit, burst),
		Semaphore:    make(chan struct{}, maxConcurrency),
		MaxRetries:   maxRetries,
		RetryBackoff: retryBackoff,
	}
}
