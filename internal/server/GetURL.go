package server

import (
	"errors"
	"net/http"
	"strings"

	"github.com/GlebZigert/url_shortener.git/internal/logger"
	"github.com/GlebZigert/url_shortener.git/internal/services"
	"go.uber.org/zap"
)

/*
Эндпоинт с методом GET и путём /{id}, где id — идентификатор сокращённого URL (например, /EwHXdJfB).
В случае успешной обработки запроса сервер возвращает ответ
с кодом 307
и оригинальным URL
в HTTP-заголовке Location.

Пример запроса к серверу:

GET /EwHXdJfB HTTP/1.1
Host: localhost:8080
Content-Type: text/plain

*/

func GetURL(w http.ResponseWriter, req *http.Request) (err error) {
	
	var deleted *services.ErrDeleted

	if req.Method == http.MethodGet {

		str := strings.Replace(req.URL.String(), "/", "", 1)
logger.Log.Info("GetURL ",zap.String("sh",str))
		res, err := services.Origin(str)
		if err == nil {
			w.Header().Add("Location", res)
			w.WriteHeader(http.StatusTemporaryRedirect)

			w.Write([]byte(res))

		} else if errors.As(err, &deleted) {

			logger.Log.Info("запрос удаленного шорта: ", zap.String("", err.Error()))
			//w.Header().Add("Location", res)
			w.WriteHeader(http.StatusGone)

			w.Write([]byte(res))
		} else {

			w.Header().Set("Location", "")
			w.WriteHeader(http.StatusTemporaryRedirect)

			w.Write([]byte{})
		}
	}
	return
}
