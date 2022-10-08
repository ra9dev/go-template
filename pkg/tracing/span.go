package tracing

import (
	"context"

	"go.opentelemetry.io/otel/trace"
)

// SpanFromContext returns a span from the context.
func SpanFromContext(ctx context.Context, pkgName, methodName string) (context.Context, trace.Span) { //nolint:ireturn
	tracer := Tracer(pkgName)

	return tracer.Start(ctx, methodName, trace.WithSpanKind(trace.SpanKindInternal))
}
