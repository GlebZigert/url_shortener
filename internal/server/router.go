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
		r.Use(middleware.Log)
		r.Post(`/`, middleware.Auth(middleware.Gzip(CreateShortURL)))
		r.Post(`/api/shorten`, CreateShortURLfromJSON)
		r.Post(`/api/shorten/batch`, Batcher)
		r.Get(`/api/user/urls`, middleware.Auth(GetURLs))
		r.Get(`/ping`, Ping)
		r.Get(`/*`, GetURL)
		r.Delete(`/api/user/urls`, middleware.Auth(Delete))

	})

	http.ListenAndServe(config.RunAddr, r)
}
