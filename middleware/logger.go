package middleware

import (
	"log"
	"net/http"
	"time"
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
	})
}
