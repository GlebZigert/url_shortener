package config

import (
	"reflect"
	"strings"
	"testing"
)

func TestParseFlagsCorrect(t *testing.T) {
	var tests = []struct {
		args []string
		conf Config
	}{
		{[]string{},
			Config{
				RunAddr:         "localhost:8080",
				BaseURL:         "http://localhost:8080",
				FlagLogLevel:    "info",
				FileStoragePath: "",
				DatabaseDSN:     "",
				SECRETKEY:       "supersecretkey",
				TOKENEXP:        3,
				NumWorkers:      3,
			}},

		// ... many more test entries here
	}

	for _, tt := range tests {
		t.Run(strings.Join(tt.args, " "), func(t *testing.T) {
			config := NewConfig("", tt.args)

			if !reflect.DeepEqual(*config, tt.conf) {
				t.Errorf("conf got %+v, want %+v", *config, tt.conf)
			}
		})
	}
}
