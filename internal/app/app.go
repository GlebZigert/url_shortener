package app

import (
	"log"
	"net/http"

	"github.com/GlebZigert/url_shortener.git/internal/config"
	"github.com/GlebZigert/url_shortener.git/internal/transport"
	"github.com/go-chi/chi"
)

func Run() {

	r := chi.NewRouter()

	// r.Get(`/`, Get)

	r.Post(`/`, transport.Post)
	r.Get(`/*`, transport.Get)

	log.Fatal(http.ListenAndServe(config.RunAddr, r))
}
