package redis

import (
	"context"
	"fmt"
	"os"
	"strings"
	"time"

	redis "github.com/go-redis/redis/v8"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

type ctxKey string

const startTimeKey ctxKey = "startTime"

var (
	labelNames = []string{"app_name", "pod_name", "external_service_name", "command", "full_command", "status"}
)

// metricHook implements redis.Hook
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
		externalServiceName: "redis",
		appName:             appName,
		podName:             podName,
	}
}

// BeforeProcess is called before a redis command is executed
func (m *metricHook) BeforeProcess(ctx context.Context, cmd redis.Cmder) (context.Context, error) {
	startTime := time.Now()
	return context.WithValue(ctx, startTimeKey, startTime), nil
}

// AfterProcess is called after a redis command is executed
func (m *metricHook) AfterProcess(ctx context.Context, cmd redis.Cmder) error {
	startTime, _ := ctx.Value(startTimeKey).(time.Time)
	duration := time.Since(startTime).Seconds()

	// Get the Redis command (e.g., "SET", "GET", "DEL")
	command := cmd.Name()

	// Get the full Redis command as a string
	fullCommand := getFullRedisCommand(cmd)

	status := "success"
	if cmd.Err() != nil {
		status = "error"
	}

	label := prometheus.Labels{
		"app_name":              m.appName,
		"pod_name":              m.podName,
		"external_service_name": m.externalServiceName,
		"command":               command,
		"full_command":          fullCommand,
		"status":                status,
	}

	// Update prometheus metrics
	redisRequestTotal.With(label).Inc()
	redisRequestLatency.With(label).Observe(duration)

	return nil
}

// BeforeProcessPipeline is called before executing multiple commands in a pipeline
func (m *metricHook) BeforeProcessPipeline(ctx context.Context, cmds []redis.Cmder) (context.Context, error) {
	startTime := time.Now()
	return context.WithValue(ctx, startTimeKey, startTime), nil
}

// AfterProcessPipeline is called after executing multiple commands in a pipeline
func (m *metricHook) AfterProcessPipeline(ctx context.Context, cmds []redis.Cmder) error {
	startTime, _ := ctx.Value(startTimeKey).(time.Time)
	duration := time.Since(startTime).Seconds()

	// Process each command in the pipeline
	for _, cmd := range cmds {
		command := cmd.Name()
		fullCommand := getFullRedisCommand(cmd)

		status := "success"
		if cmd.Err() != nil {
			status = "error"
		}

		label := prometheus.Labels{
			"app_name":              m.appName,
			"pod_name":              m.podName,
			"external_service_name": m.externalServiceName,
			"command":               command,
			"full_command":          fullCommand,
			"status":                status,
		}

		// Update prometheus metrics
		redisRequestTotal.With(label).Inc()
		redisRequestLatency.With(label).Observe(duration)
	}

	return nil
}

var (
	redisRequestTotal = promauto.NewCounterVec(
		prometheus.CounterOpts{
			Name: "external_redis_requests_total",
			Help: "Total number of Redis requests processed",
		},
		labelNames,
	)

	redisRequestLatency = promauto.NewHistogramVec(
		prometheus.HistogramOpts{
			Name: "external_redis_requests_latency_seconds",
			Help: "Latency of Redis request in seconds",
		},
		labelNames,
	)
)

// getFullRedisCommand converts Redis Cmder into a full command string
func getFullRedisCommand(cmd redis.Cmder) string {
	args := cmd.Args()
	commandStrings := make([]string, len(args))

	for i, arg := range args {
		commandStrings[i] = fmt.Sprintf("%v", arg)
	}

	return strings.Join(commandStrings, " ")
}
