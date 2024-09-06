package middleware

import (
	"context"
	"net/http"

	"github.com/GlebZigert/url_shortener.git/internal/config"
	"github.com/GlebZigert/url_shortener.git/internal/logger"
	"github.com/uptrace/bunrouter"
	"go.uber.org/zap"
)

func ErrHandler(next bunrouter.HandlerFunc) bunrouter.HandlerFunc {
	return func(w http.ResponseWriter, req bunrouter.Request) error {
		err := next(w, req)

		switch err := err.(type) {
		case nil:
			// no error

		default:
			logger.Log.Error("err: ", zap.String("", err.Error()))
		}

		return err
	}
}
