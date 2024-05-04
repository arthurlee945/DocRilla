package project

import (
	"net/http"
	"strconv"

	"github.com/arthurlee945/Docrilla/internal/errors"
	"github.com/arthurlee945/Docrilla/internal/model"
	"github.com/arthurlee945/Docrilla/internal/server"
)

func RegisterHandler(router *http.ServeMux, service Service) {
	handler := NewHandler(service)
	router.HandleFunc("GET /projects", handler.GetAll)
	router.HandleFunc("GET /projects/{id}/overview", handler.GetOverviewById)
	router.HandleFunc("GET /projects/{id}/detail", handler.GetDetailById)
	router.HandleFunc("POST /projects", handler.Create)
	router.HandleFunc("PUT /projects", handler.Update)
	router.HandleFunc("DELETE /projects/{id}", handler.Delete)
}

type Handler struct {
	service Service
}

func NewHandler(s Service) *Handler {
	return &Handler{
		service: s,
	}
}

type GetAllResponse struct {
	projects []model.Project
	cursor   string
}

func (h *Handler) GetAll(w http.ResponseWriter, r *http.Request) {
	limit, cursor := r.FormValue("limit"), r.FormValue("cursor")
	var limitToPass uint8
	if intLimit, err := strconv.Atoi(limit); err == nil && intLimit > 0 {
		limitToPass = uint8(intLimit)
	} else {
		limitToPass = 10
	}
	proj, cursor, err := h.service.GetAll(r.Context(), GetAllRequest{limitToPass, cursor})
	if err != nil {
		errors.ServerHandleError(r.Context(), w, err)
		return
	}
	server.Encode(w, http.StatusOK, &GetAllResponse{
		projects: proj,
		cursor:   cursor,
	})
}

func (h *Handler) GetOverviewById(w http.ResponseWriter, r *http.Request) {
	// id := r.PathValue("id")

}

func (h *Handler) GetDetailById(w http.ResponseWriter, r *http.Request) {

}

func (h *Handler) Create(w http.ResponseWriter, r *http.Request) {

}

func (h *Handler) Update(w http.ResponseWriter, r *http.Request) {

}

func (h *Handler) Delete(w http.ResponseWriter, r *http.Request) {

}
