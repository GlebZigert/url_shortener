// Package logger is a package for logging
package logger

// Logger interface
type Logger interface {
	//	Debug(msg string, fields map[string]interface{})
	Info(msg string, fields map[string]interface{})
	//	Warn(msg string, fields map[string]interface{})
	Error(msg string, fields map[string]interface{})
	//	Fatal(msg string, fields map[string]interface{})
}
