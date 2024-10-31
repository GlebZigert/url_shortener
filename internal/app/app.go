package app

import (
	"context"
	"os"

	"github.com/GlebZigert/url_shortener.git/internal/auth"
	"github.com/GlebZigert/url_shortener.git/internal/config"
	"github.com/GlebZigert/url_shortener.git/internal/db"
	"github.com/GlebZigert/url_shortener.git/internal/logger"
	"github.com/GlebZigert/url_shortener.git/internal/middleware"
	"github.com/GlebZigert/url_shortener.git/internal/server"
	"github.com/GlebZigert/url_shortener.git/internal/services"
	"github.com/GlebZigert/url_shortener.git/internal/storager"
)

// Запуск
func Run() (err error) {

	cfg, err := config.NewConfig(os.Args[0], os.Args[1:])
	if err != nil {
		return
	}
	ctx := context.Background()

	err = db.Init(cfg.DatabaseDSN)
	if err != nil {
		return
	}
	store := storager.New(cfg)

	logger := logger.NewLogrusLogger(cfg.FlagLogLevel, ctx)

	service := services.NewService(logger, store)

	auc := auth.NewAuth(cfg.SECRETKEY, cfg.TOKENEXP)
	mdl := middleware.NewMiddlewares(auc, logger)
	server, err := server.NewServer(cfg, mdl, logger, service)

	if err != nil {
		return
	}
	err = server.Start()

	return
}
