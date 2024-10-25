package middleware

import (
	"context"
	"net/http"

	"github.com/GlebZigert/url_shortener.git/internal/config"
	"github.com/GlebZigert/url_shortener.git/internal/packerr"
)

func (mdl *Middleware) Auth(h http.Handler) http.Handler {
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
			mdl.logger.Info("auth: ", map[string]interface{}{
				"auth": authv,
			})

			userid, err = mdl.GetUserID(authv.Value)
			//ctx = r.Context()
		}

		if err != nil || authv == nil {
			jwt, _ := mdl.BuildJWTString(userid)
			userid, _ = mdl.GetUserID(jwt)
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
