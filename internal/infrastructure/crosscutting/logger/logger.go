package logger

import (
	"log/slog"
	"os"

	"github.com/AskhatZRPV/rprofile-grpc/internal/core/config"
)

const (
	envLocal = "local"
	envDev   = "dev"
	envProd  = "prod"
)

func New(config *config.Config) *slog.Logger {
	var log *slog.Logger

	switch config.Env {
	case envLocal:
		log = slog.New(
			slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}),
		)
	case envDev:
		log = slog.New(
			slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}),
		)
	case envProd:
		log = slog.New(
			slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}),
		)
	default:
		log = slog.New(
			slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}),
		)
	}

	return log
}
