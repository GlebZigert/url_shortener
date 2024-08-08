package app

import (
	"github.com/GlebZigert/url_shortener.git/internal/config"
	"github.com/GlebZigert/url_shortener.git/internal/db"
	"github.com/GlebZigert/url_shortener.git/internal/logger"
	"github.com/GlebZigert/url_shortener.git/internal/server"
	"github.com/GlebZigert/url_shortener.git/internal/services"
	"go.uber.org/zap"
)

func Run() error {

	config.ParseFlags()

	if err := logger.Initialize(config.FlagLogLevel); err != nil {
		return err
	}
	db.Init()
	services.Init()
	logger.Log.Info("Running server", zap.String("address", config.RunAddr))

	server.InitRouter()

	return nil
}
