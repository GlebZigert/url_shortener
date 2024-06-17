package transport

import (
	"log"
	"net/http"

	"github.com/GlebZigert/url_shortener.git/internal/config"
	"github.com/go-chi/chi"
)

//Назначение роутера

func InitRouter() {
	r := chi.NewRouter()

	// r.Get(`/`, Get)

	r.Post(`/`, CreateShortURL)
	r.Get(`/*`, GetURL)

	log.Fatal(http.ListenAndServe(config.RunAddr, r))
}
