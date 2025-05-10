package http

import (
	"context"
	"net/http"

	"golang.org/x/net/context/ctxhttp"
)

type Runner struct{}

func (r *Runner) Process(ctx context.Context, client *http.Client, req *http.Request) (res *http.Response, err error) {
	return ctxhttp.Do(ctx, client, req)
}

func (r *Runner) SetNext(next Middleware) {
	// this middleware should not have any next
}
