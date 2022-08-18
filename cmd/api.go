package main

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"time"

	chi "github.com/go-chi/chi/v5"
	"github.com/spf13/cobra"
	"go.uber.org/zap"
	"golang.org/x/sync/errgroup"

	"github.com/ra9dev/go-template/pkg/shutdown"
)

const (
	readTimeout  = 5 * time.Second
	writeTimeout = 30 * time.Second
)

func APIServerCMD(cfg Config) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "api",
		Short: "api server of project",
	}

	cmd.RunE = func(cmd *cobra.Command, args []string) error {
		group, _ := errgroup.WithContext(cmd.Context())
		group.Go(func() error {
			httpSrv := httpServer(cfg.Ports.HTTP)

			zap.S().Infof("Listening http on %d...", cfg.Ports.HTTP)

			shutdown.Add(func(ctx context.Context) {
				zap.S().Info("Shutting down http")
				_ = httpSrv.Shutdown(ctx)
				zap.S().Info("HTTP shutdown succeeded!")
			})

			if err := httpSrv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
				return fmt.Errorf("http server failed to serve: %w", err)
			}

			return nil
		})

		if err := group.Wait(); err != nil {
			return fmt.Errorf("one of the listeners failed to run: %w", err)
		}

		return nil
	}

	return cmd
}

func httpServer(httpPort uint) *http.Server {
	mux := chi.NewMux()

	srv := &http.Server{
		Addr:         fmt.Sprintf(":%d", httpPort),
		Handler:      mux,
		ReadTimeout:  readTimeout,
		WriteTimeout: writeTimeout,
	}

	return srv
}
