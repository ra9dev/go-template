package tracing

import (
	"context"
	"net/http"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/codes"
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/trace"
)

// StartHTTPTrace starts a new trace for an HTTP request.
func StartHTTPTrace(ctx context.Context, name string, r *http.Request) (context.Context, *http.Request) {
	// Get the tracer from the context
	tracer := Tracer("http")
	// Start a new span
	ctx, span := tracer.Start(ctx, name, trace.WithSpanKind(trace.SpanKindServer))
	// Add the request headers to the span
	span.SetAttributes(
		attribute.String("http.method", r.Method),
		attribute.String("http.host", r.Host),
		attribute.String("http.target", r.URL.Path),
		attribute.String("http.scheme", r.URL.Scheme),
	)
	// Inject the trace headers into the request
	r = r.WithContext(otel.GetTextMapPropagator().Extract(ctx, propagation.HeaderCarrier(r.Header)))
	return ctx, r
}

// StartGRPCTrace starts a new trace for a gRPC
func StartGRPCTrace(ctx context.Context, name string) (context.Context, trace.Span) {
	// Get the tracer from the context
	tracer := Tracer("grpc")
	// Start a new span
	ctx, span := tracer.Start(ctx, name, trace.WithSpanKind(trace.SpanKindServer))
	return ctx, span
}

func SpanFromContext(ctx context.Context, pkgName, methodName string) (context.Context, trace.Span) {
	tracer := Tracer(pkgName)
	ctx, span := tracer.Start(ctx, methodName, trace.WithSpanKind(trace.SpanKindInternal))
	return ctx, span
}

func TraceError(span trace.Span, err error) {
	span.SetStatus(codes.Error, err.Error())
}
