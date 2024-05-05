package auth

import (
	"context"

	"github.com/arthurlee945/Docrilla/internal/errors"
	"github.com/arthurlee945/Docrilla/internal/model"
)

type contextKey string

const AuthKey contextKey = "server.user"

func GetUser(ctx context.Context) (*model.User, error) {
	if user, ok := ctx.Value(AuthKey).(*model.User); ok {
		return user, nil
	}
	return nil, errors.ErrUnauthorized
}
