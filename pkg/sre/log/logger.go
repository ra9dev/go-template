package log

import (
	"context"
	"fmt"

	"go.opentelemetry.io/otel/trace"
	"go.uber.org/zap"

	"github.com/ra9dev/go-template/pkg/sre"
)

var _ ContextWriter[Logger] = (*Logger)(nil)

type Logger struct {
	driver *zap.SugaredLogger
}

func NewLogger(params Params) (Logger, error) {
	params = params.withDefault()

	var cfg zap.Config

	switch params.Env {
	case sre.EnvLocal, sre.EnvDev:
		cfg = zap.NewDevelopmentConfig()
	default:
		cfg = zap.NewProductionConfig()
	}

	cfg.Level = params.Level.ToZapAtomic()

	driver, err := cfg.Build()
	if err != nil {
		return Logger{}, fmt.Errorf("could not build logger: %w", err)
	}

	return Logger{
		driver: driver.Sugar(),
	}, nil
}

func (l Logger) Sync() error {
	return l.driver.Sync() // nolint:wrapcheck
}

func (l Logger) NoContext() Writer[noContextLogger] {
	return noContextLogger{logger: l}
}

func (l Logger) With(key string, value any) Logger {
	l.driver = l.driver.With(key, value)

	return l
}

func (l Logger) withKeys(ctx context.Context) Logger {
	loggerWithKeys := l

	if span := trace.SpanFromContext(ctx); span.IsRecording() {
		loggerWithKeys = loggerWithKeys.
			With(sre.KeyTraceID.String(), span.SpanContext().TraceID()).
			With(sre.KeySpanID.String(), span.SpanContext().SpanID())
	}

	if requestID := ctx.Value(sre.KeyRequestID); requestID != nil {
		loggerWithKeys = loggerWithKeys.With(sre.KeyRequestID.String(), requestID)
	}

	return loggerWithKeys
}

func (l Logger) Debug(ctx context.Context, args ...any) {
	l.withKeys(ctx).driver.Debug(args...)
}

func (l Logger) Debugf(ctx context.Context, template string, args ...any) {
	l.withKeys(ctx).driver.Debugf(template, args...)
}

func (l Logger) Info(ctx context.Context, args ...any) {
	l.withKeys(ctx).driver.Info(args...)
}

func (l Logger) Infof(ctx context.Context, template string, args ...any) {
	l.withKeys(ctx).driver.Infof(template, args...)
}

func (l Logger) Warn(ctx context.Context, args ...any) {
	l.withKeys(ctx).driver.Warn(args...)
}

func (l Logger) Warnf(ctx context.Context, template string, args ...any) {
	l.withKeys(ctx).driver.Warnf(template, args...)
}

func (l Logger) Error(ctx context.Context, args ...any) {
	l.withKeys(ctx).driver.Error(args...)
}

func (l Logger) Errorf(ctx context.Context, template string, args ...any) {
	l.withKeys(ctx).driver.Errorf(template, args...)
}

func (l Logger) Panic(ctx context.Context, args ...any) {
	l.withKeys(ctx).driver.Panic(args...)
}

func (l Logger) Panicf(ctx context.Context, template string, args ...any) {
	l.withKeys(ctx).driver.Panicf(template, args...)
}

func (l Logger) Fatal(ctx context.Context, args ...any) {
	l.withKeys(ctx).driver.Fatal(args...)
}

func (l Logger) Fatalf(ctx context.Context, template string, args ...any) {
	l.withKeys(ctx).driver.Fatalf(template, args...)
}
