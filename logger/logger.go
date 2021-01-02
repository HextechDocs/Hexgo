// Package logger and its subpackages house everything related to logging
package logger

import (
	"github.com/sirupsen/logrus"
	"os"
)

// HexLogger is the logger instance with the settings we need
var HexLogger *logrus.Logger

// InitLogger initializes HexLogger and sets the settings we want
func InitLogger() {
	HexLogger = logrus.New()
	HexLogger.SetFormatter(&logrus.TextFormatter{})
	HexLogger.SetOutput(os.Stdout)

	if os.Getenv("LOGRUS_DEVEL") == "" {
		logrus.SetLevel(logrus.WarnLevel)
	} else {
		logrus.SetLevel(logrus.TraceLevel)
	}
}
