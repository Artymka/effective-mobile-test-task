package middleware

import (
	"fmt"
	"net/http"
	"time"

	"github.com/Artymka/effective-mobile-test-task/internal/lib"
)

func LoggerMiddleware(next http.Handler, logger *lib.Logger) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		// Оборачиваем ResponseWriter для захвата статуса ответа
		wrapped := &responseWriter{
			ResponseWriter: w,
			statusCode:     http.StatusOK,
		}

		// Обработка запроса
		next.ServeHTTP(wrapped, r)

		// Логируем информацию о запросе
		logger.Info("request", fmt.Sprintf(
			"[%s] %s %s - %d (%v)",
			r.Method,
			r.URL.Path,
			r.RemoteAddr,
			wrapped.statusCode,
			time.Since(start),
		))
	})
}

type responseWriter struct {
	http.ResponseWriter
	statusCode int
}

func (rw *responseWriter) WriteHeader(code int) {
	rw.statusCode = code
	rw.ResponseWriter.WriteHeader(code)
}
