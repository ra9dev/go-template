package log

import "context"

var _ Writer[noContextLogger] = (*noContextLogger)(nil)

type noContextLogger struct {
	logger Logger
}

func (n noContextLogger) With(key string, value any) noContextLogger {
	n.logger = n.logger.With(key, value)

	return n
}

func (n noContextLogger) Debug(args ...any) {
	n.logger.Debug(context.Background(), args...)
}

func (n noContextLogger) Debugf(template string, args ...any) {
	n.logger.Debugf(context.Background(), template, args...)
}

func (n noContextLogger) Info(args ...any) {
	n.logger.Info(context.Background(), args...)
}

func (n noContextLogger) Infof(template string, args ...any) {
	n.logger.Infof(context.Background(), template, args...)
}

func (n noContextLogger) Warn(args ...any) {
	n.logger.Warn(context.Background(), args...)
}

func (n noContextLogger) Warnf(template string, args ...any) {
	n.logger.Warnf(context.Background(), template, args...)
}

func (n noContextLogger) Error(args ...any) {
	n.logger.Error(context.Background(), args...)
}

func (n noContextLogger) Errorf(template string, args ...any) {
	n.logger.Errorf(context.Background(), template, args...)
}

func (n noContextLogger) Panic(args ...any) {
	n.logger.Panic(context.Background(), args...)
}

func (n noContextLogger) Panicf(template string, args ...any) {
	n.logger.Panicf(context.Background(), template, args...)
}

func (n noContextLogger) Fatal(args ...any) {
	n.logger.Fatal(context.Background(), args...)
}

func (n noContextLogger) Fatalf(template string, args ...any) {
	n.logger.Fatalf(context.Background(), template, args...)
}
