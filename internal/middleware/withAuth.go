package middleware

import (
	"context"
	"net/http"

	"github.com/GlebZigert/url_shortener.git/internal/auth"
	"github.com/GlebZigert/url_shortener.git/internal/config"
	"github.com/GlebZigert/url_shortener.git/internal/logger"
	"go.uber.org/zap"
)

func AuthMiddleware(h http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// по умолчанию устанавливаем оригинальный http.ResponseWriter как тот,
		// который будем передавать следующей функции
		ow := w

		// проверяем, что клиент умеет получать от сервера сжатые данные в формате gzip
		authv := r.Header.Get("Authorization")
		logger.Log.Info("auth: ", zap.String("", authv))

		userid, err := auth.GetUserID(authv)
		ctx := r.Context()
		if err != nil {
			logger.Log.Error("auth err: ", zap.String("", err.Error()))

			jwt, _ := auth.BuildJWTString()
			userid, err = auth.GetUserID(jwt)
			if err != nil {
				logger.Log.Error("auth err: ", zap.String("", err.Error()))
			}
			ctx = context.WithValue(ctx, config.JWTkey, string(jwt))
			ctx = context.WithValue(ctx, config.NEWkey, bool(true))
		}

		ctx = context.WithValue(ctx, config.UIDkey, int(userid))
		r = r.WithContext(ctx)
		h.ServeHTTP(ow, r)
	}
}
