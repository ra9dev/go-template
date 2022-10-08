package tracing

import (
	"context"

	"go.opentelemetry.io/otel/trace"
)

// StartGRPCTrace starts a new trace for a gRPC.
func StartGRPCTrace(ctx context.Context, name string) (context.Context, trace.Span) { //nolint:ireturn
	tracer := Tracer("grpc")

	return tracer.Start(ctx, name, trace.WithSpanKind(trace.SpanKindServer))
}

// SpanFromContext returns a span from the context.
func SpanFromContext(ctx context.Context, pkgName, methodName string) (context.Context, trace.Span) { //nolint:ireturn
	tracer := Tracer(pkgName)

	return tracer.Start(ctx, methodName, trace.WithSpanKind(trace.SpanKindInternal))
}
