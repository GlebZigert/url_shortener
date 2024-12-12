package middleware

import (
	"context"
	"errors"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/GlebZigert/url_shortener.git/internal/auth"
	"github.com/GlebZigert/url_shortener.git/internal/config"
	"github.com/GlebZigert/url_shortener.git/internal/filereader"
	"github.com/GlebZigert/url_shortener.git/internal/logger"

	"github.com/GlebZigert/url_shortener.git/internal/packerr"
)

func TestErr(t *testing.T) {
	type want struct {
		code        int
		contentType string
		sendsGzip   bool
	}

	type request struct {
		method          string
		url             string
		user            int
		body            io.Reader
		acceptEncoding  bool
		contentEncoding bool
	}
	originFirst := "asefsfsdfjj"
	tests := []struct {
		name     string
		request  request
		endpoint *func(w http.ResponseWriter, req *http.Request)
		want     want
	}{
		{
			name: "",
			request: request{
				http.MethodGet,
				"/",
				-1, // в контексте ревеста не будет данных о пользователе
				strings.NewReader(originFirst),
				true,
				true,
			},
			want: want{
				code:      http.StatusOK,
				sendsGzip: true,
			},
		},
	}

	cfg, err := config.NewConfig("prog", []string{}, filereader.New())
	if err != nil {
		t.Errorf("error parse config")
	}

	ctx := context.Background()

	//store := storager.New(cfg, filereader.New())

	logger := logger.NewLogrusLogger(cfg.FlagLogLevel, ctx)

	auc := auth.NewAuth(cfg.SECRETKEY, cfg.TOKENEXP)
	mdl := NewMiddlewares(auc, logger, cfg, nil)

	//srv, _ := NewServer(cfg, mdl, logger, service)

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			//t.Log("req: ", test.request.method, " ", test.request.url)

			r := httptest.NewRequest(test.request.method, test.request.url, test.request.body)
			w := httptest.NewRecorder()

			if test.request.user >= 0 {
				ctx = auc.SetUID(ctx, test.request.user)
			}
			var err error
			packerr.AddErrToReqContext(r, &err)
			r = r.WithContext(ctx)

			testHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				_, err = w.Write([]byte("[]"))
				err = errors.New("some err")
				packerr.AddErrToReqContext(r, &err)
			})

			if test.request.acceptEncoding {
				r.Header.Set("Accept-Encoding", "gzip")
			}

			if test.request.contentEncoding {
				r.Header.Set("Content-Encoding", "gzip")
			}
			handler := mdl.ErrHandler(testHandler)
			handler.ServeHTTP(w, r)

		})
	}
}
