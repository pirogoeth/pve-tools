package logging

import (
	"io"
	"os"

	"github.com/sirupsen/logrus"
)

// New prepares a logger instance with an opinionated output format.
func New() {
	logLevel, err := logrus.ParseLevel(os.Getenv("LOG_LEVEL"))
	if err != nil {
		logLevel = logrus.InfoLevel
	}

	logrus.SetFormatter(&logrus.TextFormatter{
		ForceColors: true,
	})
	logrus.SetOutput(os.Stdout)
	logrus.SetLevel(logLevel)
}

// Fatalf logs a message at the fatal level and exits.
func Fatalf(format string, args ...interface{}) {
	logrus.Fatalf(format, args...)
}

// Errorf logs a message at the error level.
func Errorf(format string, args ...interface{}) {
	logrus.Errorf(format, args...)
}

// Warnf logs a message at the warn level.
func Warnf(format string, args ...interface{}) {
	logrus.Warnf(format, args...)
}

// Infof logs a message at the info level.
func Infof(format string, args ...interface{}) {
	logrus.Infof(format, args...)
}

// Debugf logs a message at the debug level.
func Debugf(format string, args ...interface{}) {
	logrus.Debugf(format, args...)
}

// Tracef logs a message at the trace level.
func Tracef(format string, args ...interface{}) {
	logrus.Tracef(format, args...)
}

// GetLogger returns the logger instance.
func GetLogger() *logrus.Logger {
	return logrus.StandardLogger()
}

// LoggerWriter returns an io.Writer that writes to the logger.
func LoggerWriter() io.Writer {
	return GetLogger().Writer()
}
