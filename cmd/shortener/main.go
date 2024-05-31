package main

import (
	"net/http"
)

/*
Эндпоинт с методом POST и путём /.
Сервер принимает в теле запроса строку URL как text/plain
 и возвращает ответ с кодом 201
 и сокращённым URL как text/plain.

*/
func mainPage(w http.ResponseWriter, req *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte{})
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
