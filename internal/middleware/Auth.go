package middleware

import (
	"context"
	"net/http"

	"github.com/GlebZigert/url_shortener.git/internal/auth"
	"github.com/GlebZigert/url_shortener.git/internal/config"
	"github.com/GlebZigert/url_shortener.git/internal/logger"
	"go.uber.org/zap"
)

func Auth(h MyHandlerFunc) MyHandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) (err error) {
		// по умолчанию устанавливаем оригинальный http.ResponseWriter как тот,
		// который будем передавать следующей функции

		// проверяем, что клиент умеет получать от сервера сжатые данные в формате gzip
		authv := r.Header.Get("Authorization")
		logger.Log.Info("auth: ", zap.String("", authv))

		userid, err := auth.GetUserID(authv)
		ctx := r.Context()
		if err != nil {

			jwt, _ := auth.BuildJWTString()
			userid, _ = auth.GetUserID(jwt)
			ctx = context.WithValue(ctx, config.JWTkey, string(jwt))
			ctx = context.WithValue(ctx, config.NEWkey, bool(true))
		}

		ctx = context.WithValue(ctx, config.UIDkey, int(userid))
		r = r.WithContext(ctx)
		h(w, r)
		return
	}
}
