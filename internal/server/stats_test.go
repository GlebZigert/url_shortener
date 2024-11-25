package server

import (
	"testing"
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

	/*
		var CIDR string
		ctrl := gomock.NewController(t)
		mcfg := mocks.NewMockSrvConfig(ctrl)
		//мок на srvConfig

		tests := []struct {
			name    string
			urls    int       //количество сокращённых URL в сервисе
			users   int       //количество пользователей в сервисе
			XRealIP string    //переданный в заголовке запроса IP-адрес клиента
			cfg     SrvConfig //интерфейс настроек у которого должен быть метод получения бесклассовой адресации getCIDR
		}{{
			name:    "запрос который обработается",
			urls:    5,
			users:   2,
			XRealIP: CIDR,
			cfg:     mcfg,
		}}

		for _, test := range tests {
			t.Run(test.name, func(t *testing.T) {
				ctx := context.Background()

				logger := logger.NewLogrusLogger("debug", ctx)
				auth := auth.NewAuth("dd", 5)
				mdl := middleware.NewMiddlewares(auth, logger)

				mdl.Auth()

			})
		}
	*/
}
