package app

import (
	"github.com/GlebZigert/url_shortener.git/internal/config"
	"github.com/GlebZigert/url_shortener.git/internal/logger"
	"github.com/GlebZigert/url_shortener.git/internal/services"
	"github.com/GlebZigert/url_shortener.git/internal/transport"
	"go.uber.org/zap"
)

func Run() error {

	config.ParseFlags()

	if err := logger.Initialize(config.FlagLogLevel); err != nil {
		return err
	}

	services.Init()
	logger.Log.Info("Running server", zap.String("address", config.RunAddr))

	transport.InitRouter()

	return nil
}
