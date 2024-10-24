package server

import (
	"context"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/GlebZigert/url_shortener.git/internal/auth"
	"github.com/GlebZigert/url_shortener.git/internal/config"
	"github.com/GlebZigert/url_shortener.git/internal/db"
	"github.com/GlebZigert/url_shortener.git/internal/logger"
	"github.com/GlebZigert/url_shortener.git/internal/middleware"
	"github.com/GlebZigert/url_shortener.git/internal/packerr"
	"github.com/GlebZigert/url_shortener.git/internal/server"
	"github.com/GlebZigert/url_shortener.git/internal/services"
	"github.com/GlebZigert/url_shortener.git/internal/storager"
	"github.com/GlebZigert/url_shortener.git/mocks"
	"github.com/golang/mock/gomock"
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
	origin_first := "asefsfsdf"
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
				strings.NewReader(origin_first),
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
				strings.NewReader(origin_first),
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
				strings.NewReader(origin_first),
			},
			want: want{
				code: http.StatusConflict,
			},
		},
	}
	cfg := config.NewConfig()
	ctx := context.Background()

	db.Init(cfg.DatabaseDSN)
	store := storager.New(cfg)

	logger := logger.NewLogrusLogger(cfg.FlagLogLevel, ctx)

	service := services.NewService(logger, store)

	auc := auth.NewAuth(cfg.SECRETKEY, cfg.TOKENEXP, config.UIDkey)
	mdl := middleware.NewMiddlewares(auc, logger)

	srv, _ := server.NewServer(cfg, mdl, logger, service)

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Log("req: ", test.request.method, " ", test.request.url)

			r := httptest.NewRequest(test.request.method, test.request.url, test.request.body)
			w := httptest.NewRecorder()

			var err error
			//помещаем в контекст реквеста указатель на ошибку
			ctx := context.WithValue(r.Context(), config.Errkey, &err)

			if test.request.user >= 0 {
				ctx = context.WithValue(ctx, config.UIDkey, test.request.user)
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
			defer res.Body.Close()
			_, err = io.ReadAll(res.Body)
			require.NoError(t, err)

		})
	}
}

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
		/*
			{
				name: "409 запрос с повторным ориджином",
				request: request{
					http.MethodGet,
					"/",
					0,
					strings.NewReader(origin_first),
				},
				want: want{
					code: http.StatusConflict,
				},
			},
		*/
	}
	cfg := config.NewConfig()
	ctx := context.Background()

	db.Init(cfg.DatabaseDSN)

	logger := logger.NewLogrusLogger(cfg.FlagLogLevel, ctx)

	ctrl := gomock.NewController(t)
	service := mocks.NewMocksrvService(ctrl)
	service.EXPECT().Origin(gomock.Any()).DoAndReturn(func(str string) (string, error) {
		fmt.Println("mock origin ", str)
		if str == "aaa" {
			return "bbb", nil
		}
		if str == "deleted" {
			return "", &packerr.ErrDeleted{}
		}
		return "", errors.New("отстуствует")
	}).AnyTimes()

	auc := auth.NewAuth(cfg.SECRETKEY, cfg.TOKENEXP, config.UIDkey)
	mdl := middleware.NewMiddlewares(auc, logger)

	srv, _ := server.NewServer(cfg, mdl, logger, service)

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {

			t.Log("req: ", test.request.method, " ", test.request.url)

			r := httptest.NewRequest(test.request.method, test.request.url, test.request.body)
			w := httptest.NewRecorder()

			var err error
			//помещаем в контекст реквеста указатель на ошибку
			ctx := context.WithValue(r.Context(), config.Errkey, &err)

			if test.request.user >= 0 {
				ctx = context.WithValue(ctx, config.UIDkey, test.request.user)
			}

			r = r.WithContext(ctx)

			srv.GetURL(w, r)

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
			defer res.Body.Close()
			_, err = io.ReadAll(res.Body)
			require.NoError(t, err)

		})
	}
}
