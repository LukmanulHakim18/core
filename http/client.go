package http

import (
	"bytes"
	"context"
	"crypto/tls"
	"go.elastic.co/apm/module/apmhttp/v2"
	"net/http"
	"net/url"
	"time"
)

type Endpoint struct {
	Method        string
	Endpoint      string
	unescapeQuery bool
	Params        url.Values
}

func NewEndpoint(path string, qParam url.Values, method string) *Endpoint {
	return &Endpoint{
		Method:   method,
		Endpoint: path,
		Params:   qParam,
	}
}

func (e *Endpoint) UnescapeQuery() {
	e.unescapeQuery = true
}

func (e *Endpoint) buildUrl(baseUrl string) *url.URL {
	fullUrl, _ := url.Parse(baseUrl + e.Endpoint)
	if e.Params != nil {
		fullUrl.RawQuery = e.Params.Encode()
		if e.unescapeQuery {
			fullUrl.RawQuery, _ = url.QueryUnescape(fullUrl.RawQuery)
		}
	}
	return fullUrl
}

type Client struct {
	baseURL         string
	client          *http.Client
	middlewares     []Middleware
	startMiddleware Middleware
}

func NewClient(baseUrl string) *Client {
	return &Client{
		baseURL: baseUrl,
		client:  apmhttp.WrapClient(http.DefaultClient),
	}
}

func (c *Client) SetTimeout(timeout time.Duration) {
	c.client.Timeout = timeout
}

func (c *Client) SetTransport(transport http.RoundTripper) {
	c.client.Transport = transport
}

func NewInsecureClient(baseUrl string) *Client {
	tp := http.DefaultTransport.(*http.Transport).Clone()
	tp.TLSClientConfig = &tls.Config{InsecureSkipVerify: true}
	client := &http.Client{Transport: tp}
	return &Client{
		baseURL: baseUrl,
		client:  apmhttp.WrapClient(client),
	}
}

func (c *Client) Exec(ctx context.Context, ep *Endpoint, header http.Header, body []byte, option ...ClientV2Option) (*http.Response, error) {
	req, err := http.NewRequest(ep.Method, ep.buildUrl(c.baseURL).String(), bytes.NewBuffer(body))
	if err != nil {
		return nil, err
	}

	// Set Host if find in header
	hostVal := header.Get("HOST")
	if hostVal != "" {
		req.Host = hostVal
	}

	// Set Header.
	req.Header = header

	for _, cvo := range option {
		cvo(c)
	}

	if c.startMiddleware == nil {
		c.setStartMiddleware()
	}

	return c.startMiddleware.Process(ctx, c.client, req)
}

func (c *Client) Use(m Middleware) {
	c.middlewares = append(c.middlewares, m)
}

func (c *Client) setStartMiddleware() {
	if c.startMiddleware != nil {
		return
	}

	if len(c.middlewares) == 0 {
		c.startMiddleware = &Runner{}
		return
	}

	for i := 0; i < len(c.middlewares); i++ {
		if i < len(c.middlewares)-1 {
			c.middlewares[i].SetNext(c.middlewares[i+1])
		} else {
			c.middlewares[i].SetNext(&Runner{})
		}
	}
	c.startMiddleware = c.middlewares[0]
}

type ClientV2Option func(c *Client) *Client
