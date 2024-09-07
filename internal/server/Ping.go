package server

import (
	"context"
	"net/http"
	"time"

	"github.com/GlebZigert/url_shortener.git/internal/db"
	"github.com/uptrace/bunrouter"
)

func Ping(w http.ResponseWriter, req bunrouter.Request) (err error) {

	if req.Method == http.MethodGet {
		ctx, cancel := context.WithTimeout(req.Context(), 1*time.Second)
		defer cancel()
		if err = db.Ping(ctx); err == nil {
			w.WriteHeader(http.StatusOK)

		} else {
			w.WriteHeader(http.StatusInternalServerError)
		}

		w.Write([]byte{})

	}
	return err
}
