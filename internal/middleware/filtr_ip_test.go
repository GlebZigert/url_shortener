package middleware

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/GlebZigert/url_shortener.git/mocks"
	"github.com/golang/mock/gomock"
	"gotest.tools/v3/assert"
)

// Тест миддла Stat
func TestFiltrIP(t *testing.T) {

	type answer struct {
		status int
	}

	tests := []struct {
		name   string
		answer answer
	}{
		{
			name:   "запрос который не пройдет",
			answer: answer{status: http.StatusBadRequest},
		},
		{
			name:   "запрос который пройдет",
			answer: answer{status: http.StatusOK},
		},
	}

	ctrl := gomock.NewController(t)
	//мок на источник cidr
	cidr := mocks.NewMockSrcCIDR(ctrl)
	//в конструктор миддлов надо передать аргументом интерфес с методом получения CIDR
	mdl := NewMiddlewares(nil, nil, cidr)

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {

			r := httptest.NewRequest(http.MethodGet, "/api/internal/stats", nil)
			w := httptest.NewRecorder()

			testHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				w.WriteHeader(http.StatusOK)
				w.Write(nil)
			})

			handler := mdl.FiltrIP(testHandler)
			handler.ServeHTTP(w, r)

			result := w.Result()

			assert.Equal(t, test.answer.status, result.StatusCode)

		})
	}
}
