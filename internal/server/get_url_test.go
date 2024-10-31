package server

import (
	"context"
	"errors"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/GlebZigert/url_shortener.git/internal/auth"
	"github.com/GlebZigert/url_shortener.git/internal/config"
	"github.com/GlebZigert/url_shortener.git/internal/logger"
	"github.com/GlebZigert/url_shortener.git/internal/middleware"
	"github.com/GlebZigert/url_shortener.git/internal/packerr"
	"github.com/GlebZigert/url_shortener.git/mocks"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGetURL(t *testing.T) {

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
	//	origin_first := "aaa"
	tests := []struct {
		name     string
		request  request
		endpoint *func(w http.ResponseWriter, req *http.Request)
		want     want
	}{
		{
			name: "307 шорт для которого есть ориджин",
			request: request{
				http.MethodGet,
				"/aaa",
				-1, // в контексте ревеста не будет данных о пользователе
				nil,
			},
			want: want{
				code: http.StatusTemporaryRedirect,
			},
		},

		{
			name: "шорт который удален ",
			request: request{
				http.MethodGet,
				"/deleted",
				0,
				nil,
			},
			want: want{
				code: http.StatusGone,
			},
		},
	}
	cfg, err := config.NewConfig("prog", []string{})
	if err != nil {
		t.Errorf("error parse config")
	}
	ctx := context.Background()

	logger := logger.NewLogrusLogger(cfg.FlagLogLevel, ctx)

	ctrl := gomock.NewController(t)
	service := mocks.NewMocksrvService(ctrl)
	service.EXPECT().Origin(gomock.Any()).DoAndReturn(func(str string) (string, error) {
		t.Log("mock origin ", str)
		if str == "aaa" {
			return "bbb", nil
		}
		if str == "deleted" {
			return "", &packerr.ErrDeleted{}
		}
		return "", errors.New("отстуствует")
	}).AnyTimes()

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

			srv.GetURL(w, r)

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
