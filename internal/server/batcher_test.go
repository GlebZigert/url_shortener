package server

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/GlebZigert/url_shortener.git/internal/auth"
	"github.com/GlebZigert/url_shortener.git/internal/config"
	"github.com/GlebZigert/url_shortener.git/internal/db"
	"github.com/GlebZigert/url_shortener.git/internal/filereader"
	"github.com/GlebZigert/url_shortener.git/internal/logger"
	"github.com/GlebZigert/url_shortener.git/internal/middleware"
	"github.com/GlebZigert/url_shortener.git/internal/packerr"
	"github.com/GlebZigert/url_shortener.git/internal/services"
	"github.com/GlebZigert/url_shortener.git/internal/storager"
	"github.com/stretchr/testify/assert"
)

// TestBatcher
func TestBatcher(t *testing.T) {

	//test struct
	type batch struct {
		ID       string `json:"correlation_id"`
		Original string `json:"original_url"`
	}

	sreq := []batch{
		{"1", "11111"},
		{"2", "11112"},
		{"3", "11113"},
	}

	reqbody, err := json.Marshal(sreq)

	if err != nil {
		t.Error(err)
	}

	tests := []struct {
		name string
		body io.Reader
		code int
	}{
		{
			name: "1",

			body: strings.NewReader(string(reqbody)),

			code: http.StatusCreated,
		},
		{
			name: "nil body",

			body: nil,

			code: http.StatusBadRequest,
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

	auc := auth.NewAuth(cfg.SECRETKEY, cfg.TOKENEXP)
	mdl := middleware.NewMiddlewares(auc, logger)

	srv, err := NewServer(cfg, mdl, logger, service, dber)

	if err != nil {
		t.Error(err)
	}

	for _, test := range tests {

		t.Run(test.name, func(t *testing.T) {

			r := httptest.NewRequest(http.MethodPost, "/api/shorten/batch", test.body)
			w := httptest.NewRecorder()

			var errCtx error

			packerr.AddErrToReqContext(r, &errCtx)

			srv.Batcher(w, r)

			if errCtx != nil {
				t.Log("err: ", errCtx.Error())
			}

			res := w.Result()
			closeErr := res.Body.Close()
			if closeErr != nil {
				return
			}
			assert.Equal(t, test.code, res.StatusCode)

		})

	}
}
