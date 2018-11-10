package logs

import (
	"github.com/sirupsen/logrus"
)

type Logger struct {
	logger *logrus.Logger
}

func NewLogger() *Logger {
	return &Logger{
		logger: logrus.StandardLogger(),
	}
}

func (log *Logger) Info(args ...interface{}) {
	log.logger.Info(args...)
}

func (log *Logger) Error(args ...interface{}) {
	log.logger.Error(args...)
}

func (log *Logger) Fatal(args ...interface{}) {
	log.logger.Fatal(args...)
}

func (log *Logger) AsLogrusLogger() *logrus.Logger {
	return log.logger
}
