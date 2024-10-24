package middleware

import "context"

type mdlAuth interface {
	BuildJWTString(id int) (string, error)
	GetUserID(tokenString string) (int, error)
	CheckUID(ctx context.Context) (user int, ok bool)
}

type mdlLogger interface {
	Info(msg string, fields map[string]interface{})
	Error(msg string, fields map[string]interface{})
}

type Middleware struct {
	mdlAuth
	logger mdlLogger
}

func NewMiddlewares(auth mdlAuth, logger mdlLogger) *Middleware {
	return &Middleware{auth, logger}
}
