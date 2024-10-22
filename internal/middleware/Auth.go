package middleware

import (
	"context"
	"net/http"

	"github.com/GlebZigert/url_shortener.git/internal/auth"
	"github.com/GlebZigert/url_shortener.git/internal/config"
	"github.com/GlebZigert/url_shortener.git/internal/logger"
	"github.com/GlebZigert/url_shortener.git/internal/packerr"
	"go.uber.org/zap"
)

func Auth(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var err error
		defer packerr.AddErrToReqContext(r, &err)
		// по умолчанию устанавливаем оригинальный http.ResponseWriter как тот,
		// который будем передавать следующей функции

		// проверяем, что клиент умеет получать от сервера сжатые данные в формате gzip
		authv, err := r.Cookie("Authorization") // Header.Get("Authorization")

		var userid int
		ctx := r.Context()

		if authv != nil {
			logger.Log.Info("auth: ", zap.String("", authv.Value))

			userid, err = auth.GetUserID(authv.Value)
			//ctx = r.Context()
		}

		if err != nil || authv == nil {
			jwt, _ := auth.BuildJWTString()
			userid, _ = auth.GetUserID(jwt)
			ctx = context.WithValue(ctx, config.JWTkey, string(jwt))
			ctx = context.WithValue(ctx, config.NEWkey, bool(true))

			//	w.Header().Add("Authorization", string(jwt))
			cookie := http.Cookie{
				Name:     "Authorization",
				Value:    string(jwt),
				Path:     "/",
				HttpOnly: true,
			}
			http.SetCookie(w, &cookie)

		}

		ctx = context.WithValue(ctx, config.UIDkey, int(userid))
		r = r.WithContext(ctx)
		h.ServeHTTP(w, r)

	})
}
