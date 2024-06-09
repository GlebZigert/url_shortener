package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"

	"github.com/go-chi/chi/v5"
)

/*
Эндпоинт с методом POST и путём /.
Сервер принимает в теле запроса строку URL как text/plain

	и возвращает ответ с кодом 201
	и сокращённым URL как text/plain.
*/
func Post(w http.ResponseWriter, req *http.Request) {

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

		res := BaseURL + RunAddr

		res += Short(url)
		log += fmt.Sprintln(url, " --> ", res)

		w.Write([]byte(res))

	}
}
func Get(w http.ResponseWriter, req *http.Request) {
	log := ""
	defer fmt.Println(log)
	log += fmt.Sprintf("URL: %s\r\n", req.URL)
	log += fmt.Sprintf("Method: %s\r\n", req.Method)
	if req.Method == http.MethodGet {

		str := strings.Replace(req.URL.String(), "/", "", 1)
		fmt.Println(str)
		res, err := Origin(str)
		if err == nil {
			w.Header().Add("Location", res)
			w.WriteHeader(http.StatusTemporaryRedirect)

			w.Write([]byte(res))
		}
		log += fmt.Sprintln(res, " <-- ", str)
		w.Header().Set("Location", res)
		w.WriteHeader(http.StatusTemporaryRedirect)

		w.Write([]byte{})
	}

	fmt.Println(log)

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

var RunAddr string

var BaseURL string

func ParseFlags() {
	// регистрируем переменную flagRunAddr
	// как аргумент -a со значением :8080 по умолчанию
	flag.StringVar(&RunAddr, "a", ":8080", "address and port to run server")
	flag.StringVar(&BaseURL, "b", "http://localhost", "base address for short URL")
	// парсим переданные серверу аргументы в зарегистрированные переменные
	flag.Parse()
}

func main() {

	ParseFlags()
	fmt.Println("Running server on", RunAddr, " with BasURL ", BaseURL)
	/*
		Сервер должен быть доступен по адресу http://localhost:8080
	*/

	r := chi.NewRouter()

	// r.Get(`/`, Get)

	r.Post(`/`, Post)
	r.Get(`/*`, Get)

	log.Fatal(http.ListenAndServe(RunAddr, r))
}
