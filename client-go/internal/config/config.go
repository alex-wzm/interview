package config

import (
	"github.com/caarlos0/env/v11"
)

type Application struct {
	ServerAddr     string `env:"SERVER_ADDR" envDefault:"localhost:8080"`
	AuthServerAddr string `env:"AUTH_SERVER_ADDR" envDefault:"localhost:8081"`
	Debug          bool   `env:"DEBUG"`
}

func LoadFromEnv() (*Application, error) {
	cfg := Application{}
	parseErr := env.Parse(&cfg)
	return &cfg, parseErr
}
