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
	"errors"
	"log"
	"net/http"
	_ "net/http/pprof"

	"github.com/GlebZigert/url_shortener.git/internal/app"
)

var (
	buildVersion string
	buildDate    string
	buildCommit  string
)

// func check return N/A if value is emptyt
func check(value string) string {
	if value == "" {
		return "N/A"
	}
	return value
}

func main() {

	log.Println("Build version:", check(buildVersion))
	log.Println("Build date:", check(buildDate))
	log.Println("Build commit:", check(buildCommit))

	err := app.Run()

	if err != nil && !errors.Is(err, http.ErrServerClosed) {
		log.Fatalf("HTTP server error: %v", err)
	}

}
