package middleware

type mdlAuth interface {
	BuildJWTString(id int) (string, error)
	GetUserID(tokenString string) (int, error)
}

type mdlLogger interface {
	Info(msg string, fields map[string]interface{})
	Error(msg string, fields map[string]interface{})
}

type Middleware struct {
	auch   mdlAuth
	logger mdlLogger
}

func NewMiddlewares(auth mdlAuth, logger mdlLogger) *Middleware {
	return &Middleware{auth, logger}
}
