package middleware

import (
	"net/http"
	"sync"

	"github.com/GlebZigert/url_shortener.git/internal/packerr"
)

// Миддл который проверяет,
// что переданный в заголовке запроса
// X-Real-IP IP-адрес клиента входит в доверенную подсеть,
// в противном случае возвращать статус ответа 403 Forbidden.
var CIDR string

func (mdl *Middleware) FiltrIP(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var err error
		defer packerr.AddErrToReqContext(r, &err)

		//считываем CIDR однократно - при первом использовании миддла
		var once sync.Once
		onceBody := func() {
			CIDR = mdl.SrcCIDR.GetCIDR()
		}

		once.Do(onceBody)

		//смотрим есть ли в заголовке запроса поле  X-Real-IP
		ip := r.Header.Get("X-Real-IP")

		//если нет
		if ip == "" {
			//ставим статус о том что сервер понял запрос, но отказывается его авторизовать.
			w.WriteHeader(http.StatusForbidden)
			return //err
		}

		//если ip есть - проверяем его

		h.ServeHTTP(w, r)

	})
}
