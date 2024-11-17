package services

import (
	"context"
	"strconv"
	"testing"

	"github.com/GlebZigert/url_shortener.git/internal/config"
	"github.com/GlebZigert/url_shortener.git/internal/file_reader"
	"github.com/GlebZigert/url_shortener.git/internal/logger"
	"github.com/GlebZigert/url_shortener.git/internal/storager"
)

func BenchmarkSimplest(b *testing.B) {

	cfg, err := config.NewConfig("prog", []string{}, file_reader.New())
	if err != nil {
		b.Errorf("error parse config")
	}
	ctx := context.Background()

	store := storager.New(cfg)

	logger := logger.NewLogrusLogger(cfg.FlagLogLevel, ctx)

	service := NewService(logger, store)
	for i := 0; i < 100; i++ {
		_, err = service.Short(strconv.Itoa(i), 0)
		if err != nil {
			b.Log(err.Error())
		}
	}

}
