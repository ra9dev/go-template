package main

import (
	"context"
	"fmt"

	"github.com/spf13/cobra"
	"go.uber.org/zap"

	"github.com/ra9dev/go-template/internal/config"
	"github.com/ra9dev/go-template/pkg/log"
	"github.com/ra9dev/go-template/pkg/shutdown"
)

func main() {
	rootCmd := &cobra.Command{
		Use:   "go-template",
		Short: "Main entry-point command for the application",
	}

	cfg, err := config.NewConfig()
	if err != nil {
		zap.S().Panicf("failed to prepare config: %w", err)
	}

	if err = setupLogger(cfg); err != nil {
		zap.S().Panic(err)
	}

	rootCmd.AddCommand(
		APIServerCMD(cfg),
	)

	defer func() {
		if shutdownErr := shutdown.Wait(); shutdownErr != nil {
			zap.S().Error(shutdownErr)

			return
		}

		zap.S().Infof("Shutdown completed in %.1f seconds", shutdown.Timeout().Seconds())
	}()

	if err = rootCmd.ExecuteContext(shutdown.Context()); err != nil {
		zap.S().Errorf("failed to execute root cmd: %v", err)

		return
	}
}

func setupLogger(cfg config.Config) error {
	_, err := log.NewLogger(cfg.LogLevel.ToZapAtomic())
	if err != nil {
		return fmt.Errorf("failed to prepare logger: %w", err)
	}

	shutdown.Add(func(_ context.Context) {
		_ = zap.L().Sync()
	})

	return nil
}
