package server

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"

	"github.com/GlebZigert/url_shortener.git/internal/config"
	"github.com/GlebZigert/url_shortener.git/internal/logger"
	"github.com/GlebZigert/url_shortener.git/internal/packerr"
	"github.com/GlebZigert/url_shortener.git/internal/services"
)

type Batch struct {
	CorrelationID string `json:"correlation_id"`
	OriginalURL   string `json:"original_url"`
}

type BatchBack struct {
	CorrelationID string `json:"correlation_id"`
	ShortURL      string `json:"short_url"`
}

func Batcher(w http.ResponseWriter, req *http.Request) {

	var err error
	defer packerr.AddErrToReqContext(req, &err)

	logger.Log.Info("Batcher")
	if req.Method != http.MethodPost {
		w.WriteHeader(http.StatusBadRequest)
		return // errors.New("req.Method != http.MethodPost")
	}

	var batches []Batch

	body, err := io.ReadAll(req.Body)
	if err != nil {
		return //err
	}

	if err := json.Unmarshal(body, &batches); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return // err
	}
	ll := len(batches)
	batchback := make([]BatchBack, ll)
	var conflict *services.ErrConflict409
	for i, b := range batches {

		ress, err := services.Short(b.OriginalURL, -1)
		if err == nil || errors.As(err, &conflict) {
			res := config.BaseURL + "/" + ress
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

	w.Write(resp)

}
