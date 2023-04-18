package main

import (
	"flag"
	"interview-service/cmd/service"
	"os"

	log "github.com/sirupsen/logrus"
)

func init() {
	log.SetFormatter(&log.JSONFormatter{})
	log.SetOutput(os.Stdout)
	log.SetLevel(log.InfoLevel)
}

func main() {
	loglevel := flag.Int("logLevel", 4, "Useful Log levels: Warn = 3; Info = 4; Debug = 5;")
	grpcLogs := flag.Bool("grpcLogs", false, "Turn ON/OFF grpc middleware logs")
	flag.Parse()

	if *grpcLogs {
		os.Setenv("GRPC_GO_LOG_VERBOSITY_LEVEL", "99")
		os.Setenv("GRPC_GO_LOG_SEVERITY_LEVEL", "info")
	}

	if *loglevel >= 0 && *loglevel <= 6 {
		log.SetLevel(log.Level(*loglevel))
		log.Debug("Running in Debug mode!")
	}
	service.Start()
}
