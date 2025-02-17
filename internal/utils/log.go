package utils

import (
	"os"
	"time"

	"github.com/sirupsen/logrus"
)

// Logger is a global logger instance
var Logger *logrus.Logger

func InitLogger() {
	Logger = logrus.New()

	// Set log format (JSON or Text)
	Logger.SetFormatter(&logrus.JSONFormatter{
		TimestampFormat: time.RFC3339, // Custom timestamp format
	})

	// Set output to file and console
	os.MkdirAll("logs", 0755)
	file, err := os.OpenFile("logs/app.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err == nil {
		Logger.SetOutput(file) // Log to file
	} else {
		Logger.SetOutput(os.Stdout) // Fallback to console
		Logger.Warn("Failed to log to file, using default stdout")
	}

	// Set log level
	Logger.SetLevel(logrus.DebugLevel) // Debug, Info, Warn, Error, Fatal, Panic
}

// LogInfo logs an info message
func LogInfo(message string, fields logrus.Fields) {
	Logger.WithFields(fields).Info(message)
}

// LogError logs an error message
func LogError(message string, err error, fields logrus.Fields) {
	Logger.WithFields(fields).WithError(err).Error(message)
}
