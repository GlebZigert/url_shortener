package main

import (
	"github.com/GlebZigert/url_shortener.git/internal/app"

	"go.uber.org/zap"
)

var sugar zap.SugaredLogger

func main() {

	app.Run()
}
