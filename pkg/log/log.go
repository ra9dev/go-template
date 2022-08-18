package log

import (
	"fmt"

	"go.uber.org/zap"
)

func NewLogger(level zap.AtomicLevel) (*zap.Logger, error) {
	cfg := zap.NewProductionConfig()
	cfg.Level = level

	logger, err := cfg.Build()
	if err != nil {
		return nil, fmt.Errorf("could not build logger: %w", err)
	}

	zap.ReplaceGlobals(logger)

	return logger, nil
}
