package server

import (
	"errors"
	"net/http"
	"strings"

	"github.com/GlebZigert/url_shortener.git/internal/packerr"
	"github.com/GlebZigert/url_shortener.git/internal/services"
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

func (srv *Server) GetURL(w http.ResponseWriter, req *http.Request) {
	var err error
	defer packerr.AddErrToReqContext(req, &err)

	var deleted *services.ErrDeleted

	str := strings.Replace(req.URL.String(), "/", "", 1)

	srv.logger.Info("GetURL: ", map[string]interface{}{

		"short": str,
	})

	res, err := srv.service.Origin(str)
	if err == nil {
		w.Header().Add("Location", res)
		w.WriteHeader(http.StatusTemporaryRedirect)

		w.Write([]byte(res))
		return

	}
	if errors.As(err, &deleted) {

		srv.logger.Info("запрос удаленного шорта: ", map[string]interface{}{

			"": err.Error(),
		})

		//w.Header().Add("Location", res)
		w.WriteHeader(http.StatusGone)

		w.Write([]byte(res))
		return
	}

	w.Header().Set("Location", "")
	w.WriteHeader(http.StatusTemporaryRedirect)

	w.Write([]byte{})

}
