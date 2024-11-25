package middleware

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/GlebZigert/url_shortener.git/mocks"
	"github.com/golang/mock/gomock"
	"gotest.tools/v3/assert"

	convert "github.com/GlebZigert/url_shortener.git/pkg/convertIPtoCIDR"
)

// Тест миддла Stat
func TestFiltrIP(t *testing.T) {

	type answer struct {
		status int
	}

	tests := []struct {
		name   string
		ip     string
		answer answer
	}{
		{
			name:   "запрос который не пройдет",
			ip:     "",
			answer: answer{status: http.StatusForbidden},
		},
		{
			name:   "запрос который также не пройдет",
			ip:     "192.168.2.104",
			answer: answer{status: http.StatusForbidden},
		},
		{
			name:   "запрос который  пройдет",
			ip:     "192.168.1.12",
			answer: answer{status: http.StatusOK},
		},
	}

	ctrl := gomock.NewController(t)
	//мок на источник cidr
	cidr := mocks.NewMockSrcCIDR(ctrl)

	cidr.EXPECT().GetCIDR().DoAndReturn(func() []string {
		cidr, err := convert.IPv4RangeToCIDR("192.168.1.10", "192.168.1.17")
		if err != nil {
			return []string{}
		}
		return cidr
	}).AnyTimes()
	//в конструктор миддлов надо передать аргументом интерфес с методом получения CIDR
	mdl := NewMiddlewares(nil, nil, cidr)

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {

			r := httptest.NewRequest(http.MethodGet, "/api/internal/stats", nil)

			if test.ip != "" {
				r.Header.Set("X-Real-IP", test.ip)
			}

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
