package log

import "go.uber.org/zap"

func NewNoopLogger() Logger {
	return Logger{driver: zap.NewNop().Sugar()}
}
