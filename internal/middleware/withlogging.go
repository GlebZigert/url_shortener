package middleware

import (
	"net/http"
	"strconv"
	"time"

	"github.com/GlebZigert/url_shortener.git/internal/logger"
	"go.uber.org/zap"
)

type (
	// берём структуру для хранения сведений об ответе
	responseData struct {
		status int
		size   int
	}

	// добавляем реализацию http.ResponseWriter
	loggingResponseWriter struct {
		http.ResponseWriter // встраиваем оригинальный http.ResponseWriter
		responseData        *responseData
	}
)

func (r *loggingResponseWriter) Write(b []byte) (int, error) {
	// записываем ответ, используя оригинальный http.ResponseWriter
	size, err := r.ResponseWriter.Write(b)
	r.responseData.size += size // захватываем размер
	return size, err
}

func (r *loggingResponseWriter) WriteHeader(statusCode int) {
	// записываем код статуса, используя оригинальный http.ResponseWriter
	r.ResponseWriter.WriteHeader(statusCode)
	r.responseData.status = statusCode // захватываем код статуса
}

// RequestLogger — middleware-логер для входящих HTTP-запросов.
func RequestLogger(h http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		t1 := time.Now()

		responseData := &responseData{
			status: 0,
			size:   0,
		}

		lw := loggingResponseWriter{
			ResponseWriter: w, // встраиваем оригинальный http.ResponseWriter
			responseData:   responseData,
		}

		h(&lw, r)

		logger.Log.Info("got incoming HTTP request",
			zap.String("method", r.Method),
			zap.String("path", r.URL.Path),
			zap.String("dt", time.Since(t1).String()),
			zap.String("size", strconv.Itoa(responseData.size)),
			zap.String("status", strconv.Itoa(responseData.status)),
		)

	})
}