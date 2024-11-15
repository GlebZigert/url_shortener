// Package config is a package for config
package config

import (
	"flag"
	"fmt"
	"os"
)

// ключи в контекст
const (
	UIDkey int = iota
	NEWkey int = iota
	Errkey int = iota
	// ...
)

type Config struct {
	Values
	configFile bool
}

// Config struct is a struct for config
type Values struct {
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
func (cfg *Values) GetRunAddr() string {
	return cfg.RunAddr
}

// GetBaseURL to get BaseURL value
func (cfg *Values) GetBaseURL() string {
	return cfg.BaseURL
}

// GetFlagLogLevel to get log level value
func (cfg *Values) GetFlagLogLevel() string {
	return cfg.FlagLogLevel
}

// GetFileStoragePath to get file storage path value
func (cfg *Values) GetFileStoragePath() string {
	return cfg.FileStoragePath
}

// GetNumWorkers to get vorkers number value
func (cfg *Values) GetNumWorkers() int {
	return cfg.NumWorkers
}

// GetDatabaseDSN to get database dsn value
func (cfg *Values) GetDatabaseDSN() string {
	return cfg.DatabaseDSN
}

// GetTOKENEXP to get tocken exp value
func (cfg *Values) GetTOKENEXP() int {
	return cfg.TOKENEXP
}

// GetSECRETKEY to get sekret key value
func (cfg *Values) GetSECRETKEY() string {
	return cfg.SECRETKEY
}

// GetSECRETKEY to get sekret key value
func (cfg *Values) GetENABLEHTTPSflag() bool {
	return cfg.ENABLEHTTPS
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

	/*
		//значение по дефолту
		defaultValues := Values{
			RunAddr:         "localhost:8080",
			BaseURL:         "http://localhost:8080",
			FlagLogLevel:    "info",
			FileStoragePath: "",
			DatabaseDSN:     "",
			SECRETKEY:       "supersecretkey",
			TOKENEXP:        3,
			NumWorkers:      3,
			ENABLEHTTPS:     false,
		}
	*/

	//дефолтные значения -  низкий приоритет - перетрутся любым енвом и флагом
	var flagEnv Values

	//флаги и енвы - высокий приоритет - и еще среди них файл конфигурации
	//берем флаги
	flags := flag.NewFlagSet(progname, flag.ContinueOnError)
	flags.StringVar(&flagEnv.RunAddr, "a", "", "address and port to run server")
	flags.StringVar(&flagEnv.BaseURL, "b", "", "base address for short URL")
	flags.StringVar(&flagEnv.FlagLogLevel, "l", "", "log level")
	flags.StringVar(&flagEnv.FileStoragePath, "f", "" /*"./short-url-db.json"*/, "file storage path")
	flags.StringVar(&flagEnv.DatabaseDSN, "d", "", "database dsn")

	flags.StringVar(&flagEnv.SECRETKEY, "SECRETKEY", "", "ключ")
	flags.IntVar(&flagEnv.TOKENEXP, "TOKENEXP", 0, "время жизни токена в часах")
	flags.IntVar(&flagEnv.NumWorkers, "NumWorkers", 0, "количество воркеров в fanOut")
	flags.BoolVar(&flagEnv.ENABLEHTTPS, "s", false, "enable https")
	flags.BoolVar(&cfg.configFile, "c", false, "файл конфигурации")

	err = flags.Parse(args)
	if err != nil {
		return
	}

	visitor := func(a *flag.Flag) {
		fmt.Println(">", a.Name, "value=", a.Value)
	}

	fmt.Println("Visit()")
	flags.Visit(visitor)
	fmt.Println("VisitAll()")
	flags.VisitAll(visitor)

	//берем енвы если есть и переписываем ими флаги
	if envRunAddr := os.Getenv("RUN_ADDR"); envRunAddr != "" {
		flagEnv.RunAddr = envRunAddr
	}

	if envBaseURL := os.Getenv("BASE_URL"); envBaseURL != "" {
		flagEnv.BaseURL = envBaseURL
	}

	if envLogLevel := os.Getenv("LOG_LEVEL"); envLogLevel != "" {
		flagEnv.FlagLogLevel = envLogLevel
	}

	if envFileStoragePath := os.Getenv("FILE_STORAGE_PATH"); envFileStoragePath != "" {
		flagEnv.FileStoragePath = envFileStoragePath
	}

	if envEnableHttps := os.Getenv("ENABLE_HTTPS"); envEnableHttps == "true" {

		flagEnv.ENABLEHTTPS = true
	}

	if cfg.configFile {
		//если файл откылся

		//если файл распарсился

		//записать из файла в дефолтные значения
	}

	//теперь если есть флаги енвы пишем их - если нет пишем дефолтные

	//if flagEnv.RunAddr {
	//	cfg.RunAddr = flagEnv.RunAddr
	//}

	return
}
