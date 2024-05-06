package server

import (
	"net/http"

	"github.com/arthurlee945/Docrilla/internal/service/project"
)

func registerRoutes(router *http.ServeMux, protectedRouter *http.ServeMux, projectService project.Service) {
	project.RegisterHandler(protectedRouter, projectService)
}
