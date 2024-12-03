package server

import (
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/GlebZigert/url_shortener.git/mocks"
	"github.com/golang/mock/gomock"
	"gotest.tools/v3/assert"
)

/*
Тест ендпойнта Stat
Должен вернуть JSON

	{
	  "urls": <int>, // количество сокращённых URL в сервисе
	  "users": <int> // количество пользователей в сервисе

}
*/
func TestStats(t *testing.T) {

	t.Run("test Stats endpoint", func(t *testing.T) {

		ctrl := gomock.NewController(t)
		service := mocks.NewMockSrvService(ctrl)

		service.EXPECT().GetUsersCount().DoAndReturn(func() int {
			return 3
		}).AnyTimes()

		service.EXPECT().GetUrlsCount().DoAndReturn(func() int {
			return 3
		}).AnyTimes()

		server, err := NewServer(nil, nil, nil, service, nil)
		if err != nil {
			t.Error(err.Error())
		}

		r := httptest.NewRequest(http.MethodGet, "/api/internal/stats", nil)
		w := httptest.NewRecorder()

		server.Stats(w, r)

		res := w.Result()

		assert.Equal(t, res.Header.Get("Content-Type"), "application/json")
		assert.Equal(t, res.StatusCode, http.StatusOK)
		//Здесь должен быть получен ответ JSON

		/*
			{
			  "urls": <int>, // количество сокращённых URL в сервисе
			  "users": <int> // количество пользователей в сервисе
			}

		*/

		type StatsStruct struct {
			Urls  int `json:"urls"`
			Users int `json:"users"`
		}

		var varStat StatsStruct

		body, err := io.ReadAll(res.Body)
		if err != nil {
			t.Error(err.Error())
		}
		closeErr := res.Body.Close()
		if closeErr != nil {
			return
		}

		if err = json.Unmarshal(body, &varStat); err != nil {
			t.Error(err.Error())
		}

	})

}
