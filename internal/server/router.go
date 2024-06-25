package server

import (
	"log"
	"net/http"

	"github.com/GlebZigert/url_shortener.git/internal/config"
	"github.com/GlebZigert/url_shortener.git/internal/middleware"
	"github.com/go-chi/chi"
)

//Назначение роутера

func InitRouter() {

	r := chi.NewRouter()

	// r.Get(`/`, Get)

	r.Post(`/`, middleware.RequestLogger(middleware.GzipMiddleware(CreateShortURL)))
	r.Post(`/api/shorten`, middleware.RequestLogger(CreateShortURLfromJSON))
	r.Get(`/*`, middleware.RequestLogger(GetURL))

	log.Fatal(http.ListenAndServe(config.RunAddr, r))
}
