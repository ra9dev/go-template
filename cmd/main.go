package main

import (
	"context"
	"fmt"

	"github.com/ra9dev/shutdown"
	"github.com/spf13/cobra"
	"go.uber.org/zap"

	"github.com/ra9dev/go-template/internal/config"
	"github.com/ra9dev/go-template/pkg/log"
	"github.com/ra9dev/go-template/pkg/tracing"
)

func main() {
	rootCmd := &cobra.Command{
		Use:   "go-template",
		Short: "Main entry-point command for the application",
	}

	cfg, err := config.NewConfig()
	if err != nil {
		zap.S().Fatalf("failed to prepare config: %v", err)
	}

	if err = setupLogger(cfg); err != nil {
		zap.S().Fatal(err)
	}

	if err = setupTracing(cfg); err != nil {
		zap.S().Fatal(err)
	}

	rootCmd.AddCommand(
		APIServerCMD(cfg),
	)

	defer gracefulShutdown()

	if err = rootCmd.ExecuteContext(shutdown.Context()); err != nil {
		zap.S().Errorf("failed to execute root cmd: %v", err)

		return
	}
}

func setupTracing(cfg config.Config) error {
	provider, err := tracing.NewProvider(tracing.Config{
		ServiceName:    config.ServiceName,
		ServiceVersion: config.ServiceVersion,
		Environment:    cfg.Env,
		Endpoint:       cfg.Tracing.Endpoint,
		Enabled:        cfg.Tracing.Enabled,
	})
	if err != nil {
		return fmt.Errorf("failed to prepare tracing provider: %w", err)
	}

	shutdown.MustAdd("tracing", func(ctx context.Context) {
		zap.S().Info("Shutting down tracing provider")

		if err = provider.Shutdown(ctx); err != nil {
			zap.S().Error(err)

			return
		}

		zap.S().Info("Tracing provider shutdown succeeded!")
	})

	return nil
}

func setupLogger(cfg config.Config) error {
	_, err := log.NewLogger(cfg.LogLevel.ToZapAtomic())
	if err != nil {
		return fmt.Errorf("failed to prepare logger: %w", err)
	}

	shutdown.MustAdd("logger", func(_ context.Context) {
		zap.S().Infof("Flushing log buffer...")

		// ignoring err because there is no buffer for stderr
		_ = zap.L().Sync()

		zap.S().Infof("Log buffer flushed!")
	})

	return nil
}

func gracefulShutdown() {
	if shutdownErr := shutdown.Wait(); shutdownErr != nil {
		zap.S().Error(shutdownErr)

		return
	}

	zap.S().Infof("Shutdown completed in %.1f seconds", shutdown.Timeout().Seconds())
}
