package logger

import (
	"context"
	"os"

	"github.com/sirupsen/logrus"
)

// LogrusLogger struct
type LogrusLogger struct {
	logger *logrus.Logger
	ctx    context.Context
}

// NewLogrusLogger func is constructor for LogrusLogger struct
func NewLogrusLogger(loggerType string, ctx context.Context) *LogrusLogger {
	logger := logrus.New()
	logger.Out = os.Stdout

	return &LogrusLogger{logger: logger, ctx: ctx}
}

// Debug for debug
func (l *LogrusLogger) Debug(msg string, fields map[string]interface{}) {

	l.logger.WithFields(fields).Debug(msg)
}

// Info for info
func (l *LogrusLogger) Info(msg string, fields map[string]interface{}) {

	l.logger.WithFields(fields).Info(msg)
}

// Warn for warning
func (l *LogrusLogger) Warn(msg string, fields map[string]interface{}) {
	l.logger.WithFields(fields).Warn(msg)
}

// Error for error
func (l *LogrusLogger) Error(msg string, fields map[string]interface{}) {
	l.logger.WithFields(fields).Error(msg)
}

// Fatal for fatal
func (l *LogrusLogger) Fatal(msg string, fields map[string]interface{}) {
	l.logger.WithFields(fields).Fatal(msg)
}
