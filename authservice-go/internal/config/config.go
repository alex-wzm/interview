package config

import (
	"fmt"
	"github.com/caarlos0/env/v11"
)

type (
	Application struct {
		Debug     bool              `env:"DEBUG"`
		JWTSecret string            `env:"JWT_SECRET"`
		Users     map[string]string `env:"USERS"`
		GRPC      GRPC
	}
	GRPC struct {
		ListenAddr string `env:"LISTEN_ADDR" envDefault:":8081"`
	}
)

func LoadFromEnv() (*Application, error) {
	cfg := &Application{}

	err := env.Parse(cfg)

	if cfg.JWTSecret == "" {
		return nil, fmt.Errorf("error loading JWT secret from environment")
	}

	return cfg, err
}
