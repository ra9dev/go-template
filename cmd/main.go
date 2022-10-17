package main

import (
	"context"
	"fmt"
	"github.com/ra9dev/go-template/pkg/sre/log"
	"github.com/ra9dev/go-template/pkg/sre/tracing"

	"github.com/ra9dev/go-template/internal/config"
	"github.com/ra9dev/shutdown"
	"github.com/spf13/cobra"
)

func main() {
	defer gracefulShutdown()

	ctx := shutdown.Context()
	rootCmd := &cobra.Command{
		Use:   "go-template",
		Short: "Main entry-point command for the application",
	}

	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatalf(ctx, "failed to prepare config: %v", err)
	}

	if err = setupLogger(cfg); err != nil {
		log.Fatal(ctx, err)
	}

	if err = setupTracing(cfg); err != nil {
		log.Fatal(ctx, err)
	}

	rootCmd.AddCommand(
		APIServerCMD(cfg),
	)

	if err = rootCmd.ExecuteContext(shutdown.Context()); err != nil {
		log.Errorf(ctx, "failed to execute root cmd: %v", err)

		return
	}
}

func setupTracing(cfg config.Config) error {
	provider, err := tracing.NewProvider(tracing.Config{
		ServiceName:    config.ServiceName,
		ServiceVersion: config.ServiceVersion,
		Environment:    cfg.Env.String(),
		Endpoint:       cfg.Tracing.Endpoint,
		Enabled:        cfg.Tracing.Enabled,
	})
	if err != nil {
		return fmt.Errorf("failed to prepare tracing provider: %w", err)
	}

	shutdown.MustAdd("tracing", func(ctx context.Context) {
		log.NoContext().Info("Shutting down tracing provider")

		if err = provider.Shutdown(ctx); err != nil {
			log.NoContext().Error(err)

			return
		}

		log.NoContext().Info("Tracing provider shutdown succeeded!")
	})

	return nil
}

func setupLogger(cfg config.Config) error {
	loggerParams := log.NewParams(cfg.Env, cfg.LogLevel)

	logger, err := log.NewLogger(loggerParams)
	if err != nil {
		return fmt.Errorf("failed to prepare logger: %w", err)
	}

	log.RegisterLogger(logger)

	shutdown.MustAdd("logger", func(_ context.Context) {
		log.NoContext().Infof("Flushing log buffer...")

		// ignoring err because there is no buffer for stderr
		_ = log.Sync()

		log.NoContext().Infof("Log buffer flushed!")
	})

	return nil
}

func gracefulShutdown() {
	if shutdownErr := shutdown.Wait(); shutdownErr != nil {
		log.NoContext().Error(shutdownErr)

		return
	}

	log.NoContext().Infof("Shutdown completed in %.1f seconds", shutdown.Timeout().Seconds())
}
