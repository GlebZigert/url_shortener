package server

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"

	"github.com/GlebZigert/url_shortener.git/internal/config"
	"github.com/GlebZigert/url_shortener.git/internal/services"
	"github.com/uptrace/bunrouter"
)

/*
принимать в теле запроса JSON-объект {"url":"<some_url>"}
и возвращать в ответ объект {"result":"<ShortURL>"}.
*/

func CreateShortURLfromJSON(w http.ResponseWriter, req bunrouter.Request) error {

	var msg URLmessage

	var buf bytes.Buffer

	_, err := buf.ReadFrom(req.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err = json.Unmarshal(buf.Bytes(), &msg); err != nil {

		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	url := string(msg.URL)

	res := config.BaseURL + "/"
	var resp []byte

	short, err := services.Short(url, -1)

	fl := false
	var header int
	var conflict *services.ErrConflict409
	if err == nil {
		fl = true
		header = http.StatusCreated
	} else {

		if errors.As(err, &conflict) {
			fl = true
			header = http.StatusConflict
		} else {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}

	if fl {

		res += short

		var answer URLanswer

		answer.Result = res

		resp, err = json.Marshal(answer)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(header)
	w.Write(resp)
	return
}
