// Package config is a package for config
package config

import (
	"bufio"
	"encoding/json"
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
	Values
	configFile string
}

// Values struct is a struct for Values
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

type GetFileReader interface {
	GetReader(path string) (*bufio.Reader, error)
}

// NewConfig is constructor for Config
func NewConfig(progname string, args []string, getreader GetFileReader) (*Config, error) {

	//if ptr == nil {
	cfg := Config{}
	err := cfg.ParseFlags(progname, args, getreader)
	if err != nil {
		return nil, err
	}

	ptr = &cfg
	//}

	return ptr, nil
}

// ParseFlags to parse Config fields form flags and envs
func (cfg *Config) ParseFlags(progname string, args []string, getreader GetFileReader) (err error) {

	//дефолтные значения -  низкий приоритет - перетрутся любым енвом и флагом

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

	//флаги и енвы - высокий приоритет - и еще среди них файл конфигурации
	//берем флаги
	var flagEnv Values

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
	flags.StringVar(&cfg.configFile, "c", "", "файл конфигурации")

	err = flags.Parse(args)
	if err != nil {
		return
	}
	var aFlag bool
	var bFlag bool
	var lFlag bool
	var fFlag bool
	var dFlag bool
	var sFlag bool
	var flagSECRETKEY bool
	var flagTOKENEXP bool
	var flagNumWorkers bool

	visitor := func(a *flag.Flag) {

		switch a.Name {
		case "a":
			aFlag = true
		case "b":
			bFlag = true
		case "l":
			lFlag = true
		case "f":
			fFlag = true
		case "d":
			dFlag = true
		case "s":
			sFlag = true

		case "SECRETKEY":
			flagSECRETKEY = true
		case "TOKENEXP":
			flagTOKENEXP = true
		case "NumWorkers":
			flagNumWorkers = true
		}

	}

	flags.Visit(visitor)

	//берем енвы если есть и переписываем ими флаги
	if envRunAddr := os.Getenv("RUN_ADDR"); envRunAddr != "" {
		aFlag = true
		flagEnv.RunAddr = envRunAddr
	}

	if envBaseURL := os.Getenv("BASE_URL"); envBaseURL != "" {
		bFlag = true
		flagEnv.BaseURL = envBaseURL
	}

	if envLogLevel := os.Getenv("LOG_LEVEL"); envLogLevel != "" {
		lFlag = true
		flagEnv.FlagLogLevel = envLogLevel
	}

	if envFileStoragePath := os.Getenv("FILE_STORAGE_PATH"); envFileStoragePath != "" {
		fFlag = true
		flagEnv.FileStoragePath = envFileStoragePath
	}

	if envEnableHttps := os.Getenv("ENABLE_HTTPS"); envEnableHttps == "true" {
		sFlag = true

		flagEnv.ENABLEHTTPS = true
	}
	var errCfgFile error
	reader, errCfgFile := getreader.GetReader(cfg.configFile)
	//если файл откылся
	if errCfgFile == nil {

		var cfgFileStruct ConfigFileStruct
		data, errCfgFile := reader.ReadBytes('\n')
		if errCfgFile == nil {

			errCfgFile = json.Unmarshal(data, &cfgFileStruct)
			//если файл распарсился

			if errCfgFile == nil {

				//записать из файла в дефолтные значения
				defaultValues.RunAddr = cfgFileStruct.ServerAddress
				defaultValues.BaseURL = cfgFileStruct.BaseURL
				defaultValues.FileStoragePath = cfgFileStruct.FileStoragePath
				defaultValues.DatabaseDSN = cfgFileStruct.DatabaseDsn
				defaultValues.ENABLEHTTPS = cfgFileStruct.EnableHTTPS

			}
		}

	}

	//теперь если есть флаги енвы пишем их - если нет пишем дефолтные
	if aFlag {
		cfg.RunAddr = flagEnv.RunAddr
	} else {
		cfg.RunAddr = defaultValues.RunAddr
	}

	if bFlag {
		cfg.BaseURL = flagEnv.BaseURL
	} else {
		cfg.BaseURL = defaultValues.BaseURL
	}

	if lFlag {
		cfg.FlagLogLevel = flagEnv.FlagLogLevel
	} else {
		cfg.FlagLogLevel = defaultValues.FlagLogLevel
	}

	if fFlag {
		cfg.FileStoragePath = flagEnv.FileStoragePath
	} else {
		cfg.FileStoragePath = defaultValues.FileStoragePath
	}

	if dFlag {
		cfg.DatabaseDSN = flagEnv.DatabaseDSN
	} else {
		cfg.DatabaseDSN = defaultValues.DatabaseDSN
	}

	if sFlag {
		cfg.ENABLEHTTPS = flagEnv.ENABLEHTTPS
	} else {
		cfg.ENABLEHTTPS = defaultValues.ENABLEHTTPS
	}

	if flagSECRETKEY {
		cfg.SECRETKEY = flagEnv.SECRETKEY
	} else {
		cfg.SECRETKEY = defaultValues.SECRETKEY
	}

	if flagTOKENEXP {
		cfg.TOKENEXP = flagEnv.TOKENEXP
	} else {
		cfg.TOKENEXP = defaultValues.TOKENEXP
	}

	if flagNumWorkers {
		cfg.NumWorkers = flagEnv.NumWorkers
	} else {
		cfg.NumWorkers = defaultValues.NumWorkers
	}

	//if flagEnv.RunAddr {
	//	cfg.RunAddr = flagEnv.RunAddr
	//}

	return
}
