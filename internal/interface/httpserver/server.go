package httpserver

import (
	"context"
	"fmt"
	"net/http"

	"log/slog"

	rprofilev1 "github.com/AskhatZRPV/rprofile-grpc/gen/protos/go/proto/rprofile"
	"github.com/AskhatZRPV/rprofile-grpc/internal/core/config"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"go.uber.org/fx"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func New(lc fx.Lifecycle, log *slog.Logger, cfg *config.Config) {
	conn, err := grpc.Dial(fmt.Sprintf("localhost:%d", cfg.GRPC.Port), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Error("fail to dial: %v", err)
	}

	rmux := runtime.NewServeMux()
	client := rprofilev1.NewCompanyInfoClient(conn)

	mux := http.NewServeMux()

	mux.Handle("/", rmux)
	mux.HandleFunc("/swagger-ui/swagger.json", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./gen/proto/rprofile/rprofile.swagger.json")
	})

	mux.Handle("/swagger-ui/", http.StripPrefix("/swagger-ui/", http.FileServer(http.Dir("./swagger-ui"))))

	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			err := rprofilev1.RegisterCompanyInfoHandlerClient(ctx, rmux, client)
			if err != nil {
				log.Error(err.Error())
			}
			go func() {
				if err := http.ListenAndServe(fmt.Sprintf(":%d", cfg.GRPCGW.Port), mux); err != nil {
					log.Error("Failed to Serve gRPC-Gateway", err.Error())
				}
			}()

			return nil
		},
		OnStop: func(ctx context.Context) error {
			log.Info("Gracefully stopping gRPC server")
			defer conn.Close()
			return nil
		},
	})

}
