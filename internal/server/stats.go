package server

import (
	"encoding/json"
	"net/http"

	"github.com/GlebZigert/url_shortener.git/internal/packerr"
)

// Статистика
func (srv *Server) Stats(w http.ResponseWriter, req *http.Request) {

	var err error
	defer packerr.AddErrToReqContext(req, &err)

	type StatsStruct struct {
		Urls  int `json:"urls"`
		Users int `json:"users"`
	}

	users := srv.service.GetUsersCount()
	urls := srv.service.GetUrlsCount()

	stat := StatsStruct{urls, users}

	bytes, err := json.Marshal(stat)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(bytes)

}
