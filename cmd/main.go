package main

import (
	"context"
	"fmt"
	"os/signal"
	"syscall"

	"github.com/spf13/cobra"
	"go.uber.org/zap"

	"github.com/ra9dev/go-template/internal/config"
	"github.com/ra9dev/go-template/pkg/log"
	"github.com/ra9dev/go-template/pkg/shutdown"
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

	_, err = log.NewLogger(cfg.LogLevel.ToZapAtomic())
	if err != nil {
		zap.S().Fatalf("failed to prepare logger: %v", err)
	}

	shutdown.Add(func(_ context.Context) { _ = zap.L().Sync() })

	rootCmd.AddCommand(
		APIServerCMD(cfg),
	)

	if err = newTraceProvider(cfg); err != nil {
		zap.S().Fatalf("failed to prepare trace provider: %v", err)
	}

	osCTX, cancel := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer cancel()

	defer func() {
		zap.S().Infof("Shutdown timeout is %.1f seconds", shutdown.Timeout().Seconds())
		shutdown.Wait()
		zap.S().Info("Shutdown has been completed!")
	}()

	if err = rootCmd.ExecuteContext(osCTX); err != nil {
		zap.S().Errorf("failed to execute root cmd: %v", err)

		return
	}
}

func newTraceProvider(cfg config.Config) error {
	provider, err := tracing.NewProvider(tracing.Config{
		ServiceName:    config.ServiceName,
		ServiceVersion: config.ServiceVersion,
		Environment:    cfg.Environment,
		Endpoint:       cfg.Tracing.Endpoint,
		Enabled:        cfg.Tracing.Enabled,
	})
	if err != nil {
		return fmt.Errorf("failed to create trace provider: %w", err)
	}

	shutdown.Add(func(ctx context.Context) {
		zap.S().Info("Shutting down tracing provider")
		provider.Shutdown(ctx)
		zap.S().Info("Tracing provider shutdown succeeded!")
	})

	return nil
}
