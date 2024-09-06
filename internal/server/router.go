package server

import (
	"net/http"

	"github.com/GlebZigert/url_shortener.git/internal/config"
	"github.com/GlebZigert/url_shortener.git/internal/middleware"
	"github.com/go-chi/chi"
)

func InitRouter() {

	r := chi.NewRouter()
	r.Group(func(r chi.Router) {
		r.Use(middleware.ErrHandler)

	})

	r.Post(`/`, middleware.ErrHandler(middleware.Auth(middleware.Log(middleware.Gzip(CreateShortURL)))))
	r.Post(`/api/shorten`, middleware.ErrHandler(middleware.Log(CreateShortURLfromJSON)))
	r.Post(`/api/shorten/batch`, middleware.ErrHandler(middleware.Log(Batcher)))
	r.Get(`/api/user/urls`, middleware.ErrHandler(middleware.Auth(middleware.Log(GetURLs))))
	r.Get(`/ping`, middleware.ErrHandler(middleware.Log(Ping)))
	r.Get(`/*`, middleware.ErrHandler(middleware.Log(GetURL)))
	r.Delete(`/api/user/urls`, middleware.ErrHandler(middleware.Auth(middleware.Log(Delete))))

	http.ListenAndServe(config.RunAddr, r)
}
