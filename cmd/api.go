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

	adminAPI "github.com/ra9dev/go-template/internal/api/admin"
	"github.com/ra9dev/go-template/internal/config"
	"github.com/ra9dev/go-template/pkg/shutdown"
)

const (
	readTimeout  = 5 * time.Second
	writeTimeout = 30 * time.Second
)

func APIServerCMD(cfg config.Config) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "api",
		Short: "api server of project",
	}

	cmd.RunE = func(cmd *cobra.Command, args []string) error {
		httpSrvRun := func(srv *http.Server) error {
			zap.S().Infof("Listening http on %s...", srv.Addr)

			if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
				return fmt.Errorf("http server failed to serve: %w", err)
			}

			return nil
		}

		group, _ := errgroup.WithContext(cmd.Context())

		group.Go(func() error {
			clientHTTPHandler := newHTTPClientHandler()
			clientHTTP := newHTTPServer(cfg.Ports.HTTP)(clientHTTPHandler)

			return httpSrvRun(clientHTTP)
		})

		group.Go(func() error {
			adminHTTPHandler := newHTTPAdminHandler()
			adminHTTP := newHTTPServer(cfg.Ports.AdminHTTP)(adminHTTPHandler)

			return httpSrvRun(adminHTTP)
		})

		if err := group.Wait(); err != nil {
			return fmt.Errorf("one of the listeners failed to run: %w", err)
		}

		return nil
	}

	return cmd
}

func newHTTPClientHandler() *chi.Mux {
	mux := chi.NewMux()

	return mux
}

func newHTTPAdminHandler() *chi.Mux {
	mux := chi.NewMux()

	mux.Mount("/v1", adminAPI.NewRouter())

	return mux
}

func newHTTPServer(port uint) func(handler http.Handler) *http.Server {
	return func(handler http.Handler) *http.Server {
		addr := fmt.Sprintf(":%d", port)

		srv := http.Server{
			Addr:         addr,
			Handler:      handler,
			ReadTimeout:  readTimeout,
			WriteTimeout: writeTimeout,
		}

		shutdown.Add(func(ctx context.Context) {
			zap.S().Info("Shutting down http on %s", addr)
			_ = srv.Shutdown(ctx)
			zap.S().Info("HTTP shutdown succeeded!")
		})

		return &srv
	}
}
