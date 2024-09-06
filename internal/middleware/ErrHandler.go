package middleware

import (
	"net/http"

	"github.com/GlebZigert/url_shortener.git/internal/logger"
	"go.uber.org/zap"
)

func ErrHandler(f func(http.ResponseWriter, *http.Request) error) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		err := f(w, r)
		if err != nil {
			logger.Log.Error("err: ", zap.String("", err.Error()))
		}
	}
}
