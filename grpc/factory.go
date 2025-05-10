package grpc

import (
	"errors"
	"io"
	"time"

	"github.com/go-kit/kit/endpoint"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/sd"
	stdopentracing "github.com/opentracing/opentracing-go"
	"google.golang.org/grpc/credentials"

	ulog "github.com/LukmanulHakim18/core/log"

	"google.golang.org/grpc"
)

// ClientOption stores grpc client options
type ClientOption struct {
	//Timeout for circuit breaker
	Timeout time.Duration
	//Number of retry
	Retry int
	//Timeout for retry
	RetryTimeout time.Duration
	// ...
	MaxCallRecvMsgSize int
}

func grpcConnection(address string, creds credentials.TransportCredentials) (*grpc.ClientConn, error) {
	var conn *grpc.ClientConn
	var err error
	if creds == nil {
		conn, err = grpc.Dial(address, grpc.WithInsecure())
	} else {
		conn, err = grpc.Dial(address, grpc.WithTransportCredentials(creds))
	}
	if err != nil {
		return nil, err
	}
	return conn, nil
}

func grpcConnectionWithMaxCallRecvMsgSize(address string, creds credentials.TransportCredentials, maxCallRecvMsgSize int) (*grpc.ClientConn, error) {
	var conn *grpc.ClientConn
	var err error
	if creds == nil {
		conn, err = grpc.Dial(address, grpc.WithInsecure(), grpc.WithDefaultCallOptions(grpc.MaxCallRecvMsgSize(maxCallRecvMsgSize)))
	} else {
		conn, err = grpc.Dial(address, grpc.WithTransportCredentials(creds), grpc.WithDefaultCallOptions(grpc.MaxCallRecvMsgSize(maxCallRecvMsgSize)))
	}
	if err != nil {
		return nil, err
	}
	return conn, nil
}

// EndpointFactory returns endpoint factory
func EndpointFactory(makeEndpoint func(*grpc.ClientConn, time.Duration, stdopentracing.Tracer, log.Logger) endpoint.Endpoint, creds credentials.TransportCredentials, timeout time.Duration, tracer stdopentracing.Tracer, logger log.Logger) sd.Factory {
	return func(instance string) (endpoint.Endpoint, io.Closer, error) {

		if instance == "" {
			return nil, nil, errors.New("Empty instance")
		}

		conn, err := grpcConnection(instance, creds)
		if err != nil {
			logger.Log("host", instance, ulog.LogError, err.Error())
			return nil, nil, err
		}
		endpoint := makeEndpoint(conn, timeout, tracer, logger)

		return endpoint, conn, nil
	}
}

func EndpointFactoryWithMaxCallRecvMsgSize(makeEndpoint func(*grpc.ClientConn, time.Duration, stdopentracing.Tracer, log.Logger) endpoint.Endpoint, creds credentials.TransportCredentials, timeout time.Duration, tracer stdopentracing.Tracer, logger log.Logger, maxCallRecvMsgSize int) sd.Factory {
	return func(instance string) (endpoint.Endpoint, io.Closer, error) {

		if instance == "" {
			return nil, nil, errors.New("Empty instance")
		}

		conn, err := grpcConnectionWithMaxCallRecvMsgSize(instance, creds, maxCallRecvMsgSize)
		if err != nil {
			logger.Log("host", instance, ulog.LogError, err.Error())
			return nil, nil, err
		}
		endpoint := makeEndpoint(conn, timeout, tracer, logger)

		return endpoint, conn, nil
	}
}
