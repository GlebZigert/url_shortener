package middleware

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

	"github.com/GlebZigert/url_shortener.git/internal/packerr"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGzip(t *testing.T) {
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

	cfg, err := config.NewConfig("prog", []string{})
	if err != nil {
		t.Errorf("error parse config")
	}

	ctx := context.Background()

	//store := storager.New(cfg)

	logger := logger.NewLogrusLogger(cfg.FlagLogLevel, ctx)

	//service := services.NewService(logger, store)

	auc := auth.NewAuth(cfg.SECRETKEY, cfg.TOKENEXP)
	mdl := NewMiddlewares(auc, logger)

	//srv, _ := NewServer(cfg, mdl, logger, service)

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			//t.Log("req: ", test.request.method, " ", test.request.url)

			r := httptest.NewRequest(test.request.method, test.request.url, test.request.body)
			w := httptest.NewRecorder()

			var err error
			//помещаем в контекст реквеста указатель на ошибку
			packerr.AddErrToReqContext(r, &err)

			if test.request.user >= 0 {
				ctx = auc.SetUID(ctx, test.request.user)
			}

			r = r.WithContext(ctx)

			testHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				_, err = w.Write([]byte("[]"))
			})

			if test.request.acceptEncoding {
				r.Header.Set("Accept-Encoding", "gzip")
			}

			if test.request.contentEncoding {
				r.Header.Set("Content-Encoding", "gzip")
			}
			handler := mdl.Gzip(testHandler)
			handler.ServeHTTP(w, r)

			if err != nil {
				t.Log("err: ", err.Error())
			}

			res := w.Result()

			//t.Log("res: ", res.StatusCode, " ", string(body))

			//		c.w.Header().Set("Content-Encoding", "gzip")

			contentEncoding := r.Header.Get("Content-Encoding")
			sendsGzip := strings.Contains(contentEncoding, "gzip")

			assert.Equal(t, test.want.sendsGzip, sendsGzip)

			defer packerr.AddCloseErrToErr(&err, res.Body)
			_, err = io.ReadAll(res.Body)
			require.NoError(t, err)

		})
	}
}
