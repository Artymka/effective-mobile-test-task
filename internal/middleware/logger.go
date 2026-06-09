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

		// Используем defer для гарантированного логирования даже при панике
		defer func() {
			// Восстанавливаемся от паники
			if rec := recover(); rec != nil {
				// Логируем информацию о панике
				logger.ErrorMsg("panic recovered", fmt.Sprintf(
					"panic in request handler: %v",
					rec,
				))

				// Устанавливаем статус 500 Internal Server Error
				wrapped.statusCode = http.StatusInternalServerError

				// Отправляем ответ клиенту, если это еще не сделано
				if !wrapped.wroteHeader {
					wrapped.WriteHeader(http.StatusInternalServerError)
					lib.WriteError(wrapped, "Internal Server Error", http.StatusInternalServerError)
					// wrapped.Write([]byte("Internal Server Error"))
				}
			}

			// Логируем информацию о запросе (выполнится в любом случае)
			logger.Info("request", fmt.Sprintf(
				"[%s] %s %s - %d (%v)",
				r.Method,
				r.URL.Path,
				r.RemoteAddr,
				wrapped.statusCode,
				time.Since(start),
			))
		}()

		// Обработка запроса
		next.ServeHTTP(wrapped, r)
	})
}

type responseWriter struct {
	http.ResponseWriter
	statusCode  int
	wroteHeader bool // флаг, указывающий, был ли уже записан заголовок
}

func (rw *responseWriter) WriteHeader(code int) {
	if !rw.wroteHeader {
		rw.statusCode = code
		rw.ResponseWriter.WriteHeader(code)
		rw.wroteHeader = true
	}
}

func (rw *responseWriter) Write(b []byte) (int, error) {
	if !rw.wroteHeader {
		rw.WriteHeader(http.StatusOK)
	}
	return rw.ResponseWriter.Write(b)
}
