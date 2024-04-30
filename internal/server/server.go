package server

import (
	"net/http"

	"github.com/arthurlee945/Docrilla/internal/logger"
)

// https://grafana.com/blog/2024/02/09/how-i-write-http-services-in-go-after-13-years/
func New(logger *logger.Logger) http.Handler {
	mux := http.NewServeMux()

	// ctx := context.Background()
	// publicRouter := http.NewServeMux()
	// protectedRouter := http.NewServeMux()

	return mux
}
