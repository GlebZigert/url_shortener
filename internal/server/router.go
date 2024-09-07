package server

import (
	"net/http"

	"github.com/GlebZigert/url_shortener.git/internal/config"
	"github.com/GlebZigert/url_shortener.git/internal/middleware"
	"github.com/uptrace/bunrouter"
)

func InitRouter() {

	r := chi.NewRouter()


	group.POST(`/`, middleware.Auth(middleware.Log(middleware.Gzip(CreateShortURL))))
	group.POST(`/api/shorten`, middleware.Log(CreateShortURLfromJSON))
	group.POST(`/api/shorten/batch`, middleware.Log(Batcher))
	group.GET(`/api/user/urls`, middleware.Auth(middleware.Log(GetURLs)))
	group.GET(`/ping`, middleware.Log(Ping))
	group.GET(`/*path`, middleware.Log(GetURL))
	group.DELETE(`/api/user/urls`, middleware.Auth(middleware.Log(Delete)))

	http.ListenAndServe(config.RunAddr, router)
}
