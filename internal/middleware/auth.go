package middleware

import (
	"net/http"

	"github.com/GlebZigert/url_shortener.git/internal/packerr"
)

// мидл для атворизации
func (mdl *Middleware) Auth(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var err error
		defer packerr.AddErrToReqContext(r, &err)
		// по умолчанию устанавливаем оригинальный http.ResponseWriter как тот,
		// который будем передавать следующей функции

		authv, err := r.Cookie("Authorization") // Header.Get("Authorization")

		var userid int
		ctx := r.Context()

		if err == nil {
			mdl.logger.Info("auth: ", map[string]interface{}{
				"auth": authv,
			})

			userid, err = mdl.GetUserID(authv.Value)
			//ctx = r.Context()
		}

		if err != nil {

			userid = userid + 1
			jwt, err := mdl.BuildJWTString(userid)
			if err != nil {

				mdl.logger.Error("BuildJWTString: ", map[string]interface{}{
					"err": err.Error(),
				})

				http.Error(w, "", http.StatusInternalServerError)
				return
			}

			userid, err = mdl.GetUserID(jwt)
			if err != nil {
				mdl.logger.Error("GetUserID: ", map[string]interface{}{
					"err": err.Error(),
				})
				http.Error(w, "", http.StatusInternalServerError)
				return
			}

			ctx = mdl.SetNewFlag(ctx, true)

			//	w.Header().Add("Authorization", string(jwt))
			cookie := http.Cookie{
				Name:     "Authorization",
				Value:    string(jwt),
				Path:     "/",
				HttpOnly: true,
			}
			http.SetCookie(w, &cookie)

		}

		ctx = mdl.SetUID(ctx, userid)

		r = r.WithContext(ctx)
		h.ServeHTTP(w, r)

	})
}
