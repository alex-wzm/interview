package config

import (
	"log"
	"strings"

	"github.com/spf13/viper"
)

type (
	GrpcConfig struct {
		ServerHost   string `json:"server_host" default:"localhost"`
		UnsecurePort string `json:"unsecure_port" default:"8080"`
	}
	Config struct {
		Grpc      *GrpcConfig
		JwtSecret string
	}
)

func LoadConfig() *Config {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	viper.AddConfigPath("./config")

	// Set up to automatically read configuration from environment variables
	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file, %s", err)
	}

	grpcConfig := &GrpcConfig{}
	if err := viper.Sub("grpc").Unmarshal(grpcConfig); err != nil {
		log.Fatalf("Error unmarshalling grpc config: %s", err)
	}

	return &Config{
		Grpc:      grpcConfig,
		JwtSecret: viper.GetString("jwt"),
	}
}
