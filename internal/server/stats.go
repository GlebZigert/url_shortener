package server

import (
	"encoding/json"
	"net/http"

	"github.com/GlebZigert/url_shortener.git/internal/packerr"
)

func (srv *Server) Stats(w http.ResponseWriter, req *http.Request) {

	var err error
	defer packerr.AddErrToReqContext(req, &err)

	type StatsStruct struct {
		Urls  int `json:"urls"`
		Users int `json:"users"`
	}

	stat := StatsStruct{0, 0}

	bytes, err := json.Marshal(stat)

	if err != nil {

	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(bytes)

}
