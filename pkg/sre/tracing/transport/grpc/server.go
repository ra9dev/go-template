package grpc

import (
	trace "go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc"
	"google.golang.org/grpc"
)

// NewServer constructor with interceptors for tracing
func NewServer(params Params, extraOpts ...grpc.ServerOption) *grpc.Server {
	si := trace.StreamServerInterceptor(params.ExtraOpts...)
	ui := trace.UnaryServerInterceptor(params.ExtraOpts...)

	// TODO put request info interceptor here as well
	opts := []grpc.ServerOption{grpc.StreamInterceptor(si), grpc.UnaryInterceptor(ui)}
	opts = append(opts, extraOpts...)

	return grpc.NewServer(opts...)
}
