package transport

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/GlebZigert/url_shortener.git/internal/config"
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

		w.WriteHeader(http.StatusCreated)

		res := config.BaseURL + "/"

		res += services.Short(url)
		log += fmt.Sprintln(url, " --> ", res)

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
и возвращать в ответ объект {"result":"<short_url>"}.
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

	res += services.Short(url)

	var answer URLanswer

	answer.Result = res

	resp, err := json.Marshal(answer)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	w.Write(resp)

}
