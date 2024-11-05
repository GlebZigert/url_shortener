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
	"github.com/GlebZigert/url_shortener.git/internal/logger"
	"github.com/GlebZigert/url_shortener.git/internal/middleware"
	"github.com/GlebZigert/url_shortener.git/internal/packerr"
	"github.com/GlebZigert/url_shortener.git/internal/services"
	"github.com/GlebZigert/url_shortener.git/internal/storager"
	"github.com/stretchr/testify/assert"
)

func TestBatcher(t *testing.T) {

	type batch struct {
		ID       string `json:"correlation_id"`
		Original string `json:"original_url"`
	}

	sreq := []batch{
		batch{"1", "11111"},
		batch{"2", "11112"},
		batch{"3", "11113"},
	}

	reqbody, err := json.Marshal(sreq)

	if err != nil {
		t.Error(err)
	}

	type want struct {
		code int
	}

	type request struct {
		body io.Reader
	}

	tests := []struct {
		name    string
		request request
		want    want
	}{{
		name: "1",
		request: request{
			strings.NewReader(string(reqbody)),
		},
		want: want{
			code: http.StatusCreated,
		},
	},
		{
			name: "nil body",
			request: request{
				nil,
			},
			want: want{
				code: http.StatusBadRequest,
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
			t.Log("req: ")

			r := httptest.NewRequest(http.MethodPost, "/api/shorten/batch", test.request.body)
			w := httptest.NewRecorder()

			var err error
			//помещаем в контекст реквеста указатель на ошибку
			packerr.AddErrToReqContext(r, &err)

			srv.Batcher(w, r)

			if err != nil {
				t.Log("err: ", err.Error())
			}

			res := w.Result()

			assert.Equal(t, test.want.code, res.StatusCode)

		})

	}
}
