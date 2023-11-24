package config

import "time"

type GRPC struct {
	Port    int           `yaml:"port" env-default:"9090"`
	Timeout time.Duration `yaml:"timeout" env-default:"5s"`
}

type GRPCGW struct {
	Port    int           `yaml:"port" env-default:"8080"`
	Timeout time.Duration `yaml:"timeout" env-default:"5s"`
}

type HttpClient struct {
	BaseUrl string        `yaml:"base_url" env-default:"https://www.rusprofile.ru"`
	Timeout time.Duration `yaml:"timeout" env-default:"5s"`
}

type Config struct {
	Env        string `yaml:"env" env-default:"local"`
	HttpClient `yaml:"http_client"`
	GRPC       `yaml:"grpc"`
	GRPCGW     `yaml:"grpc_gw"`
}
