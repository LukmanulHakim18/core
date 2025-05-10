package http

import (
	"context"
	"net/http"
	"os"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

var (
	labelNames = []string{"app_name", "pod_name", "external_service_name", "method", "path", "status"}
)

type metricMiddleware struct {
	next                Middleware
	externalServiceName string
	appName             string
	namespace           string
	podName             string
}

func NewMetricMiddleware(externalServiceName string) Middleware {
	appName := os.Getenv("APP_NAME")
	if appName == "" {
		appName = "unknown service"
	}

	podName := os.Getenv("POD_NAME")
	if podName == "" {
		podName = "unknown pod"
	}

	namespace := os.Getenv("NAMESPACE")
	if namespace == "" {
		namespace = "default"
	}

	return &metricMiddleware{
		externalServiceName: externalServiceName,
		appName:             appName,
		namespace:           namespace,
		podName:             podName,
	}
}

func (m *metricMiddleware) Process(ctx context.Context, client *http.Client, req *http.Request) (res *http.Response, err error) {
	start := time.Now()

	res, err = m.next.Process(ctx, client, req)

	// Calculate latency and determine status
	latency := time.Since(start).Seconds()
	status := "error"
	if err == nil && res.StatusCode < 400 {
		status = "success"
	}

	label := prometheus.Labels{
		"app_name":              m.appName,
		"pod_name":              m.podName,
		"external_service_name": m.externalServiceName,
		"method":                req.Method,
		"path":                  req.URL.Path,
		"status":                status,
	}

	// Update prometheus metrics
	apiRequestTotal.With(label).Inc()
	apiRequestLatency.With(label).Observe(latency)

	return res, err
}

func (m *metricMiddleware) SetNext(next Middleware) {
	m.next = next
}

var (
	apiRequestTotal = promauto.NewCounterVec(
		prometheus.CounterOpts{
			Name: "external_api_requests_total",
			Help: "Total number of API requests processed",
		},
		labelNames,
	)

	apiRequestLatency = promauto.NewHistogramVec(
		prometheus.HistogramOpts{
			Name: "external_api_requests_latency_seconds",
			Help: "Latency of API requests in seconds",
		},
		labelNames,
	)
)
