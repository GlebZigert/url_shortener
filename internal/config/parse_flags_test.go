package config

import (
	"os"
	"reflect"
	"strings"
	"testing"

	"gotest.tools/assert"
)

func TestParseFlagsCorrect(t *testing.T) {
	var tests = []struct {
		args         []string
		envVars      map[string]string
		wantedConfig Config
	}{

		{[]string{"-a", "localhost:8888"},
			map[string]string{},
			Config{
				RunAddr:         "localhost:8888",
				BaseURL:         "http://localhost:8080",
				FlagLogLevel:    "info",
				FileStoragePath: "",
				DatabaseDSN:     "",
				SECRETKEY:       "supersecretkey",
				TOKENEXP:        3,
				NumWorkers:      3,
				ENABLEHTTPS:     false,
			}},
		{[]string{"-a", "localhost:8888"},
			map[string]string{"RUN_ADDR": "localhost:8889",
				"BASE_URL":          "http://localhost:8081",
				"LOG_LEVEL":         "debug",
				"FILE_STORAGE_PATH": "FILE_STORAGE_PATH",
			},
			Config{
				RunAddr:         "localhost:8889",
				BaseURL:         "http://localhost:8081",
				FlagLogLevel:    "debug",
				FileStoragePath: "FILE_STORAGE_PATH",
				DatabaseDSN:     "",
				SECRETKEY:       "supersecretkey",
				TOKENEXP:        3,
				NumWorkers:      3,
				ENABLEHTTPS:     false,
			}},
		//При передаче флага -s или переменной окружения ENABLE_HTTPS запускайте сервер с помощью метода http.ListenAndServeTLS или tls.Listen.

		{[]string{"-s"},
			map[string]string{},
			Config{
				RunAddr:         "localhost:8080",
				BaseURL:         "http://localhost:8080",
				FlagLogLevel:    "info",
				FileStoragePath: "",
				DatabaseDSN:     "",
				SECRETKEY:       "supersecretkey",
				TOKENEXP:        3,
				NumWorkers:      3,
				ENABLEHTTPS:     true,
			},
		},

		{[]string{},
			map[string]string{"ENABLE_HTTPS": "true"},
			Config{
				RunAddr:         "localhost:8080",
				BaseURL:         "http://localhost:8080",
				FlagLogLevel:    "info",
				FileStoragePath: "",
				DatabaseDSN:     "",
				SECRETKEY:       "supersecretkey",
				TOKENEXP:        3,
				NumWorkers:      3,
				ENABLEHTTPS:     true,
			},
		},

		// ... many more test entries here
	}

	for _, tt := range tests {

		os.Clearenv()
		for k, v := range tt.envVars {
			os.Setenv(k, v)
		}

		t.Run(strings.Join(tt.args, " "), func(t *testing.T) {
			config, err := NewConfig("prog", tt.args)

			if err != nil {
				t.Errorf("error parse config")
			}

			if !reflect.DeepEqual(*config, tt.wantedConfig) {
				t.Errorf("conf got %+v, want %+v", *config, tt.wantedConfig)
			}
			assert.Equal(t, tt.wantedConfig.RunAddr, config.GetRunAddr())
			assert.Equal(t, tt.wantedConfig.BaseURL, config.GetBaseURL())
			assert.Equal(t, tt.wantedConfig.FlagLogLevel, config.GetFlagLogLevel())
			assert.Equal(t, tt.wantedConfig.FileStoragePath, config.GetFileStoragePath())
			assert.Equal(t, tt.wantedConfig.NumWorkers, config.GetNumWorkers())
			assert.Equal(t, tt.wantedConfig.DatabaseDSN, config.GetDatabaseDSN())
			assert.Equal(t, tt.wantedConfig.TOKENEXP, config.GetTOKENEXP())
			assert.Equal(t, tt.wantedConfig.SECRETKEY, config.GetSECRETKEY())
			assert.Equal(t, tt.wantedConfig.ENABLEHTTPS, config.GetENABLEHTTPSflag())
		})
	}
}
