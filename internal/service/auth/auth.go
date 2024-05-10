package auth

import (
	"context"

	"github.com/arthurlee945/Docrilla/internal/errors"
)

type contextKey string

const AuthKey contextKey = "server.user"

func GetUser(ctx context.Context) (uint64, error) {
	if userId, ok := ctx.Value(AuthKey).(uint64); ok {
		return userId, nil
	}
	return 0, errors.ErrUnauthorized
}
