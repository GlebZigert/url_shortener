package server

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/GlebZigert/url_shortener.git/internal/config"
	"github.com/GlebZigert/url_shortener.git/internal/db"
	"github.com/GlebZigert/url_shortener.git/internal/filereader"
	"github.com/GlebZigert/url_shortener.git/internal/logger"
	"github.com/GlebZigert/url_shortener.git/internal/middleware"
	"github.com/GlebZigert/url_shortener.git/internal/packerr"
	"github.com/GlebZigert/url_shortener.git/internal/services"
	"github.com/GlebZigert/url_shortener.git/internal/storager"
	"github.com/GlebZigert/url_shortener.git/mocks"
	"github.com/golang/mock/gomock"
	"gotest.tools/v3/assert"
)

func TestGetURLs(t *testing.T) {

	type request struct {
		uid int
	}

	type want struct {
		code int
	}

	tests := []struct {
		name    string
		request request
		want    want
	}{
		{
			name: " запрос который не содержит ID пользователя",
			request: request{
				uid: 0,
			},
			want: want{
				code: http.StatusUnauthorized,
			},
		},
	}

	cfg, err := config.NewConfig("prog", []string{}, filereader.New())
	if err != nil {
		t.Errorf("error parse config")
	}
	ctx := context.Background()
	dber := db.Get(db.Init(cfg.GetBaseURL()))
	store := storager.New(cfg, filereader.New(), dber)

	logger := logger.NewLogrusLogger(cfg.FlagLogLevel, ctx)

	service := services.NewService(logger, store)

	ctrl := gomock.NewController(t)
	auc := mocks.NewMockMdlAuth(ctrl)

	auc.EXPECT().CheckNewFlag(gomock.Any()).DoAndReturn(func(ctx context.Context) (bool, bool) {
		return true, true
	}).AnyTimes()

	mdl := middleware.NewMiddlewares(auc, logger, cfg, service)

	srv, err := NewServer(cfg, mdl, logger, service, dber)

	if err != nil {
		t.Error(err)
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Log("req")

			r := httptest.NewRequest(http.MethodGet, "/api/user/urls", nil)
			w := httptest.NewRecorder()

			ctx := context.Background()
			r = r.WithContext(ctx)

			var err error

			packerr.AddErrToReqContext(r, &err)

			srv.GetURLs(w, r)

			res := w.Result()
			closeErr := res.Body.Close()
			if closeErr != nil {
				return
			}

			if err != nil {
				t.Log(err.Error())
			}

			assert.Equal(t, test.want.code, res.StatusCode)
		})
	}

}

func TestGetURLs1(t *testing.T) {

	type request struct {
		uid int
	}

	type want struct {
		code int
	}

	tests := []struct {
		name    string
		request request
		want    want
	}{
		{
			name: " запрос который не содержит ID пользователя",
			request: request{
				uid: 0,
			},
			want: want{
				code: http.StatusUnauthorized,
			},
		},
	}

	cfg, err := config.NewConfig("prog", []string{}, filereader.New())
	if err != nil {
		t.Errorf("error parse config")
	}
	ctx := context.Background()
	dber := db.Get(db.Init(cfg.GetBaseURL()))
	store := storager.New(cfg, filereader.New(), dber)

	logger := logger.NewLogrusLogger(cfg.FlagLogLevel, ctx)

	service := services.NewService(logger, store)

	ctrl := gomock.NewController(t)
	auc := mocks.NewMockMdlAuth(ctrl)

	auc.EXPECT().CheckNewFlag(gomock.Any()).DoAndReturn(func(ctx context.Context) (bool, bool) {
		return false, false
	}).AnyTimes()

	auc.EXPECT().CheckUID(gomock.Any()).DoAndReturn(func(ctx context.Context) (int, bool) {
		return 0, false
	}).AnyTimes()

	mdl := middleware.NewMiddlewares(auc, logger, cfg, service)

	srv, err := NewServer(cfg, mdl, logger, service, dber)

	if err != nil {
		t.Error(err)
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Log("req")

			r := httptest.NewRequest(http.MethodGet, "/api/user/urls", nil)
			w := httptest.NewRecorder()

			ctx := context.Background()
			r = r.WithContext(ctx)

			var err error

			packerr.AddErrToReqContext(r, &err)

			srv.GetURLs(w, r)

			res := w.Result()

			closeErr := res.Body.Close()
			if closeErr != nil {
				return
			}

			if err != nil {
				t.Log(err.Error())
			}

			assert.Equal(t, test.want.code, res.StatusCode)
		})
	}

}

func TestGetURLs2(t *testing.T) {

	type request struct {
		uid int
	}

	type want struct {
		code int
	}

	tests := []struct {
		name    string
		request request
		want    want
	}{
		{
			name: " запрос который не содержит ID пользователя",
			request: request{
				uid: 0,
			},
			want: want{
				code: http.StatusOK,
			},
		},
	}

	cfg, err := config.NewConfig("prog", []string{}, filereader.New())
	if err != nil {
		t.Errorf("error parse config")
	}
	ctx := context.Background()
	dber := db.Get(db.Init(cfg.GetBaseURL()))
	store := storager.New(cfg, filereader.New(), dber)

	logger := logger.NewLogrusLogger(cfg.FlagLogLevel, ctx)

	service := services.NewService(logger, store)

	ctrl := gomock.NewController(t)
	auc := mocks.NewMockMdlAuth(ctrl)

	auc.EXPECT().CheckNewFlag(gomock.Any()).DoAndReturn(func(ctx context.Context) (bool, bool) {
		return false, false
	}).AnyTimes()

	auc.EXPECT().CheckUID(gomock.Any()).DoAndReturn(func(ctx context.Context) (int, bool) {
		return 0, true
	}).AnyTimes()

	mdl := middleware.NewMiddlewares(auc, logger, cfg, service)

	srv, err := NewServer(cfg, mdl, logger, service, dber)

	if err != nil {
		t.Error(err)
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Log("req")

			r := httptest.NewRequest(http.MethodGet, "/api/user/urls", nil)
			w := httptest.NewRecorder()

			ctx := context.Background()
			r = r.WithContext(ctx)

			var err error

			packerr.AddErrToReqContext(r, &err)

			srv.GetURLs(w, r)

			res := w.Result()

			closeErr := res.Body.Close()
			if closeErr != nil {
				return
			}

			if err != nil {
				t.Log(err.Error())
			}

			assert.Equal(t, test.want.code, res.StatusCode)
		})
	}

}
