package repo

import (
	"database/sql"
	"fmt"

	"github.com/arthurlee945/Docrilla/model"
	"github.com/arthurlee945/Docrilla/model/user"
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
	enumQuery := fmt.Sprintf("CREATE TYPE IF NOT EXITS UserRole AS ENUM ('%v', '%v')", user.ADMIN, user.USER)
	if _, err := ur.db.Exec(enumQuery); err != nil {
		return err
	}

	query := fmt.Sprintf(`
	CREATE TABLE IF NOT EXISTS User (
		id SERIAL PRIMARY KEY,
		uuid UUID DEFAULT gen_random_uuid(),
		name VARCHAR(100) NOT NULL,
		email NVARCHAR(320) NOT NULL,
		email_verified BOOLEAN DEFAULT FALSE,
		email_verification_token VARCHAR(62),
		password VARCHAR(62) NOT NULL,
		password_changed_at TIMESTAMP,
		reset_password_token VARCHAR(62),
		reset_password_expires TIMESTAMP,
		access_token VARCHAR(62),
		token_expires_at TIMESTAMP,
		role  UserRole DEFAULT %v,
		active BOOLEAN DEFAULT TRUE,
		created_at TIMESTAMP DEFAULT NOW(),
		updated_at TIMESTAMP DEFAULT NOW() ON UPDATE NOW()
	)`, user.USER)

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

	return &model.User{UserID: 0, Name: name}, nil
}
