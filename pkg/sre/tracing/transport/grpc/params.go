package grpc

import (
	trace "go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc"
)

type (
	// SpanTag key-value pair
	SpanTag struct {
		Key   string
		Value any
	}

	// Params covers all possible tracing options
	Params struct {
		ExtraOpts []trace.Option
	}
)

// NewParams constructor with minimal required fields
func NewParams(extraOpts ...trace.Option) Params {
	return Params{
		ExtraOpts: extraOpts,
	}
}
