package logger

import (
	"context"

	"go.uber.org/zap"
)

// ZapLogger struct
type ZapLogger struct {
	logger *zap.Logger
	ctx    context.Context
}

// NewZapLogger func is constructor
func NewZapLogger(loggerType string, ctx context.Context) *ZapLogger {
	logger, _ := zap.NewProduction()

	return &ZapLogger{logger: logger, ctx: ctx}
}

// Debug func for debug message
func (l *ZapLogger) Debug(msg string, fields map[string]interface{}) {

	l.logger.Debug("", zap.Any("args", fields))
}

// Info func for info message
func (l *ZapLogger) Info(msg string, fields map[string]interface{}) {

	l.logger.Info(msg, zap.Any("args", fields))
}

// Warn func for warning message
func (l *ZapLogger) Warn(msg string, fields map[string]interface{}) {

	l.logger.Warn("", zap.Any("args", fields))
}

// Error func for error message
func (l *ZapLogger) Error(msg string, fields map[string]interface{}) {

	l.logger.Error(msg, zap.Any("args", fields))
}

// Fatal func for fatal message
func (l *ZapLogger) Fatal(msg string, fields map[string]interface{}) {

	l.logger.Fatal("", zap.Any("args", fields))
}
