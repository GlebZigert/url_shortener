package middleware

import (
	"net/http"

	"github.com/GlebZigert/url_shortener.git/internal/logger"
	"go.uber.org/zap"
)

type MyHandlerFunc func(w http.ResponseWriter, r *http.Request) error

// Implement the http.Handler interface.
func (fn MyHandlerFunc) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	err := fn(w, r) // Call handler function.
	if err == nil {
		return
	}
}

func ErrHandler(f MyHandlerFunc) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		err := f(w, r)
		if err != nil {
			logger.Log.Error("err: ", zap.String("", err.Error()))
		}
	}
}
