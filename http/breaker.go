package http

import (
	"context"
	"net/http"

	"github.com/LukmanulHakim18/core/microservice"
)

var cb *microservice.Breaker

type Breaker struct {
	next     Middleware
	CBConfig microservice.BreakerConfig
}

func (b *Breaker) Process(ctx context.Context, client *http.Client, req *http.Request) (res *http.Response, err error) {
	cb = microservice.NewCircuitBreaker(&b.CBConfig)
	cbRes, cbErr := cb.Execute(func() (interface{}, error) {
		res, err := b.next.Process(ctx, client, req)
		return res, err
	})
	returned, ok := cbRes.(*http.Response)
	if ok {
		return returned, cbErr
	}
	return nil, cbErr
}

func (b *Breaker) SetNext(next Middleware) {
	b.next = next
}
