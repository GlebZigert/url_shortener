package config

import (
	"flag"
	"fmt"
	"os"
)

// type key int
// ключи в контекст
const (
	UIDkey int = iota
	JWTkey int = iota
	NEWkey int = iota
	Errkey int = iota
	// ...
)

// конфиг
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

// взять адрес
func (cfg *Config) GetRunAddr() string {
	return cfg.RunAddr
}

// взять префикс
func (cfg *Config) GetBaseURL() string {
	return cfg.BaseURL
}

// взять уровень логирования
func (cfg *Config) GetFlagLogLevel() string {
	return cfg.FlagLogLevel
}

// взять путь к файловому хранилищу
func (cfg *Config) GetFileStoragePath() string {
	return cfg.FileStoragePath
}

// взять количество горутин
func (cfg *Config) GetNumWorkers() int {
	return cfg.NumWorkers
}

// взять настройки базы
func (cfg *Config) GetDatabaseDSN() string {
	return cfg.DatabaseDSN
}

// взять время жизни токена
func (cfg *Config) GetTOKENEXP() int {
	return cfg.TOKENEXP
}

// взять ключ
func (cfg *Config) GetSECRETKEY() string {
	return cfg.SECRETKEY
}

var ptr *Config

// конструктор
func NewConfig(progname string, args []string) *Config {

	if ptr == nil {
		fmt.Println("ptr==nil")
		cfg := Config{}
		cfg.ParseFlags(progname, args)
		ptr = &cfg
	}

	return ptr
}

// парсить
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
