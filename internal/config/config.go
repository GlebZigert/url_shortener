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
	Errkey key = iota
	// ...
)

type Config struct {
	RunAddr string

	BaseURL string

	FlagLogLevel string

	FileStoragePath string

	NumWorkers int

	DatabaseDSN string

	TOKENEXP int

	SECRETKEY string
}

func NewConfig() *Config {
	cfg := Config{}
	cfg.ParseFlags()
	return &cfg
}

const Conflict409 = "попытка сократить уже имеющийся в базе URL"

func (c *Config) ParseFlags() {
	flag.StringVar(&c.RunAddr, "a", "localhost:8080", "address and port to run server")
	flag.StringVar(&c.BaseURL, "b", "http://localhost:8080", "base address for short URL")
	flag.StringVar(&c.FlagLogLevel, "l", "info", "log level")
	flag.StringVar(&c.FileStoragePath, "f", "" /*"./short-url-db.json"*/, "file storage path")
	flag.StringVar(&c.DatabaseDSN, "d", "", "database dsn")

	flag.StringVar(&c.SECRETKEY, "SECRETKEY", "supersecretkey", "ключ")
	flag.IntVar(&c.TOKENEXP, "TOKENEXP", 3, "время жизни токена в часах")
	flag.IntVar(&c.NumWorkers, "NumWorkers", 3, "количество воркеров в fanOut")

	flag.Parse()

	if envRunAddr := os.Getenv("RUN_ADDR"); envRunAddr != "" {
		c.RunAddr = envRunAddr
	}

	if envBaseURL := os.Getenv("BASE_URL"); envBaseURL != "" {
		c.BaseURL = envBaseURL
	}

	if envLogLevel := os.Getenv("LOG_LEVEL"); envLogLevel != "" {
		c.FlagLogLevel = envLogLevel
	}

	if envFileStoragePath := os.Getenv("FILE_STORAGE_PATH"); envFileStoragePath != "" {
		c.FileStoragePath = envFileStoragePath
	}
}
