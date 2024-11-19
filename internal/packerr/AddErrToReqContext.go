package packerr

import (
	"context"
	"errors"
	"io"
	"net/http"
)

type key int

// ключ в контекст
var Errkey key

// Метод для добавления ошибки в контекст реквеста
func AddErrToReqContext(r *http.Request, err *error) *http.Request {
	if err == nil {
		return r
	}

	ctxerr, ok := r.Context().Value(Errkey).(*error)
	if ok {

		*ctxerr = errors.Join(*ctxerr, *err)
	} else {
		ctx := context.WithValue(r.Context(), Errkey, &err)
		r = r.WithContext(ctx)
	}
	return r
}

// Метод для добавления ошибки в контекст реквеста
func AddCloseErrToReqContext(r *http.Request, closer io.Closer) {
	err := closer.Close()
	if err != nil {
		AddErrToReqContext(r, &err)
	}

}

// Метод для чтения ошибки из контекста реквеста
func AddCloseErrToErr(err *error, closer io.Closer) {
	cerr := closer.Close()
	if cerr != nil {
		*err = errors.Join(*err, cerr)
	}

}
