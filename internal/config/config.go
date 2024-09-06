package config

import (
	"flag"
	"os"
)

type key int

const (
	UIDkey key = iota
	JWTkey key = iota
	NEWkey key = iota
	ErrKey key = iota
	// ...
)

var (
	RunAddr string

	BaseURL string

	FlagLogLevel string

	FileStoragePath string

	DatabaseDSN string

	TOKENEXP int

	SECRETKEY string
)

const Conflict409 = "попытка сократить уже имеющийся в базе URL"

func ParseFlags() {
	flag.StringVar(&RunAddr, "a", "localhost:8080", "address and port to run server")
	flag.StringVar(&BaseURL, "b", "http://localhost:8080", "base address for short URL")
	flag.StringVar(&FlagLogLevel, "l", "info", "log level")
	flag.StringVar(&FileStoragePath, "f", "" /*"./short-url-db.json"*/, "file storage path")
	flag.StringVar(&DatabaseDSN, "d", "", "database dsn")

	flag.StringVar(&SECRETKEY, "SECRETKEY", "supersecretkey", "ключ")
	flag.IntVar(&TOKENEXP, "TOKENEXP", 3, "время жизни токена в часах")

	flag.Parse()

	if envRunAddr := os.Getenv("RUN_ADDR"); envRunAddr != "" {
		RunAddr = envRunAddr
	}

	if envBaseURL := os.Getenv("BASE_URL"); envBaseURL != "" {
		BaseURL = envBaseURL
	}

	if envLogLevel := os.Getenv("LOG_LEVEL"); envLogLevel != "" {
		FlagLogLevel = envLogLevel
	}

	if envFileStoragePath := os.Getenv("FILE_STORAGE_PATH"); envFileStoragePath != "" {
		FileStoragePath = envFileStoragePath
	}
}
