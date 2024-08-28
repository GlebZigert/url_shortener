package server

import (
	"net/http"

	"github.com/GlebZigert/url_shortener.git/internal/config"
	"github.com/GlebZigert/url_shortener.git/internal/middleware"
	"github.com/go-chi/chi"
)

func InitRouter() {

	r := chi.NewRouter()

	r.Post(`/`, middleware.AuthMiddleware(middleware.RequestLogger(middleware.GzipMiddleware(CreateShortURL))))
	r.Post(`/api/shorten`, middleware.RequestLogger(CreateShortURLfromJSON))
	r.Post(`/api/shorten/batch`, middleware.RequestLogger(Batcher))
	r.Get(`/api/user/urls`, middleware.AuthMiddleware(middleware.RequestLogger(GetURLs)))
	r.Get(`/ping`, middleware.RequestLogger(Ping))
	r.Get(`/*`, middleware.RequestLogger(GetURL))

	http.ListenAndServe(config.RunAddr, r)
}
