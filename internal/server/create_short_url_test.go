package server

import (
	"context"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/GlebZigert/url_shortener.git/internal/auth"
	"github.com/GlebZigert/url_shortener.git/internal/config"
	"github.com/GlebZigert/url_shortener.git/internal/logger"
	"github.com/GlebZigert/url_shortener.git/internal/middleware"
	"github.com/GlebZigert/url_shortener.git/internal/packerr"
	"github.com/GlebZigert/url_shortener.git/internal/services"
	"github.com/GlebZigert/url_shortener.git/internal/storager"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestCreateShortURL(t *testing.T) {
	type want struct {
		code        int
		contentType string
	}

	type request struct {
		method string
		url    string
		user   int
		body   io.Reader
	}
	originFirst := "asefsfsdf"
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
				"/",
				-1, // в контексте ревеста не будет данных о пользователе
				strings.NewReader(originFirst),
			},
			want: want{
				code: http.StatusInternalServerError,
			},
		},
		{
			name: "201 запрос с новым ориджином",
			request: request{
				http.MethodGet,
				"/",
				0,
				strings.NewReader(originFirst),
			},
			want: want{
				code: http.StatusCreated,
			},
		},
		{
			name: "409 запрос с повторным ориджином",
			request: request{
				http.MethodGet,
				"/",
				0,
				strings.NewReader(originFirst),
			},
			want: want{
				code: http.StatusConflict,
			},
		},
	}
	cfg, err := config.NewConfig("prog", []string{})
	if err != nil {
		t.Errorf("error parse config")
	}
	ctx := context.Background()

	store := storager.New(cfg)

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

			r := httptest.NewRequest(test.request.method, test.request.url, test.request.body)
			w := httptest.NewRecorder()

			var err error
			//помещаем в контекст реквеста указатель на ошибку
			packerr.AddErrToReqContext(r, &err)

			if test.request.user >= 0 {
				ctx = auc.SetUID(ctx, test.request.user)
			}

			r = r.WithContext(ctx)

			srv.CreateShortURL(w, r)

			if err != nil {
				t.Log("err: ", err.Error())
			}

			res := w.Result()

			body, err := io.ReadAll(res.Body)
			if err != nil {
				return //err
			}

			t.Log("res: ", res.StatusCode, " ", string(body))

			assert.Equal(t, test.want.code, res.StatusCode)
			defer packerr.AddCloseErrToErr(&err, res.Body)
			_, err = io.ReadAll(res.Body)
			require.NoError(t, err)

		})
	}
}
