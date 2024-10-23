package server

import (
	"errors"
	"io"
	"net/http"

	"github.com/GlebZigert/url_shortener.git/internal/packerr"
	"github.com/GlebZigert/url_shortener.git/internal/services"
)

/*
Эндпоинт с методом POST и путём /.
Сервер принимает в теле запроса строку URL как text/plain

	и возвращает ответ с кодом 201
	и сокращённым URL как text/plain.
*/
func (srv *Server) CreateShortURL(w http.ResponseWriter, req *http.Request) {
	var err error
	defer packerr.AddErrToReqContext(req, &err)

	//logger.Log.Info("CreateShortURL")

	body, err := io.ReadAll(req.Body)
	if err != nil {
		return
	}
	url := string(body)

	srv.logger.Info("Запрос шорта на ориджин: ", map[string]interface{}{
		"url": url,
	})

	res := srv.cfg.GetBaseURL() + "/"
	var short string

	user, ok := srv.mdl.CheckUID(req.Context())
	if !ok {
		http.Error(w, "", http.StatusInternalServerError)
		return
	}

	short, err = srv.service.Short(url, user)

	fl := false
	var conflict *services.ErrConflict409
	if err == nil {
		fl = true
		w.WriteHeader(http.StatusCreated)
	} else {

		if errors.As(err, &conflict) {
			fl = true
			srv.logger.Error("конфликт: ", map[string]interface{}{})
			w.WriteHeader(http.StatusConflict)
		} else {
			srv.logger.Error("", map[string]interface{}{
				"err": err.Error(),
			})
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}

	if fl {

		res += short

	}

	w.Write([]byte(res))

}
