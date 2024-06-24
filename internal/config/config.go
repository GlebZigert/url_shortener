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
)

func ParseFlags() {
	// регистрируем переменную flagRunAddr
	// как аргумент -a со значением :8080 по умолчанию
	flag.StringVar(&RunAddr, "a", "localhost:8080", "address and port to run server")
	flag.StringVar(&BaseURL, "b", "http://localhost:8080", "base address for short URL")
	flag.StringVar(&FlagLogLevel, "l", "info", "log level")
	flag.StringVar(&FileStoragePath, "f", "./short-url-db.txt", "file storage path")
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
