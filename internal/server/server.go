package server

import (
	"context"
	"net/http"

	"github.com/arthurlee945/Docrilla/internal/middleware"
	"github.com/arthurlee945/Docrilla/internal/service/project"
)

// https://grafana.com/blog/2024/02/09/how-i-write-http-services-in-go-after-13-years/
func New(ctx context.Context, projectService project.Service) http.Handler {
	middlewareStack := middleware.CreateStack(middleware.Logger)

	router, protectedRouter := http.NewServeMux(), http.NewServeMux()
	router.Handle("/", middleware.Auth(protectedRouter))

	registerRoutes(router, protectedRouter, projectService)

	return middlewareStack(router)
}
