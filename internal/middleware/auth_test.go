package middleware

import (
	"context"
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

	"github.com/GlebZigert/url_shortener.git/internal/packerr"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestAuth(t *testing.T) {
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
		auth            bool
		jwt             string
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
				false,
				"",
			},
			want: want{
				code:      http.StatusOK,
				sendsGzip: false,
			},
		},

		{
			name: "",
			request: request{
				http.MethodGet,
				"/",
				-1, // в контексте ревеста не будет данных о пользователе
				strings.NewReader(originFirst),
				true,
				true,
				true,
				"wrong",
			},
			want: want{
				code:      http.StatusOK,
				sendsGzip: false,
			},
		},
	}

	cfg, err := config.NewConfig("prog", []string{})

	if err != nil {
		t.Errorf("error parse config")
	}

	ctx := context.Background()

	db.Init(cfg.DatabaseDSN)
	//store := storager.New(cfg)

	logger := logger.NewLogrusLogger(cfg.FlagLogLevel, ctx)

	//service := services.NewService(logger, store)

	//заменить на мок
	auc := auth.NewAuth(cfg.SECRETKEY, cfg.TOKENEXP)

	mdl := NewMiddlewares(auc, logger)

	//srv, _ := NewServer(cfg, mdl, logger, service)

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

			testHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				fmt.Println("testHandler")
				_, err = w.Write([]byte("[]"))
			})

			if test.request.auth {
				cookie := http.Cookie{
					Name:     "Authorization",
					Value:    string(test.request.jwt),
					Path:     "/",
					HttpOnly: true,
				}
				r.AddCookie(&cookie)
			}

			handler := mdl.Auth(testHandler)
			handler.ServeHTTP(w, r)

			if err != nil {
				t.Log("err: ", err.Error())
			}

			res := w.Result()
			auth := ""
			cookies := res.Cookies()
			for _, c := range cookies {
				if c.Name == "Authorization" {
					// Found! Use it!
					fmt.Println(c.Value) // The cookie's value
					auth = c.Value
				}
			}
			//	authv, err := res.Cookies("Authorization")

			if err != nil {
				t.Error("!!! ", err.Error())
			}

			if auth == "" {
				assert.Equal(t, http.StatusInternalServerError, res.Status)

			}

			//t.Log("res: ", res.StatusCode, " ", string(body))

			//		c.w.Header().Set("Content-Encoding", "gzip")

			contentEncoding := r.Header.Get("Content-Encoding")
			sendsGzip := strings.Contains(contentEncoding, "gzip")

			assert.Equal(t, test.want.sendsGzip, sendsGzip)

			defer res.Body.Close()
			_, err = io.ReadAll(res.Body)
			require.NoError(t, err)

		})
	}
}
