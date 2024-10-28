package logger

import (
	"context"
	"os"

	"github.com/sirupsen/logrus"
)

// реализация логрус
type LogrusLogger struct {
	logger *logrus.Logger
	ctx    context.Context
}

// конструктор
func NewLogrusLogger(loggerType string, ctx context.Context) *LogrusLogger {
	logger := logrus.New()
	logger.Out = os.Stdout

	return &LogrusLogger{logger: logger, ctx: ctx}
}

// для отладки
func (l *LogrusLogger) Debug(msg string, fields map[string]interface{}) {

	l.logger.WithFields(fields).Debug(msg)
}

// для инфо
func (l *LogrusLogger) Info(msg string, fields map[string]interface{}) {

	l.logger.WithFields(fields).Info(msg)
}

// для варнинга
func (l *LogrusLogger) Warn(msg string, fields map[string]interface{}) {
	l.logger.WithFields(fields).Warn(msg)
}

// для ошибки
func (l *LogrusLogger) Error(msg string, fields map[string]interface{}) {
	l.logger.WithFields(fields).Error(msg)
}

// для фатальной ошибки
func (l *LogrusLogger) Fatal(msg string, fields map[string]interface{}) {
	l.logger.WithFields(fields).Fatal(msg)
}
