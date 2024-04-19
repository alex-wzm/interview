package config

import (
	"encoding/json"
	"log"
	"os"
)

type (
	GrpcConfig struct {
		ServerHost   string `json:"server_host" default:"localhost"`
		UnsecurePort string `json:"unsecure_port" default:"8080"`
		AuthServer   string `json:"AuthServer" default:"127.0.0.1:9000"`
	}
)

// Function to load configuarations from config file
func LoadConfigFromFile(path string) *GrpcConfig {
	var cfg GrpcConfig
	val, err := os.ReadFile(path)
	if err != nil {
		log.Fatalln("error reading file", err)
	}

	if unMarErr := json.Unmarshal(val, &cfg); unMarErr != nil {
		log.Fatalln("error during unmarshal", err)
	}

	return &cfg
}
