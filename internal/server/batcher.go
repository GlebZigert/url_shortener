package server

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"

	"github.com/GlebZigert/url_shortener.git/internal/packerr"
)

// для запроса шорта
type Batch struct {
	CorrelationID string `json:"correlation_id"`
	OriginalURL   string `json:"original_url"`
}

// ответ на запрос шорта
type BatchBack struct {
	CorrelationID string `json:"correlation_id"`
	ShortURL      string `json:"short_url"`
}

// ендпойнт запроса  множества URL для сокращения
func (srv *Server) Batcher(w http.ResponseWriter, req *http.Request) {

	var err error
	defer packerr.AddErrToReqContext(req, &err)

	srv.logger.Info("Batcher: ", map[string]interface{}{})

	//logger.Info("Batcher")

	/*
		if req.Method != http.MethodPost {
			w.WriteHeader(http.StatusBadRequest)
			return // errors.New("req.Method != http.MethodPost")
		}
	*/

	var batches []Batch

	body, err := io.ReadAll(req.Body)
	if err != nil {

		srv.logger.Error("io.ReadAll(req.Body): ", map[string]interface{}{
			"err: ": err.Error(),
		})
		http.Error(w, err.Error(), http.StatusBadRequest)
		return //err
	}

	if err = json.Unmarshal(body, &batches); err != nil {

		srv.logger.Error("json.Unmarshal(body, &batches): ", map[string]interface{}{
			"err: ": err.Error(),
		})

		http.Error(w, err.Error(), http.StatusBadRequest)

		return // err
	}
	ll := len(batches)
	batchback := make([]BatchBack, ll)
	var conflict *packerr.ErrConflict409
	for i, b := range batches {
		var ress string
		ress, err = srv.service.Short(b.OriginalURL, -1)
		if err == nil || errors.As(err, &conflict) {
			res := srv.cfg.GetBaseURL() + "/" + ress
			batchback[i] = BatchBack{b.CorrelationID, res}
		}
	}

	resp, err := json.Marshal(batchback)
	if err != nil {

		http.Error(w, err.Error(), http.StatusInternalServerError)
		return //err
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	_, err = w.Write(resp)

}
