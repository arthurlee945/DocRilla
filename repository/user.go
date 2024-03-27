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

func (ur *UserRepository) InitializeTable() error {
	query := `
	CREATE TABLE IF NOT EXISTS user (
		id SERIAL PRIMARY KEY,
		uuid BINARY(16) NOT NULL UNIQUE,
		name VARCHAR(100) NOT NULL,
		created_at TIMESTAMP DEFAULT NOW()
		updated_at DATETIME DEFAULT NOW() ON UPDATE NOW()
	)`

	if _, err := ur.db.Exec(query); err != nil {
		return err
	}
	return nil
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

	return &model.User{ID: uuid, Name: name}, nil
}
