package config

import (
	"encoding/json"
	"fmt"
	"interview-client/internal/logger"
	"os"
)

type (
	GrpcConfig struct {
		Server     string `json:"Server"`
		Authserver string `json:"Authserver"`
	}

	EnvConfig struct {
		User string
		Pass string
	}
)

func LoadConfig() (c GrpcConfig) {
	f, err := os.Open("./config/local.json")
	defer f.Close()
	if err != nil {
		logger.LogError(fmt.Errorf("failed to open config file: %v", err), true)
	}

	decoder := json.NewDecoder(f)
	err = decoder.Decode(&c)
	if err != nil {
		logger.LogError(fmt.Errorf("failed to decode config file: %v", err), true)
	}

	return c
}

// loading user creds from enviroment for simplicity to mock login
func LoadEnv() *EnvConfig {
	user := os.Getenv("USER")
	if user == "" {
		logger.LogError(fmt.Errorf("failed to read USER from client environment"), true)
	}

	pass := os.Getenv("PASS")
	if pass == "" {
		logger.LogError(fmt.Errorf("failed to read PASS from client environment"), true)
	}

	return &EnvConfig{
		User: user,
		Pass: pass,
	}
}
