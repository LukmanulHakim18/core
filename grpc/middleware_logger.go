package grpc

import (
	"context"
	"time"

	"github.com/LukmanulHakim18/core/logger"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

// LoggerInterceptor is a gRPC client interceptor for logging requests and responses.
type LoggerInterceptor struct {
	logger *logger.Logger
}

// NewLoggerInterceptor creates a new LoggerInterceptor instance.
func NewLoggerInterceptor(logger *logger.Logger) *LoggerInterceptor {
	return &LoggerInterceptor{
		logger: logger,
	}
}

// UnaryClientInterceptor logs details for unary gRPC client calls.
func (l *LoggerInterceptor) UnaryClientInterceptor() grpc.UnaryClientInterceptor {
	return func(ctx context.Context, method string, req interface{}, reply interface{}, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
		startTime := time.Now()

		// Extract metadata
		md, _ := metadata.FromOutgoingContext(ctx)

		// Log request details
		l.logger.InfoWithContext(ctx, "Sending gRPC request",
			logger.ConvertMapToFields(map[string]interface{}{
				"method":    method,
				"metadata":  md,
				"req_body":  req,
				"grpc_type": "unary",
			})...,
		)

		// Execute the gRPC call
		err := invoker(ctx, method, req, reply, cc, opts...)
		duration := time.Since(startTime)

		// Check for errors
		if err != nil {
			l.logger.ErrorWithContext(ctx, "gRPC request failed",
				logger.ConvertMapToFields(map[string]interface{}{
					"method":      method,
					"duration_ms": duration.Milliseconds(),
					"error":       err.Error(),
					"grpc_type":   "unary",
				})...,
			)
			return err
		}

		// Log response details
		l.logger.InfoWithContext(ctx, "Received gRPC response",
			logger.ConvertMapToFields(map[string]interface{}{
				"method":      method,
				"status":      "OK",
				"duration_ms": duration.Milliseconds(),
				"grpc_type":   "unary",
				"res_body":    reply,
			})...,
		)

		return nil
	}
}

// StreamClientInterceptor logs details for streaming gRPC client calls.
func (l *LoggerInterceptor) StreamClientInterceptor() grpc.StreamClientInterceptor {
	return func(ctx context.Context, desc *grpc.StreamDesc, cc *grpc.ClientConn, method string, streamer grpc.Streamer, opts ...grpc.CallOption) (grpc.ClientStream, error) {
		startTime := time.Now()

		// Extract metadata
		md, _ := metadata.FromOutgoingContext(ctx)

		// Log request details
		l.logger.InfoWithContext(ctx, "Starting gRPC stream",
			logger.ConvertMapToFields(map[string]interface{}{
				"method":    method,
				"metadata":  md,
				"grpc_type": "stream",
			})...,
		)

		// Execute streaming call
		stream, err := streamer(ctx, desc, cc, method, opts...)
		duration := time.Since(startTime)

		// Check for errors
		if err != nil {
			l.logger.ErrorWithContext(ctx, "gRPC stream failed",
				logger.ConvertMapToFields(map[string]interface{}{
					"method":      method,
					"duration_ms": duration.Milliseconds(),
					"error":       err.Error(),
					"grpc_type":   "stream",
				})...,
			)
			return nil, err
		}

		// Log successful stream initialization
		l.logger.InfoWithContext(ctx, "gRPC stream established",
			logger.ConvertMapToFields(map[string]interface{}{
				"method":      method,
				"status":      "OK",
				"duration_ms": duration.Milliseconds(),
				"grpc_type":   "stream",
			})...,
		)

		return stream, nil
	}
}
