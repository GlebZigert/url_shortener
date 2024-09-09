package server

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/GlebZigert/url_shortener.git/internal/config"
	"github.com/GlebZigert/url_shortener.git/internal/logger"
	"github.com/GlebZigert/url_shortener.git/internal/packerr"
	"github.com/GlebZigert/url_shortener.git/internal/services"
)

func Delete(w http.ResponseWriter, req *http.Request) {
	var err error
	defer packerr.AddErrToReqContext(req, &err)
	logger.Log.Info("Delete")
	var todel []string
	body, err := io.ReadAll(req.Body)
	if err != nil {
		return
	}

	if err = json.Unmarshal(body, &todel); err != nil {

		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	user, ok := req.Context().Value(config.UIDkey).(int)
	if !ok {

		w.WriteHeader(http.StatusUnauthorized)

		w.Write([]byte{})
		return
	}

	go services.Delete(todel, user)

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusAccepted)

	w.Write([]byte{})

}
