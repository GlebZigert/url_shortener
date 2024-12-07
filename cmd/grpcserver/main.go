package main

import (
	// ...

	"errors"
	"log"
	"net/http"

	"github.com/GlebZigert/url_shortener.git/internal/grpcapp"
)

var (
	buildVersion string
	buildDate    string
	buildCommit  string
)

// func check return N/A if value is emptyt
func check(value string) string {
	if value == "" {
		return "N/A"
	}
	return value
}

func main() {

	log.Println("Build version:", check(buildVersion))
	log.Println("Build date:", check(buildDate))
	log.Println("Build commit:", check(buildCommit))

	err := grpcapp.Run()

	if err != nil && !errors.Is(err, http.ErrServerClosed) {
		log.Fatalf("HTTP server error: %v", err)
	}

}
