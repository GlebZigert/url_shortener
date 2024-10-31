package server

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/GlebZigert/url_shortener.git/internal/packerr"
)

// удаляет шорт
func (srv *Server) Delete(w http.ResponseWriter, req *http.Request) {
	var err error
	defer packerr.AddErrToReqContext(req, &err)
	//logger.Log.Info("Delete")
	var todel []string
	body, err := io.ReadAll(req.Body)
	if err != nil {
		return
	}

	if err = json.Unmarshal(body, &todel); err != nil {

		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	user, ok := srv.mdl.CheckUID(req.Context())
	if !ok {

		w.WriteHeader(http.StatusUnauthorized)

		_, err = w.Write([]byte{})
		return
	}

	go func() {
		err := srv.service.Delete(todel, user)
		if err != nil {
			srv.logger.Error("Delete ", map[string]interface{}{
				"err": err.Error(),
			})

		}
	}()

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusAccepted)

	_, err = w.Write([]byte{})

}
