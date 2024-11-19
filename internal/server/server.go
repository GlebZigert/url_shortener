package server

import (
	"context"
	"crypto/tls"
	"net/http"

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
	GetENABLEHTTPSflag() bool
}

type srvMiddleware interface {
	Auth(h http.Handler) http.Handler
	ErrHandler(f http.Handler) http.Handler
	Log(h http.Handler) http.Handler
	GetUserID(tokenString string) (int, error)
	CheckUID(ctx context.Context) (user int, ok bool)
	SetUID(ctx context.Context, user int) context.Context
	CheckNewFlag(ctx context.Context) (fl bool, ok bool)
	SetNewFlag(ctx context.Context, fl bool) context.Context
	Gzip(h http.HandlerFunc) http.HandlerFunc
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

// сервер
type Server struct {
	cfg     srvConfig
	mdl     srvMiddleware
	logger  srvLogger
	service srvService
}

// var errNoAuthMiddleware = errors.New("в миддлеварах не определен auth")
// конструктор сервера
func NewServer(cfg srvConfig, mdl srvMiddleware, logger srvLogger, service srvService) (*Server, error) {

	/*
		auch := mdl.GetAuch()
		if auch == nil {
			return nil, errNoAuthMiddleware
		}
	*/

	return &Server{cfg, mdl, logger, service}, nil
}

// запуск сервера
func (srv *Server) Start() (err error) {

	r := chi.NewRouter()
	r.Group(func(r chi.Router) {
		r.Use(srv.mdl.ErrHandler)
		r.Use(srv.mdl.Log)

		r.Group(func(r chi.Router) {
			r.Use(srv.mdl.Auth)
			r.Post(`/`, srv.mdl.Gzip(srv.CreateShortURL))
			r.Get(`/api/user/urls`, srv.GetURLs)
			r.Delete(`/api/user/urls`, srv.Delete)
		})

		r.Post(`/api/shorten`, srv.CreateShortURLfromJSON)
		r.Post(`/api/shorten/batch`, srv.Batcher)

		r.Get(`/ping`, srv.Ping)
		r.Get(`/*`, srv.GetURL)

	})

	if srv.cfg.GetENABLEHTTPSflag() {
		// конструируем менеджер TLS-сертификатов

		certFile := "cert.pem" // Your certificate file
		keyFile := "key.pem"   // Your private key file

		server := &http.Server{
			Addr:    ":443",
			Handler: r,
			TLSConfig: &tls.Config{
				MinVersion:               tls.VersionTLS12,
				CurvePreferences:         []tls.CurveID{tls.CurveP521, tls.CurveP384, tls.CurveP256},
				PreferServerCipherSuites: true,
				CipherSuites: []uint16{
					tls.TLS_ECDHE_RSA_WITH_AES_256_GCM_SHA384,
					tls.TLS_ECDHE_RSA_WITH_AES_128_GCM_SHA256,
				},
			},
		}

		err = server.ListenAndServeTLS(certFile, keyFile)
	} else {
		err = http.ListenAndServe(srv.cfg.GetRunAddr(), r)
	}

	return
}
