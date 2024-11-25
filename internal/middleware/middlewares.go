package middleware

import "context"

type mdlAuth interface {
	BuildJWTString(id int) (string, error)
	GetUserID(tokenString string) (int, error)
	CheckUID(ctx context.Context) (user int, ok bool)
	SetUID(ctx context.Context, user int) context.Context
	CheckNewFlag(ctx context.Context) (fl bool, ok bool)
	SetNewFlag(ctx context.Context, fl bool) context.Context
}

type mdlLogger interface {
	Info(msg string, fields map[string]interface{})
	Error(msg string, fields map[string]interface{})
}

// интерфейс для источника CIDR
type SrcCIDR interface {
	GetCIDR() string // метод получения CIDR
}

// струткура с методами-мидлами
type Middleware struct {
	mdlAuth
	logger  mdlLogger
	SrcCIDR //источник CIDR для миддла который фильтрует по IP адресам
}

// ее конструктор
func NewMiddlewares(auth mdlAuth, logger mdlLogger, cidr SrcCIDR) *Middleware {
	return &Middleware{auth, logger, cidr}
}
