package project

import (
	"log"
	"net/http"
)

type Handler struct{}

func (h *Handler) Create(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusCreated)
	log.Println("Received request to create a webhook endpoint")
	w.Write([]byte("Project is created"))
}
