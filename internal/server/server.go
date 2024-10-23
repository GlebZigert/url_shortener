package server

import (
	"context"
	"net/http"

	"github.com/GlebZigert/url_shortener.git/internal/middleware"
	"github.com/GlebZigert/url_shortener.git/internal/storager"
	"github.com/go-chi/chi"
)

type srvConfig interface {
	GetRunAddr() string
	GetBaseURL() string
	GetFlagLogLevel() string
	GetFileStoragePath() string
	GetNumWorkers() int
	GetDatabaseDSN() string
	GetTOKENEXP() int
	GetSECRETKEY() string
}

type srvMiddleware interface {
	Auth(h http.Handler) http.Handler
	ErrHandler(f http.Handler) http.Handler
	Log(h http.Handler) http.Handler
	CheckUID(ctx context.Context) (user int, ok bool)
}

type srvLogger interface {
	Info(msg string, fields map[string]interface{})
	Error(msg string, fields map[string]interface{})
}

type srvService interface {
	Short(oririn string, uuid int) (string, error)
	Delete(shorts []string, uid int) error
	Origin(short string) (string, error)
	GetAll() *[]*storager.Shorten
}

type Server struct {
	cfg     srvConfig
	mdl     srvMiddleware
	logger  srvLogger
	service srvService
}

//var errNoAuthMiddleware = errors.New("в миддлеварах не определен auth")

func NewServer(cfg srvConfig, mdl srvMiddleware, logger srvLogger, service srvService) (*Server, error) {

	/*
		auch := mdl.GetAuch()
		if auch == nil {
			return nil, errNoAuthMiddleware
		}
	*/

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

	err = http.ListenAndServe(srv.cfg.GetRunAddr(), r)
	return
}
