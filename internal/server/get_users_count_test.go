package server

import (
	"context"
	"net/http"
	"net/http/httptest"
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
	"gotest.tools/v3/assert"
)

//несколько раз выполним авторизацию
//проверим количество пользователей

func TestUserCount(t *testing.T) {

	tests := []struct {
		name  string
		users int
	}{{
		name:  "трое пользователей",
		users: 3,
	}, {
		name:  "пятеро пользователей",
		users: 5,
	},
	}

	for _, test := range tests {

		t.Run(test.name, func(t *testing.T) {

			cfg, err := config.NewConfig("prog", []string{}, filereader.New())
			if err != nil {
				t.Errorf("error parse config")
			}
			ctx := context.Background()

			logger := logger.NewLogrusLogger(cfg.FlagLogLevel, ctx)

			auc := auth.NewAuth(cfg.SECRETKEY, cfg.TOKENEXP)

			dber := db.Get(db.Init(cfg.GetBaseURL()))

			store := storager.New(cfg, nil, nil)

			service := services.NewService(logger, store)
			mdl := middleware.NewMiddlewares(auc, logger, cfg, service)

			srv, err := NewServer(cfg, mdl, logger, service, dber)
			if err != nil {
				t.Error(err)
			}

			testHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

				_, err = w.Write([]byte("[]"))
				if err != nil {
					packerr.AddErrToReqContext(r, &err)
				}

			})

			r := httptest.NewRequest(http.MethodGet, "/", nil)
			w := httptest.NewRecorder()

			handler := srv.mdl.Auth(testHandler)

			for i := 0; i < test.users; i++ {
				handler.ServeHTTP(w, r)
			}

			t.Log("srv.service.GetUsersCount(): ", srv.service.GetUsersCount())

			assert.Equal(t, test.users, srv.service.GetUsersCount())

		})

	}

}
