package logger

import (
	log "github.com/sirupsen/logrus"
)

func LogError(err error, fatal bool) {
	if err != nil {
		if fatal {
			log.Fatal(err)
		} else {
			log.Error(err)
		}
	}
}

func LogInfo(msg string) {
	log.Infof("info: %v", msg)
}
