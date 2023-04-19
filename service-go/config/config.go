package config

import (
	"encoding/json"
	"os"
)

type (
	GrpcConfig struct {
		ServerHost   string `json:"server_host" default:"localhost"`
		UnsecurePort string `json:"unsecure_port" default:"50051"`
	}
)

func LoadConfigFromFile(path string) (*GrpcConfig, error) {
	var cfg GrpcConfig
	val, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	if err = json.Unmarshal(val, &cfg); err != nil {
		return nil, err
	}

	return &cfg, nil
}
