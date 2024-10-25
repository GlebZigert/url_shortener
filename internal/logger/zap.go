package logger

import (
	"context"

	"go.uber.org/zap"
)

// реалтзация zap
type ZapLogger struct {
	logger *zap.Logger
	ctx    context.Context
}

// конструктор
func NewZapLogger(loggerType string, ctx context.Context) *ZapLogger {
	logger, _ := zap.NewProduction()

	return &ZapLogger{logger: logger, ctx: ctx}
}

// для отладки
func (l *ZapLogger) Debug(msg string, fields map[string]interface{}) {

	l.logger.Debug("", zap.Any("args", fields))
}

// для инфо
func (l *ZapLogger) Info(msg string, fields map[string]interface{}) {

	l.logger.Info(msg, zap.Any("args", fields))
}

// для варнинга
func (l *ZapLogger) Warn(msg string, fields map[string]interface{}) {

	l.logger.Warn("", zap.Any("args", fields))
}

// для ошибки
func (l *ZapLogger) Error(msg string, fields map[string]interface{}) {

	l.logger.Error(msg, zap.Any("args", fields))
}

// для фатальной ошибки
func (l *ZapLogger) Fatal(msg string, fields map[string]interface{}) {

	l.logger.Fatal("", zap.Any("args", fields))
}
