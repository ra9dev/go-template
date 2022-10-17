package tracing

import (
	"context"
	"fmt"

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

	endpoint := jaeger.WithEndpoint(config.Endpoint)

	collection := jaeger.WithCollectorEndpoint(
		endpoint,
	)

	exp, err := jaeger.New(collection)
	if err != nil {
		return nil, fmt.Errorf("failed to create Jaeger exporter: %w", err)
	}

	res, err := resource.Merge(
		resource.Default(),
		resource.NewWithAttributes(
			semconv.SchemaURL,
			semconv.ServiceNameKey.String(config.ServiceName),
			semconv.ServiceVersionKey.String(config.ServiceVersion),
			semconv.DeploymentEnvironmentKey.String(config.Environment),
		),
	)
	if err != nil {
		return nil, fmt.Errorf("failed to create resource: %w", err)
	}

	provider := sdktrace.NewTracerProvider(
		sdktrace.WithBatcher(exp),
		sdktrace.WithResource(res),
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

type shutdownable interface {
	Shutdown(ctx context.Context) error
}

// Shutdown shuts down the tracing provider.
func (p Provider) Shutdown(ctx context.Context) error {
	if p.provider == nil {
		return nil
	}

	prv, ok := p.provider.(shutdownable)
	if !ok {
		return nil
	}

	if err := prv.Shutdown(ctx); err != nil {
		return fmt.Errorf("failed to shutdown tracing provider: %w", err)
	}

	return nil
}

// Tracer returns a tracer for the given name.
func Tracer(name string) trace.Tracer {
	return otel.GetTracerProvider().Tracer(name)
}
