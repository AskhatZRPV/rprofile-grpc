package main

import (
	"go.uber.org/fx"
	"google.golang.org/grpc"

	"github.com/AskhatZRPV/rprofile-grpc/internal/application/rusprofile/parser"
	"github.com/AskhatZRPV/rprofile-grpc/internal/infrastructure/crosscutting/logger"
	"github.com/AskhatZRPV/rprofile-grpc/internal/infrastructure/httpclient"
	"github.com/AskhatZRPV/rprofile-grpc/internal/interface/grpc/rusprofile/search"
	"github.com/AskhatZRPV/rprofile-grpc/internal/interface/httpserver"

	rprofilev1 "github.com/AskhatZRPV/rprofile-grpc/gen/protos/go/proto/rprofile"
	colly "github.com/AskhatZRPV/rprofile-grpc/internal/infrastructure/colly/rusprofile"
	cleanenvconfig "github.com/AskhatZRPV/rprofile-grpc/internal/infrastructure/crosscutting/cleanenv_config"
	grpcserver "github.com/AskhatZRPV/rprofile-grpc/internal/interface/grpc"
)

func main() {

	app := fx.New(
		fx.Provide(cleanenvconfig.New),
		fx.Provide(logger.New),
		fx.Provide(httpclient.New),
		fx.Provide(colly.New),
		fx.Provide(
			parser.New,
		),
		fx.Provide(
			fx.Annotate(
				grpcserver.New,
				fx.As(new(grpc.ServiceRegistrar)),
			),

			fx.Annotate(
				search.New,
				fx.As(new(rprofilev1.CompanyInfoServer)),
			),
		),

		fx.Invoke(
			func(grpc.ServiceRegistrar) {},

			rprofilev1.RegisterCompanyInfoServer,
		),

		fx.Invoke(
			httpserver.New,
		),
	)

	app.Run()
}
