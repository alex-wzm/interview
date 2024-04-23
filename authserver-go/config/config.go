package config

import (
	"encoding/json"
	"fmt"
	"interview-authserver/internal/logger"
	"os"
)

type (
	GrpcConfig struct {
		ServerHost   string `json:"server_host" default:"localhost"`
		UnsecurePort string `json:"unsecure_port" default:"8081"`
	}

	EnvConfig struct {
		JwtSecret  string
		Username   string
		Usersecret string
		Host       string
		Port       string
	}
)

func LoadConfigFromFile(path string) *GrpcConfig {
	var cfg GrpcConfig
	val, err := os.ReadFile(path)
	if err != nil {
		logger.LogError(fmt.Errorf("error reading file; %v", err), true)
	}

	if unMarErr := json.Unmarshal(val, &cfg); unMarErr != nil {
		logger.LogError(fmt.Errorf("error during unmarshal; %v", err), true)
	}

	return &cfg
}

func LoadEnv() *EnvConfig {
	jwtSecret := os.Getenv("JWT_SECRET")
	if jwtSecret == "" {
		logger.LogError(fmt.Errorf("failed to read JWT_SECRET from environment"), true)
	}

	username := os.Getenv("USER_NAME")
	if username == "" {
		logger.LogError(fmt.Errorf("failed to read USER_NAME from environment"), true)
	}

	usersercret := os.Getenv("USER_SECRET")
	if usersercret == "" {
		logger.LogError(fmt.Errorf("failed to read USER_SECRET from environment"), true)
	}

	host := os.Getenv("HOST")
	if host == "" {
		logger.LogError(fmt.Errorf("failed to read HOST from environment"), true)
	}

	port := os.Getenv("PORT")
	if port == "" {
		logger.LogError(fmt.Errorf("failed to read PORT from environment"), true)
	}

	return &EnvConfig{
		JwtSecret:  jwtSecret,
		Username:   username,
		Usersecret: usersercret,
		Host:       host,
		Port:       port,
	}
}
