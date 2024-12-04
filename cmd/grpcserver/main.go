package main

import (
	// ...
	"context"
	"fmt"
	"log"
	"net"
	"os"

	"github.com/GlebZigert/url_shortener.git/internal/config"
	"github.com/GlebZigert/url_shortener.git/internal/db"
	"github.com/GlebZigert/url_shortener.git/internal/filereader"
	"github.com/GlebZigert/url_shortener.git/internal/logger"
	"github.com/GlebZigert/url_shortener.git/internal/services"
	"github.com/GlebZigert/url_shortener.git/internal/storager"
	"google.golang.org/grpc"

	// импортируем пакет со сгенерированными protobuf-файлами
	pb "github.com/GlebZigert/url_shortener.git/proto"
)

// UsersServer поддерживает все необходимые методы сервера.
type UrlShortenerServer struct {
	// нужно встраивать тип pb.Unimplemented<TypeName>
	// для совместимости с будущими версиями
	pb.UnimplementedUrlShortenerServer

	// используем sync.Map для хранения пользователей
	service *services.Service
}

func (s *UrlShortenerServer) AddUser(ctx context.Context, in *pb.CreateShortURLRequest) (*pb.CreateShortURLResponse, error) {
	var response pb.CreateShortURLResponse

	res, err := s.service.Short(in.Origin, 0)
	response.Short = res
	response.Error = err.Error()
	return &response, nil
}

func main() {
	// определяем порт для сервера
	listen, err := net.Listen("tcp", ":3200")
	if err != nil {
		log.Fatal(err)
	}
	// создаём gRPC-сервер без зарегистрированной службы
	s := grpc.NewServer()
	// регистрируем сервис

	cfg, err := config.NewConfig(os.Args[0], os.Args[1:], filereader.New())
	if err != nil {
		return
	}

	ctx := context.Background()
	dber := db.Get(db.Init(cfg.GetDatabaseDSN()))
	store := storager.New(cfg, filereader.New(), dber)

	logger := logger.NewLogrusLogger(cfg.FlagLogLevel, ctx)

	srvc := services.NewService(logger, store)

	pb.RegisterUrlShortenerServer(s, &UrlShortenerServer{service: srvc})

	fmt.Println("Сервер gRPC начал работу")
	// получаем запрос gRPC
	if err := s.Serve(listen); err != nil {
		log.Fatal(err)
	}
}
