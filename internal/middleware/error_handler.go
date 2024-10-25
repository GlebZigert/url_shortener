package middleware

import (
	"context"
	"net/http"

	"github.com/GlebZigert/url_shortener.git/internal/config"
)

// мидл для обработки ошибок
func (mdl *Middleware) ErrHandler(f http.Handler) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var err error
		//помещаем в контекст реквеста указатель на ошибку
		ctx := context.WithValue(r.Context(), config.Errkey, &err)
		r = r.WithContext(ctx)
		f.ServeHTTP(w, r)

		if err != nil {
			mdl.logger.Error("auth: ", map[string]interface{}{
				"err": err.Error(),
			})
		}
	})
}
