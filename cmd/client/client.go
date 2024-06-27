package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

/*
Эндпоинт с методом POST и путём /.
Сервер принимает в теле запроса строку URL как text/plain
и возвращает ответ с кодом 201 и сокращённым URL как text/plain.
*/
func POST(addr, data string) error {

	url := addr + "/"

	fmt.Println("POST ", url, " ", data)

	resp, err := http.Post(url, "text/plain", strings.NewReader(data))

	if err != nil {
		fmt.Println(err.Error())
		return err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		return err
	}
	if len(body) > 512 {
		body = body[:512]
	}
	fmt.Print(string(body))

	return nil
}

/*
Эндпоинт с методом GET и путём /{id},
где id — идентификатор сокращённого URL (например, /EwHXdJfB).
В случае успешной обработки запроса сервер возвращает ответ
с кодом 307 и оригинальным URL в HTTP-заголовке Location.
*/
func GET(addr, data string) error {
	url := addr + "/" + data
	fmt.Println("GET ", url)

	resp, err := http.Get(url)

	if err != nil {
		fmt.Println(err.Error())
		return err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		return err
	}
	if len(body) > 512 {
		body = body[:512]
	}
	fmt.Print(string(body))

	return err
}

func main() {
	args := os.Args

	if len(args) != 3 {
		fmt.Println("добавь метод, строку")
		return
	}

	fmt.Println(args[1], args[2])

	method := args[1]

	addr := "http://localhost:8080"
	data := args[2]

	switch method {

	case "GET":
		GET(addr, data)

	case "POST":
		POST(addr, data)
	default:
		fmt.Println("метод GET или POST")
	}

}
