package config

import (
	"flag"
	"fmt"
	"os"
)

//type key int

const (
	UIDkey int = iota
	JWTkey int = iota
	NEWkey int = iota
	Errkey int = iota
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

var ptr *Config

func NewConfig(progname string, args []string) *Config {

	if ptr == nil {
		fmt.Println("ptr==nil")
		cfg := Config{}
		cfg.ParseFlags(progname, args)
		ptr = &cfg
	}

	return ptr
}

func (cfg *Config) ParseFlags(progname string, args []string) {
	flags := flag.NewFlagSet(progname, flag.ContinueOnError)
	flags.StringVar(&cfg.RunAddr, "a", "localhost:8080", "address and port to run server")
	flags.StringVar(&cfg.BaseURL, "b", "http://localhost:8080", "base address for short URL")
	flags.StringVar(&cfg.FlagLogLevel, "l", "info", "log level")
	flags.StringVar(&cfg.FileStoragePath, "f", "" /*"./short-url-db.json"*/, "file storage path")
	flags.StringVar(&cfg.DatabaseDSN, "d", "", "database dsn")

	flags.StringVar(&cfg.SECRETKEY, "SECRETKEY", "supersecretkey", "ключ")
	flags.IntVar(&cfg.TOKENEXP, "TOKENEXP", 3, "время жизни токена в часах")
	flags.IntVar(&cfg.NumWorkers, "NumWorkers", 3, "количество воркеров в fanOut")

	flags.Parse(args)
	fmt.Println(cfg.GetRunAddr())
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
