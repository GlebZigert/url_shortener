package server

import (
	"fmt"
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
	fmt.Println("running on", config.RunAddr)

	r.Post(`/`, middleware.RequestLogger(middleware.GzipMiddleware(CreateShortURL)))
	r.Post(`/api/shorten`, middleware.RequestLogger(CreateShortURLfromJSON))
	r.Post(`/api/shorten/batch`, middleware.RequestLogger(Batcher))
	r.Get(`/api/user/urls`, middleware.RequestLogger(GetURLs))
	r.Get(`/ping`, middleware.RequestLogger(Ping))
	r.Get(`/*`, middleware.RequestLogger(GetURL))

	log.Fatal(http.ListenAndServe(config.RunAddr, r))
}
