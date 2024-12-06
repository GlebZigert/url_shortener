package interseptors

import "context"

// для получения свежего uid
type InterseptorUserStore interface {
	CreateNextUID() (int, error)
}

// Интерфейс авторизации
type InterseptorAuth interface {
	BuildJWTString(id int) (string, error)
	GetUserID(tokenString string) (int, error)
	CheckUID(ctx context.Context) (user int, ok bool)
	SetUID(ctx context.Context, user int) context.Context
	CheckNewFlag(ctx context.Context) (fl bool, ok bool)
	SetNewFlag(ctx context.Context, fl bool) context.Context
}

type InterseptorLogger interface {
	Info(msg string, fields map[string]interface{})
	Error(msg string, fields map[string]interface{})
}

// интерфейс для источника CIDR
type SrcCIDR interface {
	GetCIDR() []string // метод получения CIDR
}

// струткура с методами-мидлами
type Interseptors struct {
	InterseptorAuth
	logger  InterseptorLogger
	SrcCIDR //источник CIDR для миддла который фильтрует по IP адресам
	users   InterseptorUserStore
}

// ее конструктор
func NewInterseptors(auth InterseptorAuth,
	logger InterseptorLogger,
	cidr SrcCIDR,
	user InterseptorUserStore) *Interseptors {
	return &Interseptors{auth, logger, cidr, user}
}
