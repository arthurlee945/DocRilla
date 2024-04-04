package db

import (
	"os"

	"github.com/jmoiron/sqlx"
)

func InitializeTable(db *sqlx.DB) error {
	qs, err := GetMigrationString()
	if err != nil {
		return err
	}
	if _, err := db.Exec(qs); err != nil {
		return err
	}
	return nil
}

func Seed(db *sqlx.DB) error {
	var userID int
	userInsert := `
	INSERT INTO usr (name, email, password, role)
	VALUES ('admin', 'admin@admin.com', 'qwer1234', 'ADMIN')
	RETURNING id
	`
	uErr := db.QueryRow(userInsert).Scan(&userID)
	if uErr != nil {
		return uErr
	}
	aErr := db.QueryRow(`INSERT INTO account (user_id, type, provider) VALUES ($1, $2, $3)`, userID, "SEED", "SEED").Err()
	if aErr != nil {
		return aErr
	}

	var projID int
	projInsert := `
	INSERT INTO project (user_id, endpoint, title, description, document_url) VALUES ($1, $2, $3, $4, $5) RETURNING id
	`
	pErr := db.QueryRow(projInsert, userID, "018ea1b1-b9ba-79af-81ce-81bae9930afa", "SAMPLE TITLE", "SAMPLE DESCRIPTION", "NO URL").Scan(&projID)
	if pErr != nil {
		return pErr
	}

	fieldInsert := `
	INSERT INTO field (project_id, x1, y1, x2, y2, page, type, value)
	VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
	`
	f1Err := db.QueryRow(fieldInsert, projID, 0, 0, 24, 24, 1, "TEXT", "TEST VALUE 1").Err()
	if f1Err != nil {
		return f1Err
	}
	f2Err := db.QueryRow(fieldInsert, projID, 50, 50, 524, 124, 1, "TEXT", "TEST VALUE 2").Err()
	if f2Err != nil {
		return f2Err
	}
	return nil
}

func DropAllTable(db *sqlx.DB) error {
	if _, err := db.Exec(`
DROP Table IF EXISTS account, field, project, session, usr, verification_token;
DROP Type IF EXISTS user_role, role, project_type, type;
	`); err != nil {
		return err
	}
	return nil
}

func GetMigrationString() (string, error) {
	wd, err := os.Getwd()
	if err != nil {
		return "", err
	}
	dat, err := os.ReadFile(wd + "/migration/migration.sql")
	if err != nil {
		return "", err
	}
	return string(dat), nil
}
