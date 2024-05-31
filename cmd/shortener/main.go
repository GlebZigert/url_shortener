package main

import (
	"fmt"
	"net/http"
	"strings"
)

/*
Эндпоинт с методом POST и путём /.
Сервер принимает в теле запроса строку URL как text/plain
 и возвращает ответ с кодом 201
 и сокращённым URL как text/plain.

*/
func mainPage(w http.ResponseWriter, req *http.Request) {

	if req.Method == http.MethodPost {

		fmt.Println(".")
		//body += "\r\n"
		body := fmt.Sprintf("URL: %s\r\n", req.URL)
		body += fmt.Sprintf("Method: %s\r\n", req.Method)

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
			body += fmt.Sprintf("-- %s: %v\r\n", k, v)

			w.WriteHeader(http.StatusCreated)
			res := "http://localhost:8080/"
			res += Short(k)

			w.Write([]byte(res))
		}

		w.WriteHeader(http.StatusCreated)
		fmt.Println(body)
		w.Write([]byte{})
	}

	if req.Method == http.MethodGet {
		fmt.Printf("URL: %s\r\n", req.URL)
		str := strings.Replace(req.URL.String(), "/", "", 1)
		fmt.Println(str)
		res, err := Origin(str)
		if err == nil {
			w.WriteHeader(http.StatusTemporaryRedirect)
			w.Write([]byte(res))
		}
		w.WriteHeader(http.StatusTemporaryRedirect)
		w.Write([]byte{})
	}

}

/*
Эндпоинт с методом GET и путём /{id}, где id — идентификатор сокращённого URL (например, /EwHXdJfB). В случае успешной обработки запроса сервер возвращает ответ с кодом 307 и оригинальным URL в HTTP-заголовке Location.
Пример запроса к серверу:

GET /EwHXdJfB HTTP/1.1
Host: localhost:8080
Content-Type: text/plain

*/

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
