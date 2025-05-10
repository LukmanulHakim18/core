package http

import (
	"context"
	"net/http"
)

type Middleware interface {
	Process(ctx context.Context, client *http.Client, req *http.Request) (res *http.Response, err error)
	SetNext(Middleware)
}
