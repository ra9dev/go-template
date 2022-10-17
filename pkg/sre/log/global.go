package log

import (
	"context"
	"sync"

	"go.uber.org/zap"
)

var (
	loggerRegistration sync.Once
	logger             = NewNoopLogger()
)

// RegisterLogger by constructing Logger via NewLogger, global Logger is disabled by default
func RegisterLogger(newLogger Logger) {
	loggerRegistration.Do(func() {
		logger = newLogger

		zap.ReplaceGlobals(logger.driver.Desugar())
	})
}

func Sync() error {
	return logger.Sync()
}

func NoContext() Writer[noContextLogger] {
	return logger.NoContext()
}

func With(key string, value any) Logger {
	return logger.With(key, value)
}

func Debug(ctx context.Context, args ...any) {
	logger.Debug(ctx, args...)
}

func Debugf(ctx context.Context, template string, args ...any) {
	logger.Debugf(ctx, template, args...)
}

func Info(ctx context.Context, args ...any) {
	logger.Info(ctx, args...)
}

func Infof(ctx context.Context, template string, args ...any) {
	logger.Infof(ctx, template, args...)
}

func Warn(ctx context.Context, args ...any) {
	logger.Warn(ctx, args...)
}

func Warnf(ctx context.Context, template string, args ...any) {
	logger.Warnf(ctx, template, args...)
}

func Error(ctx context.Context, args ...any) {
	logger.Error(ctx, args...)
}

func Errorf(ctx context.Context, template string, args ...any) {
	logger.Errorf(ctx, template, args...)
}

func Panic(ctx context.Context, args ...any) {
	logger.Panic(ctx, args...)
}

func Panicf(ctx context.Context, template string, args ...any) {
	logger.Panicf(ctx, template, args...)
}

func Fatal(ctx context.Context, args ...any) {
	logger.Fatal(ctx, args...)
}

func Fatalf(ctx context.Context, template string, args ...any) {
	logger.Fatalf(ctx, template, args...)
}
