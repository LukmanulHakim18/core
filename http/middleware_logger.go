package http

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"time"

	"github.com/LukmanulHakim18/core/logger"
)

type loggerMiddleware struct {
	next   Middleware
	logger *logger.Logger
}

func NewLoggerMiddleware(logger *logger.Logger) Middleware {
	return &loggerMiddleware{
		logger: logger,
	}
}

func (l *loggerMiddleware) Process(ctx context.Context, client *http.Client, req *http.Request) (*http.Response, error) {
	startTime := time.Now()

	// Prepare request headers as a map
	reqHeaders := make(map[string][]string)
	for k, v := range req.Header {
		reqHeaders[k] = v
	}

	// Log request details
	l.logger.InfoWithContext(ctx, "Sending HTTP request",
		logger.ConvertMapToFields(map[string]interface{}{
			"method":     req.Method,
			"url":        req.URL.String(),
			"host":       req.Host,
			"user_agent": req.UserAgent(),
			"req_header": reqHeaders,
			"req_body":   truncateString(readRequestBody(req), 10000),
		})...,
	)

	// Execute the request
	res, err := l.next.Process(ctx, client, req)
	duration := time.Since(startTime)

	// Check for errors
	if err != nil {
		l.logger.ErrorWithContext(ctx, "HTTP request failed",
			logger.ConvertMapToFields(map[string]interface{}{
				"method":      req.Method,
				"url":         req.URL.String(),
				"duration_ms": duration.Milliseconds(),
				"error":       err.Error(),
			})...,
		)
		return res, err
	}

	// Prepare response headers as a map
	resHeaders := make(map[string][]string)
	if res != nil {
		for k, v := range res.Header {
			resHeaders[k] = v
		}
	}

	// Log response details
	logRespFields := logger.ConvertMapToFields(map[string]interface{}{
		"method":      req.Method,
		"url":         req.URL.String(),
		"status_code": res.StatusCode,
		"status":      res.Status,
		"duration_ms": duration.Milliseconds(),
		"res_header":  resHeaders,
		"res_body":    truncateString(readResponseBody(res), 10000),
	})

	if res.StatusCode >= 200 && res.StatusCode < 300 {
		l.logger.InfoWithContext(ctx, "Received HTTP response", logRespFields...)
	} else {
		l.logger.ErrorWithContext(ctx, "HTTP request returned non-OK status", logRespFields...)
	}

	return res, nil
}

func (l *loggerMiddleware) SetNext(next Middleware) {
	l.next = next
}

// Helper function to truncate long strings
func truncateString(input string, maxLength int) string {
	if len(input) > maxLength {
		return input[:maxLength] + "..." // Add ellipsis to indicate truncation
	}
	return input
}

// Helper function to read and reset the request body
func readRequestBody(req *http.Request) string {
	if req.Body == nil {
		return ""
	}

	body, _ := io.ReadAll(req.Body)
	req.Body = io.NopCloser(bytes.NewBuffer(body)) // Reset body for actual request processing
	return string(body)
}

// Helper function to read and reset the response body
func readResponseBody(res *http.Response) string {
	if res.Body == nil {
		return ""
	}

	body, _ := io.ReadAll(res.Body)
	res.Body = io.NopCloser(bytes.NewBuffer(body)) // Reset body for further processing if needed
	return string(body)
}
