package main

import (
	// ...
	"context"
	"fmt"
	"log"
	"net"
	"os"

	"github.com/GlebZigert/url_shortener.git/internal/auth"
	"github.com/GlebZigert/url_shortener.git/internal/config"
	"github.com/GlebZigert/url_shortener.git/internal/db"
	"github.com/GlebZigert/url_shortener.git/internal/filereader"
	"github.com/GlebZigert/url_shortener.git/internal/interseptors"
	"github.com/GlebZigert/url_shortener.git/internal/logger"
	"github.com/GlebZigert/url_shortener.git/internal/services"
	"github.com/GlebZigert/url_shortener.git/internal/storager"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	// импортируем пакет со сгенерированными protobuf-файлами
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

func (s *UrlShortenerServer) CreateShortURL(ctx context.Context, in *pb.CreateShortURLRequest) (*pb.CreateShortURLResponse, error) {
	var response pb.CreateShortURLResponse

	user, ok := s.auc.CheckUID(ctx)
	if !ok {
		s.logger.Error("Нет данных о пользователе: ", map[string]interface{}{})

		return nil, status.Error(codes.Internal, "")
	}

	res, err := s.service.Short(in.Origin, user)
	response.Short = res
	if err != nil {
		response.Error = err.Error()
	}

	return &response, nil
}

func main() {
	// определяем порт для сервера
	listen, err := net.Listen("tcp", ":3200")
	if err != nil {
		log.Fatal(err)
	}
	// создаём gRPC-сервер без зарегистрированной службы

	cfg, err := config.NewConfig(os.Args[0], os.Args[1:], filereader.New())
	if err != nil {
		return
	}

	ctx := context.Background()
	dber := db.Get(db.Init(cfg.GetDatabaseDSN()))
	store := storager.New(cfg, filereader.New(), dber)

	logger := logger.NewLogrusLogger(cfg.FlagLogLevel, ctx)

	srvc := services.NewService(logger, store)

	auc := auth.NewAuth(cfg.SECRETKEY, cfg.TOKENEXP)

	cpt := interseptors.NewInterseptors(auc, logger, cfg, srvc)

	s := grpc.NewServer(grpc.UnaryInterceptor(cpt.AuthInterceptor))
	// регистрируем сервис

	pb.RegisterUrlShortenerServer(s, &UrlShortenerServer{service: srvc, logger: logger, auc: auc})

	fmt.Println("Сервер gRPC начал работу")
	// получаем запрос gRPC
	if err := s.Serve(listen); err != nil {
		log.Fatal(err)
	}
}
