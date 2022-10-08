package main

import (
	"context"
	"errors"
	"fmt"
	"net"
	"net/http"
	"time"

	chi "github.com/go-chi/chi/v5"
	"github.com/spf13/cobra"
	"go.uber.org/zap"
	"golang.org/x/sync/errgroup"
	"google.golang.org/grpc"

	grpcAPI "github.com/ra9dev/go-template/internal/api/grpc"
	adminAPI "github.com/ra9dev/go-template/internal/api/http/admin"
	"github.com/ra9dev/go-template/internal/config"
	example "github.com/ra9dev/go-template/pb"
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

		group.Go(func() error {
			grpcListener, err := net.Listen("tcp", fmt.Sprintf(":%d", cfg.Ports.GRPC))
			if err != nil {
				return fmt.Errorf("failed to listen port %d for grpc: %w", cfg.Ports.GRPC, err)
			}

			zap.S().Infof("Listening GRPC on :%d...", cfg.Ports.GRPC)

			grpcServer := newGRPCServer(cfg.Ports.GRPC)

			if err = grpcServer.Serve(grpcListener); err != nil {
				return fmt.Errorf("GRPC server failed to serve: %w", err)
			}

			return nil
		})

		if err := group.Wait(); err != nil {
			return fmt.Errorf("api group failed: %w", err)
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
			zap.S().Infof("Shutting down HTTP on %s...", addr)

			if err := srv.Shutdown(ctx); err != nil {
				zap.S().Errorf("HTTP shutdown failed: %v", err)
				return
			}

			zap.S().Info("HTTP shutdown succeeded!")
		})

		return &srv
	}
}

func httpSrvRun(srv *http.Server) error {
	zap.S().Infof("Listening HTTP on %s...", srv.Addr)

	if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
		return fmt.Errorf("HTTP server failed to serve: %w", err)
	}

	return nil
}

func newGRPCServer(port uint) *grpc.Server {
	srv := grpc.NewServer()
	exampleService := grpcAPI.NewExampleService()

	example.RegisterGreeterServer(srv, exampleService)

	shutdown.Add(func(ctx context.Context) {
		zap.S().Infof("Shutting down GRPC on :%d...", port)
		srv.GracefulStop()
		zap.S().Info("GRPC shutdown succeeded!")
	})

	return srv
}
