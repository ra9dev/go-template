package tracing

import (
	"context"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/jaeger"
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/sdk/resource"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.12.0"
	"go.opentelemetry.io/otel/trace"
)

type Config struct {
	Endpoint       string
	ServiceName    string
	ServiceVersion string
	Environment    string
	// Set this to `false` if you want to disable tracing completely.
	Enabled bool
}

// Provider is a wrapper around the OpenTelemetry tracer provider.
type Provider struct {
	provider trace.TracerProvider
}

// NewProvider creates a new tracing provider.
func NewProvider(config Config) (*Provider, error) {
	if !config.Enabled {
		return &Provider{
			provider: trace.NewNoopTracerProvider(),
		}, nil
	}
	exp, err := jaeger.New(jaeger.WithCollectorEndpoint(jaeger.WithEndpoint(config.Endpoint)))
	if err != nil {
		return nil, err
	}

	r, err := resource.Merge(
		resource.Default(),
		resource.NewWithAttributes(
			semconv.SchemaURL,
			semconv.ServiceNameKey.String(config.ServiceName),
			semconv.ServiceVersionKey.String(config.ServiceVersion),
			semconv.DeploymentEnvironmentKey.String(config.Environment),
		),
	)

	if err != nil {
		return nil, err
	}

	provider := sdktrace.NewTracerProvider(
		sdktrace.WithBatcher(exp),
		sdktrace.WithResource(r),
	)

	otel.SetTracerProvider(provider)
	otel.SetTextMapPropagator(
		propagation.NewCompositeTextMapPropagator(
			propagation.TraceContext{},
			propagation.Baggage{},
		),
	)

	return &Provider{
		provider: provider,
	}, nil
}

// Shutdown shuts down the tracing provider.
func (p Provider) Shutdown(ctx context.Context) {
	if p.provider == nil {
		return
	}

	if prv, ok := p.provider.(*sdktrace.TracerProvider); ok {
		prv.Shutdown(ctx)
	}
}

// Tracer returns a tracer for the given name.
func Tracer(name string) trace.Tracer {
	return otel.GetTracerProvider().Tracer(name)
}

// Inject injects the trace headers into the given context.
func Inject(ctx context.Context, carrier propagation.TextMapCarrier) {
	otel.GetTextMapPropagator().Inject(ctx, carrier)
}
