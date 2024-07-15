package main

import (
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/GlebZigert/url_shortener.git/internal/server"
	"github.com/GlebZigert/url_shortener.git/internal/services"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestHandler(t *testing.T) {
	type want struct {
		code int
		//	response    string
		contentType string
	}
	tests := []struct {
		name string
		want want
	}{
		{
			name: "positive test #1",
			want: want{
				code: http.StatusCreated,
				//	response:    `{"status":"ok"}`,
				//	contentType: "text/plain",
			},
		},
	}
	services.Init()
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			request := httptest.NewRequest(http.MethodPost, "/", nil)
			// создаём новый Recorder
			w := httptest.NewRecorder()
			server.CreateShortURL(w, request)

			res := w.Result()
			// проверяем код ответа
			assert.Equal(t, test.want.code, res.StatusCode)
			// получаем и проверяем тело запроса
			defer res.Body.Close()
			_, err := io.ReadAll(res.Body)

			require.NoError(t, err)
			//	assert.JSONEq(t, test.want.response, string(resBody))
			//	assert.Equal(t, test.want.contentType, res.Header.Get("Content-Type"))
		})
	}
}
