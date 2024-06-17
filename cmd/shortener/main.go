package main

import (
	"fmt"

	"github.com/GlebZigert/url_shortener.git/internal/app"
	"github.com/GlebZigert/url_shortener.git/internal/config"
)

func main() {

	config.ParseFlags()
	fmt.Println("Running server on", config.RunAddr, " with BasURL ", config.BaseURL)
	/*
		Сервер должен быть доступен по адресу http://localhost:8080
	*/

	app.Run()
}
