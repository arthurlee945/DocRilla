package server

import (
	"net/http"

	"github.com/arthurlee945/Docrilla/internal/service/auth"
	"github.com/arthurlee945/Docrilla/internal/service/project"
)

func registerRoutes(publicRouter *http.ServeMux, protectedRouter *http.ServeMux, authService auth.Service, projectService project.Service) {
	auth.RegisterHandler(publicRouter, protectedRouter, authService)
	project.RegisterHandler(protectedRouter, projectService)
}
