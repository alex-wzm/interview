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

	grpcConfig := &GrpcConfig{
		UnsecurePort: viper.GetString("grpc.unsecure_port"),
		ServerHost:   viper.GetString("grpc.server_host"),
	}

	return &Config{
		Grpc:      grpcConfig,
		JwtSecret: viper.GetString("jwt"),
	}
}
