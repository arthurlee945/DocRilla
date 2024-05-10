package auth

import (
	"fmt"
	"time"

	"github.com/arthurlee945/Docrilla/internal/errors"
	"github.com/golang-jwt/jwt/v5"
)

type Claims struct {
	UserID uint64 `json:"userId"`
	jwt.RegisteredClaims
}

type VerifiedClaim struct {
	ID uint64
}

func VerifyToken(secret string, authToken string) (*VerifiedClaim, error) {
	token, err := jwt.Parse(authToken, func(token *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})
	if err != nil {
		return nil, err
	}
	claim, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return nil, errors.ErrInvalidToken
	}
	fmt.Println(claim, claim["userId"])
	userId, ok := claim["userId"].(float64)
	if !ok {
		return nil, errors.ErrInvalidToken
	}
	if expAt, ok := claim["exp"].(float64); !ok && time.Unix(int64(expAt), 0).Before(time.Now()) {
		return nil, errors.ErrInvalidToken
	}
	return &VerifiedClaim{
		ID: uint64(userId),
	}, nil
}

func generateToken(secret string, id uint64) (string, error) {
	tokenString, err := jwt.NewWithClaims(jwt.SigningMethodHS256,
		&Claims{
			id,
			jwt.RegisteredClaims{
				ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 24)),
				Issuer:    "Docrilla",
			},
		},
	).SignedString([]byte(secret))
	if err != nil {
		return "", errors.AuthError(err)
	}
	return tokenString, nil
}
