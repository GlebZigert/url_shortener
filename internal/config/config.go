// Package config is a package for config
package config

import (
	"flag"
	"os"
)

// ключи в контекст
const (
	UIDkey int = iota
	NEWkey int = iota
	Errkey int = iota
	// ...
)

// Config struct is a struct for config
type Config struct {
	RunAddr string

	BaseURL string

	FlagLogLevel string

	FileStoragePath string

	NumWorkers int

	DatabaseDSN string

	TOKENEXP int

	SECRETKEY string

	ENABLEHTTPS bool
}

// GetRunAddr to get RunAddr value
func (cfg *Config) GetRunAddr() string {
	return cfg.RunAddr
}

// GetBaseURL to get BaseURL value
func (cfg *Config) GetBaseURL() string {
	return cfg.BaseURL
}

// GetFlagLogLevel to get log level value
func (cfg *Config) GetFlagLogLevel() string {
	return cfg.FlagLogLevel
}

// GetFileStoragePath to get file storage path value
func (cfg *Config) GetFileStoragePath() string {
	return cfg.FileStoragePath
}

// GetNumWorkers to get vorkers number value
func (cfg *Config) GetNumWorkers() int {
	return cfg.NumWorkers
}

// GetDatabaseDSN to get database dsn value
func (cfg *Config) GetDatabaseDSN() string {
	return cfg.DatabaseDSN
}

// GetTOKENEXP to get tocken exp value
func (cfg *Config) GetTOKENEXP() int {
	return cfg.TOKENEXP
}

// GetSECRETKEY to get sekret key value
func (cfg *Config) GetSECRETKEY() string {
	return cfg.SECRETKEY
}

var ptr *Config

// NewConfig is constructor for Config
func NewConfig(progname string, args []string) (*Config, error) {

	//if ptr == nil {
	cfg := Config{}
	err := cfg.ParseFlags(progname, args)
	if err != nil {
		return nil, err
	}

	ptr = &cfg
	//}

	return ptr, nil
}

// ParseFlags to parse Config fields form flags and envs
func (cfg *Config) ParseFlags(progname string, args []string) (err error) {
	flags := flag.NewFlagSet(progname, flag.ContinueOnError)
	flags.StringVar(&cfg.RunAddr, "a", "localhost:8080", "address and port to run server")
	flags.StringVar(&cfg.BaseURL, "b", "http://localhost:8080", "base address for short URL")
	flags.StringVar(&cfg.FlagLogLevel, "l", "info", "log level")
	flags.StringVar(&cfg.FileStoragePath, "f", "" /*"./short-url-db.json"*/, "file storage path")
	flags.StringVar(&cfg.DatabaseDSN, "d", "", "database dsn")

	flags.StringVar(&cfg.SECRETKEY, "SECRETKEY", "supersecretkey", "ключ")
	flags.IntVar(&cfg.TOKENEXP, "TOKENEXP", 3, "время жизни токена в часах")
	flags.IntVar(&cfg.NumWorkers, "NumWorkers", 3, "количество воркеров в fanOut")
	flags.BoolVar(&cfg.ENABLEHTTPS, "s", false, "enable https")

	err = flags.Parse(args)
	if err != nil {
		return
	}

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

	if envEnableHttps := os.Getenv("ENABLE_HTTPS"); envEnableHttps == "true" {

		cfg.ENABLEHTTPS = true
	}

	return
}
