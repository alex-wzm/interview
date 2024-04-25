package config

import (
	"encoding/json"
	"log"
	"os"
)

type (
	GrpcConfig struct {
		ServerHost      string `json:"server_host"`
		UnsecurePort    string `json:"unsecure_port" default:"8085"`
		ResultProjectID string `json:"result_project_id" default:"my-project-id"`
		ResultTopicID   string `json:"result_topic_id"`
	}
)

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
