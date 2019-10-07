package logger

import (
	"os"
	"strconv"

	"github.com/sirupsen/logrus"
)

// New initializes a logrus logger
func New() *logrus.Logger {
	level := "5"
	logLevel := 5

	if l, found := os.LookupEnv("LOGLEVEL"); found {
		level = l
	}
	if ll, err := strconv.Atoi(level); err == nil {
		logLevel = ll
	}

	log := logrus.New()
	log.Formatter = &logrus.TextFormatter{}
	log.ReportCaller = true
	log.Level = logrus.Level(logLevel)

	return log

}
