package services

import (
	"context"
	"strconv"
	"testing"

	"github.com/GlebZigert/url_shortener.git/internal/config"
	"github.com/GlebZigert/url_shortener.git/internal/db"
	"github.com/GlebZigert/url_shortener.git/internal/logger"
	"github.com/GlebZigert/url_shortener.git/internal/storager"
)

func BenchmarkSimplest(b *testing.B) {

	cfg, err := config.NewConfig("prog", []string{})
	if err != nil {
		b.Errorf("error parse config")
	}
	ctx := context.Background()

	db.Init(cfg.DatabaseDSN)
	store := storager.New(cfg)

	logger := logger.NewLogrusLogger(cfg.FlagLogLevel, ctx)

	service := NewService(logger, store)
	for i := 0; i < 100; i++ {
		service.Short(strconv.Itoa(i), 0)
	}

}
