package app

import (
	"fmt"

	"github.com/GlebZigert/url_shortener.git/internal/config"
	"github.com/GlebZigert/url_shortener.git/internal/transport"
)

func Run() {

	config.ParseFlags()
	fmt.Println("Running server on", config.RunAddr, " with BasURL ", config.BaseURL)
	transport.InitRouter()
}
