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
		r.Post(`/`, middleware.Log(middleware.Auth(middleware.Gzip(CreateShortURL))))
		r.Post(`/api/shorten`, middleware.Log(CreateShortURLfromJSON))
		r.Post(`/api/shorten/batch`, middleware.Log(Batcher))
		r.Get(`/api/user/urls`, middleware.Log(middleware.Auth(GetURLs)))
		r.Get(`/ping`, middleware.Log(Ping))
		r.Get(`/*`, middleware.Log(GetURL))
		r.Delete(`/api/user/urls`, middleware.Log(middleware.Auth(Delete)))

	})

	http.ListenAndServe(config.RunAddr, r)
}
