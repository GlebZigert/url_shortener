package main

import (
	"net/http"
)

func main() {

	/*add server to fix
	Не удалось дождаться пока порт 8080 станет доступен для запроса: context deadline exceeded
	*/
	err := http.ListenAndServe(`:8080`, nil)
	if err != nil {
		panic(err)
	}
}
