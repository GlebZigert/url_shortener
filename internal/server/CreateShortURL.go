package server

import (
	"errors"
	"io"
	"net/http"

	"github.com/GlebZigert/url_shortener.git/internal/config"
	"github.com/GlebZigert/url_shortener.git/internal/logger"
	"github.com/GlebZigert/url_shortener.git/internal/services"
	"go.uber.org/zap"
)

/*
Эндпоинт с методом POST и путём /.
Сервер принимает в теле запроса строку URL как text/plain

	и возвращает ответ с кодом 201
	и сокращённым URL как text/plain.
*/
func CreateShortURL(w http.ResponseWriter, req *http.Request) error {

	if req.Method == http.MethodPost {

		body, err := io.ReadAll(req.Body)
		if err != nil {
			return err
		}
		url := string(body)
		logger.Log.Info("auth: ", zap.String("url", url))

		res := config.BaseURL + "/"
		var short string
		user, ok := req.Context().Value(config.UIDkey).(int)
		if ok {

			short, err = services.Short(url, user)

		}

		fl := false
		var conflict *services.ErrConflict409
		if err == nil {
			fl = true
			w.WriteHeader(http.StatusCreated)
		} else {
			logger.Log.Error(err.Error())
			if errors.As(err, &conflict) {
				fl = true
				w.WriteHeader(http.StatusConflict)
			} else {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return err
			}
		}

		if fl {

			res += short

		}

		w.Write([]byte(res))

	}
	return nil
}
