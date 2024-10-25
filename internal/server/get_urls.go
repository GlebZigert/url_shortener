package server

import (
	"encoding/json"
	"net/http"

	"github.com/GlebZigert/url_shortener.git/internal/config"
	"github.com/GlebZigert/url_shortener.git/internal/packerr"
)

func (srv *Server) GetURLs(w http.ResponseWriter, req *http.Request) {
	var err error
	defer packerr.AddErrToReqContext(req, &err)
	//	logger.Log.Info("GetURLs")
	type URLs struct {
		ShortURL    string `json:"short_url"`
		OriginalURL string `json:"original_url"`
	}

	vv, ok := req.Context().Value(config.NEWkey).(bool)

	if ok && vv {

		w.WriteHeader(http.StatusUnauthorized)

		w.Write([]byte{})
		return //errors.New("")
	}

	user, ok := srv.mdl.CheckUID(req.Context())

	if !ok {

		w.WriteHeader(http.StatusUnauthorized)

		w.Write([]byte{})
		return //errors.New("")
	}

	res := []URLs{}
	for _, sh := range *srv.service.GetAll() {
		if sh.UUID == int(user) {
			res = append(res, URLs{srv.cfg.GetBaseURL() + "/" + sh.ShortURL, sh.OriginalURL})
		}
	}

	if len(res) == 0 {

		//	w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNoContent)

		w.Write([]byte{})
		return //errors.New("StatusNoContent")
	}

	resp, err := json.Marshal(res)
	if err != nil {

		http.Error(w, err.Error(), http.StatusInternalServerError)
		return //err
	}

	if req.Method == http.MethodGet {

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		w.Write(resp)

	}

}
