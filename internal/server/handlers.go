package server

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/GlebZigert/url_shortener.git/internal/config"
	"github.com/GlebZigert/url_shortener.git/internal/db"
	"github.com/GlebZigert/url_shortener.git/internal/services"
)

/*
Эндпоинт с методом POST и путём /.
Сервер принимает в теле запроса строку URL как text/plain

	и возвращает ответ с кодом 201
	и сокращённым URL как text/plain.
*/
func CreateShortURL(w http.ResponseWriter, req *http.Request) {

	log := ""
	defer fmt.Println(log)
	log += fmt.Sprintf("URL: %s\r\n", req.URL)
	log += fmt.Sprintf("Method: %s\r\n", req.Method)
	if req.Method == http.MethodPost {

		body, err := io.ReadAll(req.Body)
		if err != nil {
			log += "\r\n"
			log += fmt.Sprintln("err: ", err.Error())
			return
		}
		url := string(body)

		res := config.BaseURL + "/"

		short, err := services.Short(url)

		fl := false
		if err == nil {
			fl = true
			w.WriteHeader(http.StatusCreated)
		} else if err.Error() == config.Conflict409 {
			fl = true
			w.WriteHeader(http.StatusConflict)
		} else {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
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
	fmt.Println("GetURL")
	log := ""
	defer fmt.Println(log)
	log += fmt.Sprintf("URL: %s\r\n", req.URL)
	log += fmt.Sprintf("Method: %s\r\n", req.Method)
	if req.Method == http.MethodGet {

		str := strings.Replace(req.URL.String(), "/", "", 1)
		fmt.Println(str)
		res, err := services.Origin(str)
		if err == nil {
			w.Header().Add("Location", res)
			w.WriteHeader(http.StatusTemporaryRedirect)

			w.Write([]byte(res))
			log += fmt.Sprintln(res, " <-- ", str)

		} else {
			log += fmt.Sprintln(res, " <-- ", str)
			w.Header().Set("Location", "")
			w.WriteHeader(http.StatusTemporaryRedirect)

			w.Write([]byte{})
		}
	}

	fmt.Println(log)

}

/*
принимать в теле запроса JSON-объект {"url":"<some_url>"}
и возвращать в ответ объект {"result":"<ShortURL>"}.
*/

func CreateShortURLfromJSON(w http.ResponseWriter, req *http.Request) {
	//fmt.Println("CreateShortURLfromJSON")

	var msg URLmessage

	var buf bytes.Buffer
	// читаем тело запроса
	_, err := buf.ReadFrom(req.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	//
	fmt.Println(buf.String())

	if err = json.Unmarshal(buf.Bytes(), &msg); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	fmt.Println(msg.URL)

	url := string(msg.URL)

	res := config.BaseURL + "/"
	var resp []byte

	short, err := services.Short(url)

	fl := false
	var header int
	if err == nil {
		fl = true
		header = http.StatusCreated
	} else if err.Error() == config.Conflict409 {
		fl = true
		header = http.StatusConflict
	} else {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
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

	log := ""
	defer fmt.Println(log)
	log += fmt.Sprintf("URL: %s\r\n", req.URL)
	log += fmt.Sprintf("Method: %s\r\n", req.Method)

	type URLs struct {
		ShortURL    string `json:"shortURL"`
		OriginalURL string `json:"originalURL"`
	}

	res := []URLs{}
	for a, b := range services.GetAll() {
		//fmt.Println(a, " ", b)
		res = append(res, URLs{a, b})
	}

	resp, err := json.Marshal(res)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if req.Method == http.MethodGet {

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		w.Write(resp)

	}

	fmt.Println(log)

}

func Ping(w http.ResponseWriter, req *http.Request) {
	fmt.Println("Ping")
	log := ""
	defer fmt.Println(log)
	log += fmt.Sprintf("URL: %s\r\n", req.URL)
	log += fmt.Sprintf("Method: %s\r\n", req.Method)

	if req.Method == http.MethodGet {

		if err := db.Ping(); err == nil {
			w.WriteHeader(http.StatusOK)
		} else {
			fmt.Println(err)
			w.WriteHeader(http.StatusInternalServerError)
		}

		w.Write([]byte{})

	}

	fmt.Println(log)

}

type Batch struct {
	Correlation_id string `json:"correlation_id"`
	OriginalURL    string `json:"original_url"`
}

type BatchBack struct {
	Correlation_id string `json:"correlation_id"`
	ShortURL       string `json:"short_url"`
}

func Batcher(w http.ResponseWriter, req *http.Request) {
	fmt.Println("Batch")
	log := ""
	defer fmt.Println(log)
	log += fmt.Sprintf("URL: %s\r\n", req.URL)
	log += fmt.Sprintf("Method: %s\r\n", req.Method)

	if req.Method == http.MethodPost {

		var batches []Batch

		var buf bytes.Buffer
		// читаем тело запроса
		_, err := buf.ReadFrom(req.Body)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		//
		fmt.Println(buf.String())

		if err := json.Unmarshal(buf.Bytes(), &batches); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		fmt.Println(batches)

		var batchback []BatchBack

		for _, b := range batches {

			ress, _ := services.Short(b.OriginalURL)
			res := config.BaseURL + "/" + ress
			batchback = append(batchback, BatchBack{b.Correlation_id, res})
		}

		resp, err := json.Marshal(batchback)
		if err != nil {
			fmt.Println("err: ", err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		fmt.Println("resp: ", string(resp))

		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)

		w.Write(resp)

	}

	fmt.Println(log)

}
