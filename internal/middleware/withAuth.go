package middleware

import (
	"context"
	"fmt"
	"net/http"

	"github.com/GlebZigert/url_shortener.git/internal/auth"
)

func AuthMiddleware(h http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// по умолчанию устанавливаем оригинальный http.ResponseWriter как тот,
		// который будем передавать следующей функции
		ow := w

		// проверяем, что клиент умеет получать от сервера сжатые данные в формате gzip
		authv := r.Header.Get("Authorization")
		fmt.Println("auth: ", authv)

		userid, err := auth.GetUserId(authv)
		ctx := r.Context()
		if err != nil {
			fmt.Println("auth err: ", err)
			jwt, _ := auth.BuildJWTString()
			userid, _ = auth.GetUserId(jwt)
			ctx = context.WithValue(ctx, "jwt", jwt)
			ctx = context.WithValue(ctx, "new", true)
		}

		ctx = context.WithValue(ctx, "user", userid)
		r = r.WithContext(ctx)
		h.ServeHTTP(ow, r)
	}
}
