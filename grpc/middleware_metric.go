package grpc

import (
	"context"
	"os"
	"strings"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"google.golang.org/grpc"
)

var (
	labelNames = []string{"app_name", "pod_name", "external_service_name", "method", "path", "status"}
)

var (
	// gRPC request counter
	grpcRequests = promauto.NewCounterVec(
		prometheus.CounterOpts{
			Name: "external_grpc_requests_total",
			Help: "Total number of gRPC requests",
		},
		labelNames,
	)

	// gRPC request latency
	grpcLatency = promauto.NewHistogramVec(
		prometheus.HistogramOpts{
			Name:    "external_grpc_requests_latency_seconds",
			Help:    "Latency of gRPC requests",
			Buckets: prometheus.DefBuckets,
		},
		labelNames,
	)
)

type MetricInterceptor struct {
	externalServiceName string
	appName             string
	podName             string
}

func NewMetricInterceptor(externalServiceName string) *MetricInterceptor {
	appName := os.Getenv("APP_NAME")
	if appName == "" {
		appName = "unknown service"
	}

	podName := os.Getenv("POD_NAME")
	if podName == "" {
		podName = "unknown pod"
	}

	return &MetricInterceptor{
		externalServiceName: externalServiceName,
		appName:             appName,
		podName:             podName,
	}
}

func (m *MetricInterceptor) ObserveGRPCRequest(label prometheus.Labels, start time.Time) {
	// Update prometheus metrics
	grpcRequests.With(label).Inc()
	grpcLatency.With(label).Observe(time.Since(start).Seconds())
}

func (m *MetricInterceptor) UnaryClientInterceptor() grpc.UnaryClientInterceptor {
	return func(ctx context.Context, method string, req interface{}, reply interface{}, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
		start := time.Now()

		// Execute the gRPC call
		err := invoker(ctx, method, req, reply, cc, opts...)
		status := "success"
		if err != nil {
			status = "error"
		}

		label := prometheus.Labels{
			"app_name":              m.appName,
			"pod_name":              m.podName,
			"external_service_name": m.externalServiceName,
			"method":                extractMethod(method),
			"path":                  method,
			"status":                status,
		}

		m.ObserveGRPCRequest(label, start)

		return err
	}
}

func extractMethod(fullMethod string) string {
	parts := strings.Split(fullMethod, "/")
	if len(parts) > 2 {
		return parts[2] // The actual method name
	}
	return fullMethod
}

func (m *MetricInterceptor) StreamClientInterceptor() grpc.StreamClientInterceptor {
	return func(ctx context.Context, desc *grpc.StreamDesc, cc *grpc.ClientConn, method string, streamer grpc.Streamer, opts ...grpc.CallOption) (grpc.ClientStream, error) {
		start := time.Now()

		// Execute streaming call
		stream, err := streamer(ctx, desc, cc, method, opts...)

		status := "success"
		if err != nil {
			status = "error"
		}

		label := prometheus.Labels{
			"app_name":              m.appName,
			"pod_name":              m.podName,
			"external_service_name": m.externalServiceName,
			"method":                extractMethod(method),
			"path":                  method,
			"status":                status,
		}

		m.ObserveGRPCRequest(label, start)

		return stream, err
	}
}
