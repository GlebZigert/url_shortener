package grpcapp

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"

	"github.com/GlebZigert/url_shortener.git/internal/auth"
	"github.com/GlebZigert/url_shortener.git/internal/config"
	"github.com/GlebZigert/url_shortener.git/internal/db"
	"github.com/GlebZigert/url_shortener.git/internal/filereader"
	"github.com/GlebZigert/url_shortener.git/internal/grpcserver"
	"github.com/GlebZigert/url_shortener.git/internal/interseptors"
	"github.com/GlebZigert/url_shortener.git/internal/logger"
	"github.com/GlebZigert/url_shortener.git/internal/services"
	"github.com/GlebZigert/url_shortener.git/internal/storager"
	pb "github.com/GlebZigert/url_shortener.git/proto"
	"google.golang.org/grpc"
)

func Run() (err error) {

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

	pb.RegisterUrlShortenerServer(s, grpcserver.New(cfg, srvc, logger, auc))

	fmt.Println("Сервер gRPC начал работу")
	// получаем запрос gRPC
	if err := s.Serve(listen); err != nil {
		log.Fatal(err)
	}

	return
}
