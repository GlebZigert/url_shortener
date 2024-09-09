package server

import (
	"context"
	"net/http"
	"time"

	"github.com/GlebZigert/url_shortener.git/internal/db"
	"github.com/GlebZigert/url_shortener.git/internal/logger"
	"github.com/GlebZigert/url_shortener.git/internal/packerr"
)

func Ping(w http.ResponseWriter, req *http.Request) {
	var err error
	defer packerr.AddErrToReqContext(req, &err)
	logger.Log.Info("Ping")

	ctx, cancel := context.WithTimeout(req.Context(), 1*time.Second)
	defer cancel()
	if err := db.Ping(ctx); err == nil {
		w.WriteHeader(http.StatusOK)

	} else {
		w.WriteHeader(http.StatusInternalServerError)
	}

	w.Write([]byte{})

}
