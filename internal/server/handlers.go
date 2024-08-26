package server

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"strings"
	"time"

	"github.com/GlebZigert/url_shortener.git/internal/config"
	"github.com/GlebZigert/url_shortener.git/internal/db"
	"github.com/GlebZigert/url_shortener.git/internal/logger"
	"github.com/GlebZigert/url_shortener.git/internal/services"
)

/*
Эндпоинт с методом POST и путём /.
Сервер принимает в теле запроса строку URL как text/plain

	и возвращает ответ с кодом 201
	и сокращённым URL как text/plain.
*/
func CreateShortURL(w http.ResponseWriter, req *http.Request) {

	if req.Method == http.MethodPost {

		body, err := io.ReadAll(req.Body)
		if err != nil {
			return
		}
		url := string(body)

		res := config.BaseURL + "/"

		short, err := services.Short(url)

		fl := false
		var conflict *services.ErrConflict409
		if err == nil {
			fl = true
			w.WriteHeader(http.StatusCreated)
		} else {
			logger.Log.Error(err.Error())
			if errors.As(err, &conflict) {
				fl = true
				w.WriteHeader(http.StatusConflict)
			} else {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
		}

		if fl {

			res += short

		}

		w.Write([]byte(res))

	}
}

/*
Эндпоинт с методом GET и путём /{id}, где id — идентификатор сокращённого URL (например, /EwHXdJfB).
В случае успешной обработки запроса сервер возвращает ответ
с кодом 307
и оригинальным URL
в HTTP-заголовке Location.

Пример запроса к серверу:

GET /EwHXdJfB HTTP/1.1
Host: localhost:8080
Content-Type: text/plain

*/

func GetURL(w http.ResponseWriter, req *http.Request) {

	if req.Method == http.MethodGet {

		str := strings.Replace(req.URL.String(), "/", "", 1)

		res, err := services.Origin(str)
		if err == nil {
			w.Header().Add("Location", res)
			w.WriteHeader(http.StatusTemporaryRedirect)

			w.Write([]byte(res))

		} else {
			logger.Log.Error(err.Error())
			w.Header().Set("Location", "")
			w.WriteHeader(http.StatusTemporaryRedirect)

			w.Write([]byte{})
		}
	}

}

/*
принимать в теле запроса JSON-объект {"url":"<some_url>"}
и возвращать в ответ объект {"result":"<ShortURL>"}.
*/

func CreateShortURLfromJSON(w http.ResponseWriter, req *http.Request) {

	var msg URLmessage

	var buf bytes.Buffer

	_, err := buf.ReadFrom(req.Body)
	if err != nil {
		logger.Log.Error(err.Error())
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err = json.Unmarshal(buf.Bytes(), &msg); err != nil {
		logger.Log.Error(err.Error())
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	url := string(msg.URL)

	res := config.BaseURL + "/"
	var resp []byte

	short, err := services.Short(url)

	fl := false
	var header int
	var conflict *services.ErrConflict409
	if err == nil {
		fl = true
		header = http.StatusCreated
	} else {
		logger.Log.Error(err.Error())
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

}

/*

 */

func GetURLs(w http.ResponseWriter, req *http.Request) {

	type URLs struct {
		ShortURL    string `json:"shortURL"`
		OriginalURL string `json:"originalURL"`
	}

	res := []URLs{}
	for a, b := range services.GetAll() {
		res = append(res, URLs{a, b})
	}

	resp, err := json.Marshal(res)
	if err != nil {
		logger.Log.Error(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if req.Method == http.MethodGet {

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		w.Write(resp)

	}

}

func Ping(w http.ResponseWriter, req *http.Request) {

	if req.Method == http.MethodGet {

		ctx, cancel := context.WithTimeout(req.Context(), 1*time.Second)
		defer cancel()
		if err := db.Ping(ctx); err == nil {
			w.WriteHeader(http.StatusOK)
		} else {
			w.WriteHeader(http.StatusInternalServerError)
		}

		w.Write([]byte{})

	}

}

type Batch struct {
	CorrelationID string `json:"correlation_id"`
	OriginalURL   string `json:"original_url"`
}

type BatchBack struct {
	CorrelationID string `json:"correlation_id"`
	ShortURL      string `json:"short_url"`
}

func Batcher(w http.ResponseWriter, req *http.Request) {

	if req.Method == http.MethodPost {

		var batches []Batch

		var buf bytes.Buffer

		_, err := buf.ReadFrom(req.Body)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		if err := json.Unmarshal(buf.Bytes(), &batches); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		ll := len(batches)
		batchback := make([]BatchBack, ll)
		var conflict *services.ErrConflict409
		for i, b := range batches {

			ress, err := services.Short(b.OriginalURL)
			if err == nil || errors.As(err, &conflict) {
				res := config.BaseURL + "/" + ress
				batchback[i] = BatchBack{b.CorrelationID, res}
			}
		}

		resp, err := json.Marshal(batchback)
		if err != nil {

			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)

		w.Write(resp)

	}

}
