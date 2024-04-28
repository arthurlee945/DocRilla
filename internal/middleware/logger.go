package middleware

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/arthurlee945/Docrilla/internal/logger"
	"go.uber.org/zap"
)

type wrapperWriter struct {
	http.ResponseWriter
	statusCode int
}

func (ww *wrapperWriter) WriterHeader(statusCode int) {
	ww.ResponseWriter.WriteHeader(statusCode)
	ww.statusCode = statusCode
}

func Logger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		wrapped := &wrapperWriter{w, http.StatusOK}

		next.ServeHTTP(wrapped, r)

		log.Println(wrapped.statusCode, r.Method, r.URL.Path, time.Since(start))
		logger.New().Info(
			fmt.Sprintf("%s %s", r.Method, r.URL.Path),
			zap.Int("statusCode", wrapped.statusCode),
			zap.String("duration", time.Since(start).String()),
		)
	})
}
