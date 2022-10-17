package log

import "context"

type (
	Writer[T any] interface {
		With(key string, value any) T
		Debug(args ...any)
		Debugf(template string, args ...any)
		Info(args ...any)
		Infof(template string, args ...any)
		Warn(args ...any)
		Warnf(template string, args ...any)
		Error(args ...any)
		Errorf(template string, args ...any)
		Panic(args ...any)
		Panicf(template string, args ...any)
		Fatal(args ...any)
		Fatalf(template string, args ...any)
	}

	ContextWriter[T any] interface {
		With(key string, value any) T
		Debug(ctx context.Context, args ...any)
		Debugf(ctx context.Context, template string, args ...any)
		Info(ctx context.Context, args ...any)
		Infof(ctx context.Context, template string, args ...any)
		Warn(ctx context.Context, args ...any)
		Warnf(ctx context.Context, template string, args ...any)
		Error(ctx context.Context, args ...any)
		Errorf(ctx context.Context, template string, args ...any)
		Panic(ctx context.Context, args ...any)
		Panicf(ctx context.Context, template string, args ...any)
		Fatal(ctx context.Context, args ...any)
		Fatalf(ctx context.Context, template string, args ...any)
	}
)
