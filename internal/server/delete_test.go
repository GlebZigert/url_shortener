package server

import (
	"bytes"
	"context"
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
	"github.com/GlebZigert/url_shortener.git/mocks"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestDelete(t *testing.T) {

	cfg, err := config.NewConfig("prog", []string{}, filereader.New())
	if err != nil {
		t.Errorf("error parse config")
	}

	ctx := context.Background()

	store := storager.New(cfg)

	logger := logger.NewLogrusLogger(cfg.FlagLogLevel, ctx)

	service := services.NewService(logger, store)

	ctrl := gomock.NewController(t)
	auc := mocks.NewMockmdlAuth(ctrl)

	auc.EXPECT().CheckNewFlag(gomock.Any()).DoAndReturn(func(ctx context.Context) (bool, bool) {
		return true, true
	}).AnyTimes()

	mdl := middleware.NewMiddlewares(auc, logger)

	srv, err := NewServer(cfg, mdl, logger, service)

	if err != nil {
		t.Error(err)
	}

	t.Run("service delete endpoint", func(t *testing.T) {

		r := httptest.NewRequest(http.MethodDelete, "/api/user/urls", nil)
		w := httptest.NewRecorder()

		srv.Delete(w, r)

		res := w.Result()
		closeErr := res.Body.Close()
		if closeErr != nil {
			return
		}

		if err != nil {
			t.Log(err.Error())
		}

		assert.Equal(t, http.StatusBadRequest, res.StatusCode)

	})
}

func TestDelete1(t *testing.T) {

	cfg, err := config.NewConfig("prog", []string{}, filereader.New())
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

	todel := "[\"1\", \"2\", \"3\"]"

	t.Run("service delete endpoint", func(t *testing.T) {

		r := httptest.NewRequest(http.MethodDelete, "/api/user/urls", bytes.NewReader([]byte(todel)))
		w := httptest.NewRecorder()

		ctx := context.Background()
		ctx = auc.SetUID(ctx, 0)
		r = r.WithContext(ctx)

		var err error

		packerr.AddErrToReqContext(r, &err)

		srv.Delete(w, r)

		res := w.Result()
		closeErr := res.Body.Close()
		if closeErr != nil {
			return
		}

		if err != nil {
			t.Log(err.Error())
		}

		assert.Equal(t, http.StatusAccepted, res.StatusCode)

	})
}
