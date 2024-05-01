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

func NewRepository(db *sqlx.DB) Repository {
	return &repository{
		db,
	}
}

func (r *repository) Get(ctx context.Context, uuid string) (*model.Field, error) {
	field := &model.Field{}
	if err := r.db.GetContext(ctx, field, `SELECT * FROM field WHERE uuid = $1`, uuid); err != nil {
		return nil, err
	}

	return nil, nil
}
