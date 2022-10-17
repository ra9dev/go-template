package main

import (
	"context"
	"errors"
	"fmt"
	"github.com/ra9dev/go-template/pkg/sre/log"
	"net"
	"net/http"
	"time"

	chi "github.com/go-chi/chi/v5"
	"github.com/ra9dev/shutdown"
	"github.com/spf13/cobra"
	"golang.org/x/sync/errgroup"
	"google.golang.org/grpc"

	grpcAPI "github.com/ra9dev/go-template/internal/api/grpc"
	adminAPI "github.com/ra9dev/go-template/internal/api/http/admin"
	"github.com/ra9dev/go-template/internal/config"
	example "github.com/ra9dev/go-template/pb"
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
		ctx := cmd.Context()
		group, groupCTX := errgroup.WithContext(ctx)

		group.Go(func() error {
			clientHTTPHandler := newHTTPClientHandler()
			clientHTTP := newHTTPServer(cfg.Ports.HTTP)(clientHTTPHandler)

			return httpSrvRun(groupCTX, clientHTTP)
		})

		group.Go(func() error {
			adminHTTPHandler := newHTTPAdminHandler()
			adminHTTP := newHTTPServer(cfg.Ports.AdminHTTP)(adminHTTPHandler)

			return httpSrvRun(groupCTX, adminHTTP)
		})

		group.Go(func() error {
			return grpcSrvRun(groupCTX, cfg.Ports.GRPC)
		})

		if err := group.Wait(); err != nil {
			return fmt.Errorf("api group failed: %w", err)
		}

		return nil
	}

	return cmd
}

func httpSrvRun(ctx context.Context, srv *http.Server) error {
	log.Infof(ctx, "Listening HTTP on %s...", srv.Addr)

	if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
		return fmt.Errorf("HTTP server failed to serve: %w", err)
	}

	return nil
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

		shutdownKey := fmt.Sprintf("http:%d", port)

		shutdown.MustAdd(shutdownKey, func(ctx context.Context) {
			log.NoContext().Infof("Shutting down HTTP on %s...", addr)

			if err := srv.Shutdown(ctx); err != nil {
				log.NoContext().Errorf("HTTP shutdown failed: %v", err)

				return
			}

			log.NoContext().Info("HTTP shutdown succeeded!")
		})

		return &srv
	}
}

func grpcSrvRun(ctx context.Context, port uint) error {
	grpcListener, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		return fmt.Errorf("failed to listen port %d for grpc: %w", port, err)
	}

	log.Infof(ctx, "Listening GRPC on :%d...", port)

	grpcServer := newGRPCServer(port)

	if err := grpcServer.Serve(grpcListener); err != nil {
		return fmt.Errorf("GRPC server failed to serve: %w", err)
	}

	return nil
}

func newGRPCServer(port uint) *grpc.Server {
	srv := grpc.NewServer()
	exampleService := grpcAPI.NewExampleService()

	example.RegisterGreeterServer(srv, exampleService)

	shutdownKey := fmt.Sprintf("grpc:%d", port)

	shutdown.MustAdd(shutdownKey, func(ctx context.Context) {
		log.NoContext().Infof("Shutting down GRPC on :%d...", port)

		srv.GracefulStop()

		log.NoContext().Info("GRPC shutdown succeeded!")
	})

	return srv
}
