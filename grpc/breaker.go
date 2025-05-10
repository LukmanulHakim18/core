package grpc

import (
	"github.com/LukmanulHakim18/core/microservice"
	"golang.org/x/net/context"
	libGrpc "google.golang.org/grpc"
)

func BreakerClientUnaryInterceptor(cb *microservice.Breaker) libGrpc.UnaryClientInterceptor {
	return func(
		ctx context.Context,
		method string,
		req, reply interface{},
		cc *libGrpc.ClientConn,
		invoker libGrpc.UnaryInvoker,
		opts ...libGrpc.CallOption,
	) error {
		_, err := cb.Execute(func() (interface{}, error) {
			err := invoker(ctx, method, req, reply, cc, opts...)
			return nil, err
		})
		return err
	}
}
