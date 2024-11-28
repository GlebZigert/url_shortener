package middleware

import "context"

// для получения свежего uid
type MdlUserStore interface {
	CreateNextUID() (int, error)
}

type MdlAuth interface {
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
	GetCIDR() []string // метод получения CIDR
}

// струткура с методами-мидлами
type Middleware struct {
	MdlAuth
	logger  mdlLogger
	SrcCIDR //источник CIDR для миддла который фильтрует по IP адресам
	MdlUserStore
}

// ее конструктор
func NewMiddlewares(auth MdlAuth, logger mdlLogger, cidr SrcCIDR, user MdlUserStore) *Middleware {
	return &Middleware{auth, logger, cidr, user}
}
