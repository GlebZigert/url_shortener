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

func (cfg *Config) GetRunAddr() string {
	return cfg.RunAddr
}

func (cfg *Config) GetBaseURL() string {
	return cfg.BaseURL
}

func (cfg *Config) GetFlagLogLevel() string {
	return cfg.FlagLogLevel
}

func (cfg *Config) GetFileStoragePath() string {
	return cfg.FileStoragePath
}

func (cfg *Config) GetNumWorkers() int {
	return cfg.NumWorkers
}

func (cfg *Config) GetDatabaseDSN() string {
	return cfg.DatabaseDSN
}

func (cfg *Config) GetTOKENEXP() int {
	return cfg.TOKENEXP
}

func (cfg *Config) GetSECRETKEY() string {
	return cfg.SECRETKEY
}

func NewConfig() *Config {
	cfg := Config{}
	cfg.ParseFlags()
	return &cfg
}

func (cfg *Config) ParseFlags() {
	flag.StringVar(&cfg.RunAddr, "a", "localhost:8080", "address and port to run server")
	flag.StringVar(&cfg.BaseURL, "b", "http://localhost:8080", "base address for short URL")
	flag.StringVar(&cfg.FlagLogLevel, "l", "info", "log level")
	flag.StringVar(&cfg.FileStoragePath, "f", "" /*"./short-url-db.json"*/, "file storage path")
	flag.StringVar(&cfg.DatabaseDSN, "d", "", "database dsn")

	flag.StringVar(&cfg.SECRETKEY, "SECRETKEY", "supersecretkey", "ключ")
	flag.IntVar(&cfg.TOKENEXP, "TOKENEXP", 3, "время жизни токена в часах")
	flag.IntVar(&cfg.NumWorkers, "NumWorkers", 3, "количество воркеров в fanOut")

	flag.Parse()

	if envRunAddr := os.Getenv("RUN_ADDR"); envRunAddr != "" {
		cfg.RunAddr = envRunAddr
	}

	if envBaseURL := os.Getenv("BASE_URL"); envBaseURL != "" {
		cfg.BaseURL = envBaseURL
	}

	if envLogLevel := os.Getenv("LOG_LEVEL"); envLogLevel != "" {
		cfg.FlagLogLevel = envLogLevel
	}

	if envFileStoragePath := os.Getenv("FILE_STORAGE_PATH"); envFileStoragePath != "" {
		cfg.FileStoragePath = envFileStoragePath
	}
}
