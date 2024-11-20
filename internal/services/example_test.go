package services

import (
	"context"

	"github.com/GlebZigert/url_shortener.git/internal/config"
	"github.com/GlebZigert/url_shortener.git/internal/db"
	"github.com/GlebZigert/url_shortener.git/internal/filereader"
	"github.com/GlebZigert/url_shortener.git/internal/logger"
	"github.com/GlebZigert/url_shortener.git/internal/storager"
)

func Example() {

	//Инициируем компоненты сервиса
	cfg, err := config.NewConfig("prog", []string{}, filereader.New())
	if err != nil {
		return
	}
	ctx := context.Background()
	dber := db.Get()
	err = dber.Init(cfg.DatabaseDSN)
	if err != nil {
		return
	}
	store := storager.New(cfg, filereader.New(), dber)

	logger := logger.NewLogrusLogger(cfg.FlagLogLevel, ctx)

	service := NewService(logger, store)

	originFirst := "originFirst"
	uid1 := 1

	//Метод Short на входе принимает оригинальный url и id пользователя
	// создает и возвращает сокращенный url либо ошибку

	short, err := service.Short(originFirst, uid1)

	if err != nil {
		return
	}

	//Метод Origin на входе принимает сокращенный url
	// возвращает оригинальный url либо ошибку

	_, err = service.Origin(short)

}
