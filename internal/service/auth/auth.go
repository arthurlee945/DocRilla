package auth

import (
	"context"
	"time"

	"github.com/arthurlee945/Docrilla/internal/errors"
	"github.com/arthurlee945/Docrilla/internal/logger"
	"github.com/golang-jwt/jwt/v5"
	"go.uber.org/zap"
)

type contextKey string

const AuthKey contextKey = "server.user"

func GetUser(ctx context.Context) (uint64, error) {
	if userId, ok := ctx.Value(AuthKey).(uint64); ok {
		return userId, nil
	}
	return 0, errors.ErrUnauthorized
}

type Claims struct {
	ID  uint64    `json:"id"`
	EXP time.Time `json:"exp"`
	jwt.RegisteredClaims
}

func GenerateToken(secret string, id uint64) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodES256,
		Claims{
			ID:  id,
			EXP: time.Now().Add(time.Hour * 24),
		},
	)
	tokenString, err := token.SignedString(secret)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func VerifyToken(secret string, authToken string) error {
	claim := new(Claims)
	token, err := jwt.ParseWithClaims(authToken, claim, func(token *jwt.Token) (interface{}, error) {
		return secret, nil
	})
	if err != nil {
		return err
	}
	if !token.Valid {
		return errors.ErrInvalidToken
	}
	if claim.EXP.Before(time.Now()) {
		return errors.ErrUnauthorized
	}
	return nil
}

func Authenticate(ctx context.Context, email, password string) (jwt string, err error) {
	logger.From(ctx).Info("authentication req", zap.String("email", email))

	return "", nil
}
