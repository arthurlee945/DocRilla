package repo

import (
	"context"
	"database/sql"

	"github.com/arthurlee945/Docrilla/internal/model"
)

type UserRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{
		db,
	}
}

func (ur *UserRepository) Get(ctx context.Context, userId string) (*model.User, error) {
	query := `
	SELECT uuid, name 
	FROM users
	WHERE id = $1
	`
	var (
		uuid string
		name string
	)
	err := ur.db.QueryRowContext(ctx, query, userId).Scan(&uuid, &name)
	if err != nil {
		return nil, err
	}

	return &model.User{ID: 0, Name: name}, nil
}
