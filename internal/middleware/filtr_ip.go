package middleware

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"strings"
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
			log.Println("CIDR: ", CIDR)
		}

		once.Do(onceBody)

		//смотрим есть ли в заголовке запроса поле  X-Real-IP
		ip, err := checkIP(r)

		if err != nil {
			log.Println(err.Error())
			w.WriteHeader(http.StatusForbidden)
			return //err
		}

		log.Println("ip: ", ip)

		//если ip есть - проверяем его

		h.ServeHTTP(w, r)

	})
}

func checkIP(r *http.Request) (net.IP, error) {

	// смотрим заголовок запроса X-Real-IP
	ipStr := r.Header.Get("X-Real-IP")
	log.Println("ipStr: ", ipStr)
	// парсим ip

	ip := net.ParseIP(ipStr)
	if ip == nil {
		// если заголовок X-Real-IP пуст, пробуем X-Forwarded-For
		// этот заголовок содержит адреса отправителя и промежуточных прокси
		// в виде 203.0.113.195, 70.41.3.18, 150.172.238.178
		ips := r.Header.Get("X-Forwarded-For")
		// разделяем цепочку адресов
		ipStrs := strings.Split(ips, ",")
		// интересует только первый
		ipStr = ipStrs[0]
		// парсим
		ip = net.ParseIP(ipStr)
	}
	if ip == nil {
		return nil, fmt.Errorf("failed parse ip from http header")
	}
	return ip, nil

}
