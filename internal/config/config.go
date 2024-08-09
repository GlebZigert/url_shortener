package config

import (
	"flag"
	"os"
)

var (
	RunAddr string

	BaseURL string

	FlagLogLevel string

	FileStoragePath string

	DatabaseDSN string
)

const Conflict409 = "попытка сократить уже имеющийся в базе URL"

func ParseFlags() {
	// регистрируем переменную flagRunAddr
	// как аргумент -a со значением :8080 по умолчанию
	flag.StringVar(&RunAddr, "a", "localhost:8080", "address and port to run server")
	flag.StringVar(&BaseURL, "b", "http://localhost:8080", "base address for short URL")
	flag.StringVar(&FlagLogLevel, "l", "info", "log level")
	flag.StringVar(&FileStoragePath, "f", "" /*"./short-url-db.json"*/, "file storage path")
	flag.StringVar(&DatabaseDSN, "d", "", "database dsn")
	// парсим переданные серверу аргументы в зарегистрированные переменные
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
