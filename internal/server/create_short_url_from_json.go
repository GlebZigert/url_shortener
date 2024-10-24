package server

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"

	"github.com/GlebZigert/url_shortener.git/internal/packerr"
)

/*
принимать в теле запроса JSON-объект {"url":"<some_url>"}
и возвращать в ответ объект {"result":"<ShortURL>"}.
*/

func (srv *Server) CreateShortURLfromJSON(w http.ResponseWriter, req *http.Request) {
	var err error
	defer packerr.AddErrToReqContext(req, &err)
	//logger.Log.Info("CreateShortURLfromJSON")
	var msg URLmessage

	body, err := io.ReadAll(req.Body)
	if err != nil {
		return //err
	}
	if err = json.Unmarshal(body, &msg); err != nil {

		http.Error(w, err.Error(), http.StatusBadRequest)
		return //err
	}

	url := string(msg.URL)

	res := srv.cfg.GetBaseURL() + "/"
	var resp []byte

	short, err := srv.service.Short(url, -1)

	fl := false
	var header int
	var conflict packerr.ErrConflict409
	if err == nil {
		fl = true
		header = http.StatusCreated
	} else {

		if errors.As(err, &conflict) {
			fl = true
			header = http.StatusConflict
		} else {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return //err
		}
	}

	if fl {

		res += short

		var answer URLanswer

		answer.Result = res

		resp, err = json.Marshal(answer)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return //err
		}
	}
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(header)
	w.Write(resp)

}
