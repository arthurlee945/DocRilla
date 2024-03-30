package repo

import (
	"database/sql"
)

type ProjectRepository struct {
	db *sql.DB
}

func NewProjectRepository(db *sql.DB) *ProjectRepository {
	return &ProjectRepository{
		db,
	}
}

func (pr *ProjectRepository) InitializeTable() error {
	query := `
	CREATE TABLE IF NOT EXISTS Project (
		id SERIAL PRIMARY KEY,
		endpoint UUID DEFAULT gen_random_uuid(),
		user_id INT NOT NULL,
		title VARCHAR(128) NOT NULL,
		description NVARCHAR(512),
		document_url varchar(256) NOT NULL,
		archived BOOLEAN DEFAULT FALSE,
		created_at TIMESTAMP DEFAULT NOW(),
		updated_at TIMESTAMP DEFAULT NOW(),
		CONSTRAINT fk_User FOREIGN KEY(user_id) REFERENCES user_account(id) ON DELETE CASCADE
	)`

	if _, err := pr.db.Exec(query); err != nil {
		return err
	}
	return nil
}
