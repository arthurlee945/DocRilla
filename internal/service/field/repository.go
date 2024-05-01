package field

import (
	"context"

	"github.com/arthurlee945/Docrilla/internal/model"
	"github.com/jmoiron/sqlx"
)

type Repository interface {
}

type repository struct {
	db *sqlx.DB
}

func Get(ctx context.Context, uuid string) (*model.Field, error) {
	return nil, nil
}
