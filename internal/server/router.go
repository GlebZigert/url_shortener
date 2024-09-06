package server

import (
	"net/http"

	"github.com/GlebZigert/url_shortener.git/internal/config"
	"github.com/GlebZigert/url_shortener.git/internal/middleware"
	"github.com/go-chi/chi"
)

func InitRouter() {

	r := chi.NewRouter()

	r.Post(`/`, middleware.ErrorHandler(middleware.AuthMiddleware(middleware.RequestLogger(middleware.GzipMiddleware(CreateShortURL)))))
	r.Post(`/api/shorten`, middleware.ErrorHandler(middleware.RequestLogger(CreateShortURLfromJSON)))
	r.Post(`/api/shorten/batch`, middleware.ErrorHandler(middleware.RequestLogger(Batcher)))
	r.Get(`/api/user/urls`, middleware.ErrorHandler(middleware.AuthMiddleware(middleware.RequestLogger(GetURLs))))
	r.Get(`/ping`, middleware.ErrorHandler(middleware.RequestLogger(Ping)))
	r.Get(`/*`, middleware.ErrorHandler(middleware.RequestLogger(GetURL)))
	r.Delete(`/api/user/urls`, middleware.ErrorHandler(middleware.AuthMiddleware(middleware.RequestLogger(Delete))))
	http.ListenAndServe(config.RunAddr, r)
}
