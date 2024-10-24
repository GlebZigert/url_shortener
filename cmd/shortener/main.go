package main

import (
	_ "net/http/pprof"

	"github.com/GlebZigert/url_shortener.git/internal/app"
)

func main() {

	app.Run()
}
