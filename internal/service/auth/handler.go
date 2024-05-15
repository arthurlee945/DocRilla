package auth

import (
	"net/http"

	"github.com/arthurlee945/Docrilla/internal/errors"
	"github.com/arthurlee945/Docrilla/internal/util/json"
)

func RegisterHandler(publicRouter *http.ServeMux, protectedRouter *http.ServeMux, service Service) {
	handler := NewHandler(service)
	publicRouter.HandleFunc("POST /login", handler.LogIn)
	publicRouter.HandleFunc("POST /signup", handler.SignUp)
	// U know this aint right
	protectedRouter.HandleFunc("DELETE /delete", handler.Delete)
}

type Handler struct {
	service Service
}

func NewHandler(s Service) *Handler {
	return &Handler{s}
}

func (h *Handler) LogIn(w http.ResponseWriter, r *http.Request) {
	loginReq := new(LogInReq)
	if err := json.Decode(r, loginReq); err != nil {
		errors.ServerError(r.Context(), w, err)
		return
	}
	jwt, err := h.service.LogIn(r.Context(), *loginReq)
	if err != nil {
		errors.ServerError(r.Context(), w, err)
		return
	}
	if err := json.Encode(w, http.StatusOK, struct {
		Token string `json:"token"`
	}{jwt}); err != nil {
		errors.ServerError(r.Context(), w, err)
	}
}

func (h *Handler) SignUp(w http.ResponseWriter, r *http.Request) {
	signupReq := new(SignUpRequest)
	if err := json.Decode(r, signupReq); err != nil {
		errors.ServerError(r.Context(), w, err)
		return
	}
	jwt, err := h.service.SignUp(r.Context(), *signupReq)
	if err != nil {
		errors.ServerError(r.Context(), w, err)
		return
	}
	if err := json.Encode(w, http.StatusOK, struct {
		Token string `json:"token"`
	}{jwt}); err != nil {
		errors.ServerError(r.Context(), w, err)
	}
}
func (h *Handler) Delete(w http.ResponseWriter, r *http.Request) {
	if err := h.service.Delete(r.Context()); err != nil {
		errors.ServerError(r.Context(), w, err)
		return
	}
	if err := json.Encode(w, http.StatusAccepted, struct{}{}); err != nil {
		errors.ServerError(r.Context(), w, err)
	}
}
