package server

import (
	"context"
	"net/http"

	"github.com/arthurlee945/Docrilla/internal/config"
	"github.com/arthurlee945/Docrilla/internal/logger"
	"github.com/arthurlee945/Docrilla/internal/middleware"
)

// https://grafana.com/blog/2024/02/09/how-i-write-http-services-in-go-after-13-years/
func New(ctx context.Context, logger *logger.Logger, config *config.Config) http.Handler {
	middlewareStack := middleware.CreateStack(middleware.Logger)

	router := http.NewServeMux()
	protectedRouter := http.NewServeMux()

	registerRoutes(router, protectedRouter)

	//merge auth router here

	return middlewareStack(router)
}
