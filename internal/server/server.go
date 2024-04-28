package server

import (
	"net/http"
)

// https://grafana.com/blog/2024/02/09/how-i-write-http-services-in-go-after-13-years/
func New() http.Handler {
	mux := http.NewServeMux()
	// ctx := context.Background()
	// publicRouter := http.NewServeMux()
	// protectedRouter := http.NewServeMux()

	return mux
}
