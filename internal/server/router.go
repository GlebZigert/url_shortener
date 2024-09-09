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

		r.Group(func(r chi.Router) {
			r.Use(middleware.Auth)
			r.Post(`/`, middleware.Gzip(CreateShortURL))
			r.Get(`/api/user/urls`, GetURLs)
			r.Delete(`/api/user/urls`, Delete)
		})

		r.Post(`/api/shorten`, CreateShortURLfromJSON)
		r.Post(`/api/shorten/batch`, Batcher)

		r.Get(`/ping`, Ping)
		r.Get(`/*`, GetURL)

	})

	http.ListenAndServe(config.RunAddr, r)
}
