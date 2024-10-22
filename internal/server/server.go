package server

import (
	"errors"
	"net/http"

	"github.com/GlebZigert/url_shortener.git/internal/config"
	"github.com/GlebZigert/url_shortener.git/internal/logger"
	"github.com/GlebZigert/url_shortener.git/internal/middleware"
	"github.com/GlebZigert/url_shortener.git/internal/services"
	"github.com/go-chi/chi"
)

type Server struct {
	cfg     *config.Config
	mdl     *middleware.Middleware
	logger  logger.Logger
	service *services.Service
}

var errNoAuthMiddleware = errors.New("в миддлеварах не определен auth")

func NewServer(cfg *config.Config, mdl *middleware.Middleware, logger logger.Logger, service *services.Service) (*Server, error) {
	auch := mdl.GetAuch()
	if auch == nil {
		return nil, errNoAuthMiddleware
	}
	return &Server{cfg, mdl, logger, service}, nil
}

func (srv *Server) Start() (err error) {

	r := chi.NewRouter()
	r.Group(func(r chi.Router) {
		r.Use(srv.mdl.ErrHandler)
		r.Use(srv.mdl.Log)

		r.Group(func(r chi.Router) {
			r.Use(srv.mdl.Auth)
			r.Post(`/`, middleware.Gzip(srv.CreateShortURL))
			r.Get(`/api/user/urls`, srv.GetURLs)
			r.Delete(`/api/user/urls`, srv.Delete)
		})

		r.Post(`/api/shorten`, srv.CreateShortURLfromJSON)
		r.Post(`/api/shorten/batch`, srv.Batcher)

		r.Get(`/ping`, srv.Ping)
		r.Get(`/*`, srv.GetURL)

	})

	err = http.ListenAndServe(srv.cfg.RunAddr, r)
	return
}
