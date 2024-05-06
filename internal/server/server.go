package server

import (
	"context"
	"net/http"

	"github.com/arthurlee945/Docrilla/internal/config"
	"github.com/arthurlee945/Docrilla/internal/middleware"
	"github.com/arthurlee945/Docrilla/internal/service/project"
)

// https://grafana.com/blog/2024/02/09/how-i-write-http-services-in-go-after-13-years/
func New(ctx context.Context, config *config.Config, projectService project.Service) http.Handler {
	middlewareStack := middleware.CreateStack(middleware.Logger)

	router := http.NewServeMux()
	protectedRouter := http.NewServeMux()

	registerRoutes(router, protectedRouter, projectService)

	//merge auth router here

	return middlewareStack(router)
}
