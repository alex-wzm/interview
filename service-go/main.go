package main

import (
	"flag"
	"interview-service/cmd/service"
	"os"

	"github.com/joho/godotenv"

	log "github.com/sirupsen/logrus"
)

func main() {
	loglevel := flag.Int("logLevel", 4, "Useful Log levels: Warn = 3; Info = 4; Debug = 5;")
	logGrpc := flag.Bool("logGrpc", false, "Turn ON/OFF grpc middleware logs")
	flag.Parse()

	log.SetFormatter(&log.JSONFormatter{})
	log.SetOutput(os.Stdout)

	if *loglevel >= 0 && *loglevel <= 6 {
		log.SetLevel(log.Level(*loglevel))
	}

	if *logGrpc {
		err := godotenv.Load("logGrpc.env")
		if err != nil {
			log.Infof("Error loading .env file: %+v", err)
		}
		log.Debugf("Loaded .env file")
	}
	log.Debugf("Severity: %s", os.Getenv("GRPC_GO_LOG_SEVERITY_LEVEL"))
	log.Debugf("Verbosity: %s", os.Getenv("GRPC_GO_LOG_VERBOSITY_LEVEL"))

	service.Start()
}
