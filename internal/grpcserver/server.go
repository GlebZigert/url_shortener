package grpcserver

import (
	"context"

	"github.com/GlebZigert/url_shortener.git/internal/services"
	pb "github.com/GlebZigert/url_shortener.git/proto"
)

type GrpcServerAuc interface {
	CheckUID(context.Context) (int, bool)
}

type GrpcServerLogger interface {
	Info(msg string, fields map[string]interface{})
	Error(msg string, fields map[string]interface{})
}

// UsersServer поддерживает все необходимые методы сервера.
type UrlShortenerServer struct {
	// нужно встраивать тип pb.Unimplemented<TypeName>
	// для совместимости с будущими версиями
	pb.UnimplementedUrlShortenerServer

	// используем sync.Map для хранения пользователей
	service *services.Service

	logger GrpcServerLogger

	auc GrpcServerAuc
}

func New(srvc *services.Service,
	logger GrpcServerLogger,
	auc GrpcServerAuc) *UrlShortenerServer {
	return &UrlShortenerServer{service: srvc, logger: logger, auc: auc}
}
