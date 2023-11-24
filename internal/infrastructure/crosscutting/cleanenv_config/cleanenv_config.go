package cleanenv

import (
	"log"
	"os"

	"github.com/AskhatZRPV/rprofile-grpc/internal/core/config"
	"github.com/ilyakaznacheev/cleanenv"

	"github.com/pkg/errors"
)

func New() (*config.Config, error) {
	configPath := os.Getenv("CONFIG_PATH")
	if configPath == "" {
		log.Fatal("CONFIG_PATH is not set")
	}

	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		return nil, errors.Wrap(err, "config file does not exist")
	}

	var cfg config.Config
	if err := cleanenv.ReadConfig(configPath, &cfg); err != nil {
		return nil, errors.Wrap(err, "cannot read config")
	}

	return &cfg, nil
}
