package main

import (
	"context"
	"os/signal"
	"syscall"

	"go.uber.org/zap"

	"github.com/spf13/cobra"

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
		zap.S().Fatalf("failed to prepare config: %s", err.Error())
	}

	_, err = log.NewLogger(cfg.LogLevel.ToZapAtomic())
	if err != nil {
		zap.S().Fatalf("failed to prepare logger: %s", err.Error())
	}

	shutdown.Add(func(_ context.Context) { _ = zap.L().Sync() })

	rootCmd.AddCommand(
		APIServerCMD(cfg),
	)

	osCTX, cancel := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer cancel()

	defer func() {
		zap.S().Infof("Shutdown timeout is %.1f seconds", shutdown.Timeout().Seconds())
		shutdown.Wait()
		zap.S().Info("Shutdown has been completed!")
	}()

	if err = rootCmd.ExecuteContext(osCTX); err != nil {
		zap.S().Errorf("failed to execute root cmd: %s", err.Error())

		return
	}
}
