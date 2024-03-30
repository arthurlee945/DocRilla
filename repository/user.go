package repo

import (
	"database/sql"

	"github.com/arthurlee945/Docrilla/model"
)

type UserRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{
		db,
	}
}

func (ur *UserRepository) Get(userId string) (*model.User, error) {
	query := `
	SELECT uuid, name 
	FROM users
	WHERE id = $1
	`
	var (
		uuid string
		name string
	)
	err := ur.db.QueryRow(query, userId).Scan(&uuid, &name)
	if err != nil {
		return nil, err
	}

	return &model.User{UserID: 0, Name: name}, nil
}
