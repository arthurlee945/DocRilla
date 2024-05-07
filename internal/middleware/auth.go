package middleware

import (
	"context"
	"net/http"

	"github.com/arthurlee945/Docrilla/internal/model/mock"
	"github.com/arthurlee945/Docrilla/internal/service/auth"
)

func Auth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authReq := r.WithContext(context.WithValue(r.Context(), auth.AuthKey, &mock.User))
		next.ServeHTTP(w, authReq)
	})
}
