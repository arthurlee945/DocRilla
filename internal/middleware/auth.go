package middleware

import (
	"context"
	"net/http"

	"github.com/arthurlee945/Docrilla/internal/config"
	"github.com/arthurlee945/Docrilla/internal/logger"
	"github.com/arthurlee945/Docrilla/internal/service/auth"
	"go.uber.org/zap"
)

func Auth(next http.Handler, cfg *config.Config) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		tokenStr := r.Header.Get("Authorization")
		if tokenStr == "" {
			w.WriteHeader(http.StatusUnauthorized)
			logger.From(r.Context()).Error("Unauthorized User Access")
			return
		}
		tokenStr = tokenStr[len("Bearer "):]
		claim, err := auth.VerifyToken(cfg.JwtSecret, tokenStr)
		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			logger.From(r.Context()).Error("Token Verification Failed", zap.Error(err))
			return
		}
		authReq := r.WithContext(context.WithValue(r.Context(), auth.AuthKey, claim.ID))
		next.ServeHTTP(w, authReq)
	})
}
