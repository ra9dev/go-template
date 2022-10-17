package grpc

import (
	"fmt"

	trace "go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc"
	"google.golang.org/grpc"
)

// Dial to connect to grpc.Server at target address and trace client calls
func Dial(target string, params Params, extraOpts ...grpc.DialOption) (*grpc.ClientConn, error) {
	si := trace.StreamClientInterceptor(params.ExtraOpts...)
	ui := trace.UnaryClientInterceptor(params.ExtraOpts...)

	opts := []grpc.DialOption{grpc.WithStreamInterceptor(si), grpc.WithUnaryInterceptor(ui)}
	opts = append(opts, extraOpts...)

	conn, err := grpc.Dial(target, opts...)
	if err != nil {
		return nil, fmt.Errorf("failed to grpc dial: %w", err)
	}

	return conn, nil
}
