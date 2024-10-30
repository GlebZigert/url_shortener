package services

import (
	"context"
	"fmt"

	"github.com/GlebZigert/url_shortener.git/internal/config"
	"github.com/GlebZigert/url_shortener.git/internal/db"
	"github.com/GlebZigert/url_shortener.git/internal/logger"
	"github.com/GlebZigert/url_shortener.git/internal/storager"
)

func Example() {

	//Инициируем компоненты сервиса
	cfg, err := config.NewConfig("prog", []string{})
	if err != nil {
		fmt.Errorf("error parse config")
	}
	ctx := context.Background()

	db.Init(cfg.DatabaseDSN)
	store := storager.New(cfg)

	logger := logger.NewLogrusLogger(cfg.FlagLogLevel, ctx)

	service := NewService(logger, store)

	originFirst := "originFirst"
	uid1 := 1

	//Метод Short на входе принимает оригинальный url и id пользователя
	// создает и возвращает сокращенный url либо ошибку

	short, err := service.Short(originFirst, uid1)

	if err != nil {
		fmt.Println("ошибка при генерации сокращенного url")
		return
	}

	//Метод Origin на входе принимает сокращенный url
	// возвращает оригинальный url либо ошибку

	origin, err := service.Origin(short)
	fmt.Println("Метод Origin принимает сокращенный URL возвращает оригинальный URL: ", origin)

}
