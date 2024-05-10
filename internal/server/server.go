package server

import (
	"context"
	"net/http"

	"github.com/arthurlee945/Docrilla/internal/config"
	"github.com/arthurlee945/Docrilla/internal/middleware"
	"github.com/arthurlee945/Docrilla/internal/service/auth"
	"github.com/arthurlee945/Docrilla/internal/service/project"
)

type Test struct {
	Test string `json:"test"`
}

// https://grafana.com/blog/2024/02/09/how-i-write-http-services-in-go-after-13-years/
func New(ctx context.Context, cfg *config.Config, authService auth.Service, projectService project.Service) http.Handler {
	stack := middleware.CreateStack(middleware.Logger)

	router, protectedRouter := http.NewServeMux(), http.NewServeMux()
	registerRoutes(router, protectedRouter, authService, projectService)

	router.Handle("/", middleware.Auth(protectedRouter, cfg))

	return stack(router)
}
