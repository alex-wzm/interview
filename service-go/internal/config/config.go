package config

import (
	"github.com/caarlos0/env/v11"
)

type (
	Application struct {
		Debug bool `env:"DEBUG"`
		GRPC  GRPC
	}

	GRPC struct {
		ServerAddr     string `env:"SERVER_ADDR" envDefault:":8080"`
		AuthServerAddr string `env:"AUTH_SERVER_ADDR" envDefault:"localhost:8081"`
	}
)

func LoadFromEnv() (*Application, error) {
	cfg := Application{}

	// keeping this outside of the return in the case we want to do additional checks
	// after parsing before we return
	err := env.Parse(&cfg)

	return &cfg, err
}
