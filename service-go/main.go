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
	loglevel := flag.Int("loglevel", 4, "Useful Log levels: Warn = 3; Info = 4; Debug = 5;")
	flag.Parse()
	if *loglevel >= 0 && *loglevel <= 6 {
		log.SetLevel(log.Level(*loglevel))
		log.Debug("Running in Debug mode!")
	}
	service.Start()
}
