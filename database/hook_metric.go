package database

import (
	"context"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

var (
	labelNames = []string{"app_name", "pod_name", "external_service_name", "query", "status"}
)

type metricHook struct {
	externalServiceName string
	appName             string
	podName             string
}

func newMetricHook() *metricHook {
	appName := os.Getenv("APP_NAME")
	if appName == "" {
		appName = "unknown service"
	}

	podName := os.Getenv("POD_NAME")
	if podName == "" {
		podName = "unknown pod"
	}

	return &metricHook{
		externalServiceName: "database",
		appName:             appName,
		podName:             podName,
	}
}

// AfterProcess is called after a DB command is executed
func (m *metricHook) AfterProcess(ctx context.Context, startTime time.Time, err error, query string, args ...any) {
	duration := time.Since(startTime).Seconds()

	queryStr := buildQuery(ctx, query, args...)

	status := "success"
	if err != nil {
		status = "error"
	}

	label := prometheus.Labels{
		"app_name":              m.appName,
		"pod_name":              m.podName,
		"external_service_name": m.externalServiceName,
		"query":                 queryStr,
		"status":                status,
	}

	// Update prometheus metrics
	dbRequestTotal.With(label).Inc()
	dbRequestLatency.With(label).Observe(duration)
}

var (
	dbRequestTotal = promauto.NewCounterVec(
		prometheus.CounterOpts{
			Name: "external_db_requests_total",
			Help: "Total number of DB requests processed",
		},
		labelNames,
	)

	dbRequestLatency = promauto.NewHistogramVec(
		prometheus.HistogramOpts{
			Name: "external_db_requests_latency_seconds",
			Help: "Latency of DB request in seconds",
		},
		labelNames,
	)
)

// buildQuery replaces PostgreSQL-style ($1, $2, ...) placeholders with actual values.
func buildQuery(ctx context.Context, query string, args ...interface{}) string {
	for i, arg := range args {
		placeholder := fmt.Sprintf("$%d", i+1)
		replacement := fmt.Sprintf("'%v'", arg) // Enclose values in single quotes for safety.
		query = strings.Replace(query, placeholder, replacement, 1)
	}
	return query
}
