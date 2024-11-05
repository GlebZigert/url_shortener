package server

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/GlebZigert/url_shortener.git/internal/packerr"
)

// ендпойнт запрос нескольких шортов
func (srv *Server) GetURLs(w http.ResponseWriter, req *http.Request) {
	var err error
	defer packerr.AddErrToReqContext(req, &err)
	//	logger.Log.Info("GetURLs")
	type URLs struct {
		ShortURL    string `json:"short_url"`
		OriginalURL string `json:"original_url"`
	}

	vv, ok := srv.mdl.CheckNewFlag(req.Context())

	if ok && vv {

		srv.logger.Error("CheckNewFlag: ", map[string]interface{}{})

		w.WriteHeader(http.StatusUnauthorized)

		err = errors.Join(err, &packerr.NewUserTryGetsURLs)
		_, werr := w.Write([]byte{})
		if werr != nil {
			err = errors.Join(err, werr)
		}
		return //errors.New("")
	}

	user, ok := srv.mdl.CheckUID(req.Context())

	if !ok {
		srv.logger.Error("Check NO UID: ", map[string]interface{}{})
		w.WriteHeader(http.StatusUnauthorized)

		_, err = w.Write([]byte{})
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

		_, err = w.Write([]byte{})
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

		_, err = w.Write(resp)

	}

}
