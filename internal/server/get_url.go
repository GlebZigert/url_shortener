package server

import (
	"errors"
	"net/http"
	"strings"

	"github.com/GlebZigert/url_shortener.git/internal/packerr"
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

На любой некорректный запрос сервер должен возвращать ответ с кодом 400.
*/

func (srv *Server) GetURL(w http.ResponseWriter, req *http.Request) {
	var err error
	defer packerr.AddErrToReqContext(req, &err)

	var deleted *packerr.ErrDeleted

	src := req.URL.String()

	str := strings.Replace(src, "/", "", 1)

	srv.logger.Info("GetURL: ", map[string]interface{}{
		"src":   src,
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
