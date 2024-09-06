package middleware

import (
	"context"
	"net/http"

	"github.com/GlebZigert/url_shortener.git/internal/config"
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

func ErrHandler(f http.HandlerFunc) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		var err error
		ctx = context.WithValue(ctx, config.ErrKey, &err)
		r = r.WithContext(ctx)
		f(w, r)

		if err != nil {

			logger.Log.Error("user id: ", zap.String("", err.Error()))
		}
	}
}
