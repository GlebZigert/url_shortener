package main

import (
	"fmt"
	"io"
	"net/http"
)

/*
Эндпоинт с методом POST и путём /.
Сервер принимает в теле запроса строку URL как text/plain
 и возвращает ответ с кодом 201
 и сокращённым URL как text/plain.

*/
func mainPage(w http.ResponseWriter, req *http.Request) {
	w.WriteHeader(http.StatusCreated)
	fmt.Println(".")
	body := fmt.Sprintf("Method: %s\r\n", req.Method)
	body += "Header ===============\r\n"
	for k, v := range req.Header {
		body += fmt.Sprintf("%s: %v\r\n", k, v)
	}
	body += "Query parameters ===============\r\n"
	if err := req.ParseForm(); err != nil {
		w.Write([]byte(err.Error()))
		return
	}
	for k, v := range req.Form {
		body += fmt.Sprintf("%s: %v\r\n", k, v)
	}

	body_, err := io.ReadAll(req.Body)
	if err == nil {
		body += "\r\n"
		body += string(body_)
	}

	fmt.Println(body)
	w.Write([]byte("http://localhost:8080/EwHXdJfB "))
}

func main() {

	/*
		Сервер должен быть доступен по адресу http://localhost:8080
	*/

	mux := http.NewServeMux()
	mux.HandleFunc(`/`, mainPage)

	err := http.ListenAndServe(`:8080`, mux)
	if err != nil {
		panic(err)
	}
}

//url -X POST --data "https://practicum.yandex.ru/ " http://localhost:8080/
