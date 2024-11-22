package logger

import (
	"context"
	"io"
	"testing"
)

func TestLogrust(t *testing.T) {
	type want struct {
		code        int
		contentType string
		sendsGzip   bool
	}

	type request struct {
		method          string
		url             string
		user            int
		body            io.Reader
		acceptEncoding  bool
		contentEncoding bool
		auth            bool
		jwt             string
	}

	ctx := context.Background()

	tests := []struct {
		name   string
		logger Logger
	}{
		{
			name:   "logrus",
			logger: NewLogrusLogger("debug", ctx),
		},

		{
			name:   "zap",
			logger: NewZapLogger("debug", ctx),
		},
	}

	//store := storager.New(cfg, filereader.New())

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {

			test.logger.Info("Тест инфо: ", map[string]interface{}{})
			test.logger.Debug("Тест инфо: ", map[string]interface{}{})
			test.logger.Warn("Тест инфо: ", map[string]interface{}{})
			test.logger.Error("Тест инфо: ", map[string]interface{}{})
			//	test.logger.Fatal("Тест инфо: ", map[string]interface{}{})
		})
	}
}
