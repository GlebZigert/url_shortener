// Package area provides functions for calculating the area of various shapes.
package main

// @Title BucketStorage API
// @Description Сервис хранения данных bucket-ов.
// @Version 1.0

// @Contact.email support@ultimatestore.io

// @BasePath /api/v1
// @Host ultimatestore.io:8080

// @SecurityDefinitions.apikey ApiKeyAuth
// @In header
// @Name authorization

// @Tag.name Info
// @Tag.description "Группа запросов состояния сервиса"

// @Tag.name Storage
// @Tag.description "Группа для работы с данными внутри bucket-ов"

import (
	_ "net/http/pprof"

	"github.com/GlebZigert/url_shortener.git/internal/app"
)

func main() {

	app.Run()
}
