package server

import (
	"bytes"
	"encoding/json"
	"net/http"

	"github.com/GlebZigert/url_shortener.git/internal/config"
	"github.com/GlebZigert/url_shortener.git/internal/services"
)

func Delete(w http.ResponseWriter, req *http.Request) {

	var todel []string
	var buf bytes.Buffer

	_, err := buf.ReadFrom(req.Body)

	if err != nil {

		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := json.Unmarshal(buf.Bytes(), &todel); err != nil {

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
