package server

import (
	"context"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/GlebZigert/url_shortener.git/internal/auth"
	"github.com/GlebZigert/url_shortener.git/internal/config"
	"github.com/GlebZigert/url_shortener.git/internal/filereader"
	"github.com/GlebZigert/url_shortener.git/internal/logger"
	"github.com/GlebZigert/url_shortener.git/internal/middleware"
	"github.com/GlebZigert/url_shortener.git/internal/packerr"
	"github.com/GlebZigert/url_shortener.git/internal/services"
	"github.com/GlebZigert/url_shortener.git/internal/storager"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestPing(t *testing.T) {
	type want struct {
		code        int
		contentType string
	}

	type request struct {
		method string
		url    string
	}

	tests := []struct {
		name     string
		request  request
		endpoint *func(w http.ResponseWriter, req *http.Request)
		want     want
	}{
		{
			name: "500 запрос без данных о пользователе",
			request: request{
				http.MethodGet,
				"/Ping",
			},
			want: want{
				code: http.StatusInternalServerError,
			},
		},
	}
	cfg, err := config.NewConfig("prog", []string{}, filereader.New())
	if err != nil {
		t.Errorf("error parse config")
	}
	ctx := context.Background()

	store := storager.New(cfg, filereader.New())

	logger := logger.NewLogrusLogger(cfg.FlagLogLevel, ctx)

	service := services.NewService(logger, store)

	auc := auth.NewAuth(cfg.SECRETKEY, cfg.TOKENEXP)
	mdl := middleware.NewMiddlewares(auc, logger)

	srv, err := NewServer(cfg, mdl, logger, service)
	if err != nil {
		t.Error(err)
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Log("req: ", test.request.method, " ", test.request.url)

			r := httptest.NewRequest(test.request.method, test.request.url, nil)
			w := httptest.NewRecorder()

			var err error
			//помещаем в контекст реквеста указатель на ошибку
			packerr.AddErrToReqContext(r, &err)

			r = r.WithContext(ctx)

			srv.Ping(w, r)

			if err != nil {
				t.Log("err: ", err.Error())
			}

			res := w.Result()

			body, err := io.ReadAll(res.Body)
			closeErr := res.Body.Close()
			if closeErr != nil {
				return
			}
			if err != nil {
				return //err
			}

			t.Log("res: ", res.StatusCode, " ", string(body))

			assert.Equal(t, test.want.code, res.StatusCode)

			require.NoError(t, err)

		})
	}
}
