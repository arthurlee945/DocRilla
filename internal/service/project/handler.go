package project

import (
	"net/http"
)

// TODO: NEED Auth Implemented
type Handler struct {
	service Service
}

func NewHandler(s Service) *Handler {
	return &Handler{
		service: s,
	}
}

func (h *Handler) GetAll(w http.ResponseWriter, r *http.Request) {
}

func (h *Handler) GetOverviewById(w http.ResponseWriter, r *http.Request) {

}

func (h *Handler) GetDetailById(w http.ResponseWriter, r *http.Request) {

}

func (h *Handler) Create(w http.ResponseWriter, r *http.Request) {

}

func (h *Handler) Update(w http.ResponseWriter, r *http.Request) {

}

func (h *Handler) Delete(w http.ResponseWriter, r *http.Request) {

}
