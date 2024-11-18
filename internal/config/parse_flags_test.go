package config

import (
	"bufio"
	"bytes"
	"os"
	"reflect"
	"strings"
	"testing"

	"github.com/GlebZigert/url_shortener.git/internal/filereader"
	"github.com/GlebZigert/url_shortener.git/mocks"
	"github.com/golang/mock/gomock"
	"gotest.tools/assert"
)

func TestParseFlagsCorrect(t *testing.T) {

	ctrl := gomock.NewController(t)
	mockreader := mocks.NewMockGetFileReader(ctrl)
	mockreader.EXPECT().GetReader(gomock.Any()).DoAndReturn(func(path string) (*bufio.Reader, error) {
		data := []byte(`{"server_address": "localhost:8083","base_url": "http://localhost","file_storage_path": "/path/to/file.db","database_dsn": "ddd","enable_https": true}
		`)
		buffer := bytes.NewBuffer(data)
		reader := bufio.NewReader(buffer)
		return reader, nil
	})

	var tests = []struct {
		args         []string
		envVars      map[string]string
		wantedConfig Config
		reader       GetFileReader
	}{

		{[]string{"-a", "localhost:8888",
			"-b", "http://localhost:8888",
			"-l", "debug",
			"-f", "/path/to/file",
			"-d", "dsn",
			"-SECRETKEY", "SECRETKEY",
			"-TOKENEXP", "8",
			"-NumWorkers", "5",
			"-s"},
			map[string]string{},
			Config{
				Values{
					RunAddr:         "localhost:8888",
					BaseURL:         "http://localhost:8888",
					FlagLogLevel:    "debug",
					FileStoragePath: "/path/to/file",
					DatabaseDSN:     "dsn",
					SECRETKEY:       "SECRETKEY",
					TOKENEXP:        8,
					NumWorkers:      5,
					ENABLEHTTPS:     true,
				},
				"",
			},
			filereader.New(),
		},
		{[]string{"-a", "localhost:8888"},
			map[string]string{"RUN_ADDR": "localhost:8889",
				"BASE_URL":          "http://localhost:8081",
				"LOG_LEVEL":         "debug",
				"FILE_STORAGE_PATH": "FILE_STORAGE_PATH",
			},
			Config{
				Values{
					RunAddr:         "localhost:8889",
					BaseURL:         "http://localhost:8081",
					FlagLogLevel:    "debug",
					FileStoragePath: "FILE_STORAGE_PATH",
					DatabaseDSN:     "",
					SECRETKEY:       "supersecretkey",
					TOKENEXP:        3,
					NumWorkers:      3,
					ENABLEHTTPS:     false,
				},
				""},
			filereader.New(),
		},
		//При передаче флага -s или переменной окружения ENABLE_HTTPS запускайте сервер с помощью метода http.ListenAndServeTLS или tls.Listen.

		{[]string{"-s"},
			map[string]string{},
			Config{
				Values{
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
				""},
			filereader.New(),
		},

		{[]string{},
			map[string]string{"ENABLE_HTTPS": "true"},
			Config{
				Values{
					RunAddr:         "localhost:8080",
					BaseURL:         "http://localhost:8080",
					FlagLogLevel:    "info",
					FileStoragePath: "",
					DatabaseDSN:     "",
					SECRETKEY:       "supersecretkey",
					TOKENEXP:        3,
					NumWorkers:      3,
					ENABLEHTTPS:     true,
				}, ""},
			filereader.New(),
		},

		{[]string{"-c", "/some_path"},
			map[string]string{"ENABLE_HTTPS": "true"},
			Config{
				Values{
					RunAddr:         "localhost:8083",
					BaseURL:         "http://localhost",
					FlagLogLevel:    "info",
					FileStoragePath: "/path/to/file.db",
					DatabaseDSN:     "ddd",
					SECRETKEY:       "supersecretkey",
					TOKENEXP:        3,
					NumWorkers:      3,
					ENABLEHTTPS:     true,
				},
				"/some_path",
			},
			mockreader,
		},

		// ... many more test entries here
	}

	for _, tt := range tests {

		os.Clearenv()
		for k, v := range tt.envVars {
			os.Setenv(k, v)
		}

		t.Run(strings.Join(tt.args, " "), func(t *testing.T) {
			config, err := NewConfig("prog", tt.args, tt.reader)

			if err != nil {
				t.Error("error parse config", err)
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
