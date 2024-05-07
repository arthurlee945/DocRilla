package middleware

import (
	"context"
	"fmt"
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
		servLogger := logger.New()
		loggedReq := r.WithContext(context.WithValue(r.Context(), logger.LoggerKey, servLogger))
		next.ServeHTTP(wrapped, loggedReq)

		servLogger.Info(
			fmt.Sprintf("%s %s", r.Method, r.URL.Path),
			zap.Int("statusCode", wrapped.statusCode),
			zap.String("duration", time.Since(start).String()),
		)
	})
}
