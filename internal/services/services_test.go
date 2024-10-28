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
	cfg := config.NewConfig("prog", []string{})
	ctx := context.Background()

	db.Init(cfg.DatabaseDSN)
	store := storager.New(cfg)

	logger := logger.NewLogrusLogger(cfg.FlagLogLevel, ctx)

	service := NewService(logger, store)

	originFirst := "originFirst"
	uid1 := 1

	//создаем шорт
	short, err := service.Short(originFirst, uid1)

	if err != nil {
		fmt.Println("что-то пошло не так")
		return
	}
	fmt.Println("Метод Short прринимает на входе оригинальный URL: ", originFirst)
	fmt.Println("Создает и возвращает сокращенный url: ", short)

	origin, err := service.Origin(short)
	fmt.Println("Метод Origin принимает сокращенный URL возвращает оригинальный URL: ", origin)

}
