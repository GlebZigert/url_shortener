package packerr

import (
	"context"
	"errors"
	"net/http"
)

type key int

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

//Метод для чтения ошибки из контекста реквеста
