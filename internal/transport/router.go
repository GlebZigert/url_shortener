package transport

import (
	"log"
	"net/http"

	"github.com/GlebZigert/url_shortener.git/internal/config"
	"github.com/GlebZigert/url_shortener.git/internal/logger"
	"github.com/go-chi/chi"
)

//Назначение роутера

func InitRouter() {

	r := chi.NewRouter()

	// r.Get(`/`, Get)

	r.Post(`/`, logger.RequestLogger(CreateShortURL))
	r.Post(`/api/shorten`, logger.RequestLogger(CreateShortURLfromJSON))
	r.Get(`/*`, logger.RequestLogger(GetURL))

	log.Fatal(http.ListenAndServe(config.RunAddr, r))
}
