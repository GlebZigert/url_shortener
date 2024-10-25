package packerr

import (
	"context"
	"errors"
	"net/http"
)

type key int

var Errkey key

// Метод для добавления ошибки в контекст реквеста
func AddErrToReqContext(r *http.Request, err *error) {
	if err == nil {
		return
	}

	ctxerr, ok := r.Context().Value(Errkey).(*error)
	if ok {

		*ctxerr = errors.Join(*ctxerr, *err)
	} else {
		ctx := context.WithValue(r.Context(), Errkey, &err)
		r = r.WithContext(ctx)
	}
}

//Метод для чтения ошибки из контекста реквеста
