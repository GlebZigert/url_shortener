package grpcserver

import (
	"context"

	"github.com/GlebZigert/url_shortener.git/internal/services"
	pb "github.com/GlebZigert/url_shortener.git/proto"
)

type GrpcServerConfig interface {
	GetRunAddr() string
	GetBaseURL() string
	GetFlagLogLevel() string
	GetFileStoragePath() string
	GetNumWorkers() int
	GetDatabaseDSN() string
	GetTOKENEXP() int
	GetSECRETKEY() string
	GetENABLEHTTPSflag() bool
}

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

	cfg GrpcServerConfig
	// используем sync.Map для хранения пользователей
	service *services.Service

	logger GrpcServerLogger

	auc GrpcServerAuc
}

func New(cfg GrpcServerConfig,
	srvc *services.Service,
	logger GrpcServerLogger,
	auc GrpcServerAuc) *UrlShortenerServer {
	return &UrlShortenerServer{cfg: cfg, service: srvc, logger: logger, auc: auc}
}
