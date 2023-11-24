package grpcserver

import (
	"context"
	"fmt"
	"log/slog"
	"net"

	"github.com/AskhatZRPV/rprofile-grpc/internal/core/config"
	"github.com/AskhatZRPV/rprofile-grpc/internal/interface/grpc/interceptors"
	"github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/logging"
	"github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/recovery"
	"go.uber.org/fx"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func New(lc fx.Lifecycle, log *slog.Logger, cfg *config.Config) *grpc.Server {

	loggingOpts := []logging.Option{
		logging.WithLogOnEvents(
			logging.PayloadReceived, logging.PayloadSent,
		),
	}

	recoveryOpts := []recovery.Option{
		recovery.WithRecoveryHandler(func(p interface{}) (err error) {
			log.Error("Recovered from panic", slog.Any("panic", p))
			return status.Errorf(codes.Internal, "internal error")
		}),
	}

	srv := grpc.NewServer(grpc.ChainUnaryInterceptor(
		recovery.UnaryServerInterceptor(recoveryOpts...),
		logging.UnaryServerInterceptor(interceptors.InterceptorLogger(log), loggingOpts...),
	))

	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			log.Info("Starting gRPC server")

			ln, err := net.Listen("tcp", fmt.Sprintf(":%d", cfg.GRPC.Port))
			if err != nil {
				return err
			}

			go func() {
				if err := srv.Serve(ln); err != nil {
					log.Error("Failed to Serve gRPC", err.Error())
				}
			}()

			return nil
		},
		OnStop: func(ctx context.Context) error {
			log.Info("Gracefully stopping gRPC server")

			srv.GracefulStop()

			return nil
		},
	})

	return srv
}
