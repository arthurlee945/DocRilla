package db

import (
	"os"

	"github.com/arthurlee945/Docrilla/internal/errors"
	"github.com/jmoiron/sqlx"
)

const (
	ErrFailedToGetMigrationString = errors.Error("err_failed_to_get_migration_string: database failed to initialize table")
	ErrDBFailedToInitializeTable  = errors.Error("err_db_failed_to_initialize: database failed to initialize table")
	ErrFailedToSeedDB             = errors.Error("err_failed_to_seed_database: ")
)

func InitializeTable(db *sqlx.DB) error {
	qs, err := getMigrationString()
	if err != nil {
		return ErrFailedToGetMigrationString.Wrap(err)
	}
	if _, err := db.Exec(qs); err != nil {
		return ErrDBFailedToInitializeTable.Wrap(err)
	}
	return nil
}

func Seed(db *sqlx.DB) error {
	tx, err := db.Beginx()
	if err != nil {
		return ErrFailedToSeedDB.Wrap(err)
	}
	var userID int
	userInsert := `
	INSERT INTO usr (name, email, password, role)
	VALUES ('admin', 'admin@admin.com', 'qwer1234', 'ADMIN')
	RETURNING id
	`
	err = tx.QueryRow(userInsert).Scan(&userID)
	if err != nil {
		tx.Rollback()
		return ErrFailedToSeedDB.Wrap(err)
	}
	err = tx.QueryRow(`INSERT INTO account (user_id, type, provider) VALUES ($1, $2, $3)`, userID, "SEED", "SEED").Err()
	if err != nil {
		tx.Rollback()
		return ErrFailedToSeedDB.Wrap(err)
	}

	var projID int
	projInsert := `
	INSERT INTO project (user_id, title, description, document_url) VALUES ($1, $3, $4, $5) RETURNING id
	`
	err = tx.QueryRow(projInsert, userID, "SAMPLE TITLE", "SAMPLE DESCRIPTION", "NO URL").Scan(&projID)
	if err != nil {
		tx.Rollback()
		return err
	}

	fieldInsert := `
	INSERT INTO field (project_id, x1, y1, x2, y2, page, type, value)
	VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
	`
	err = tx.QueryRow(fieldInsert, projID, 0, 0, 24, 24, 1, "TEXT", "TEST VALUE 1").Err()
	if err != nil {
		tx.Rollback()
		return ErrFailedToSeedDB.Wrap(err)
	}
	err = tx.QueryRow(fieldInsert, projID, 50, 50, 524, 124, 1, "TEXT", "TEST VALUE 2").Err()
	if err != nil {
		tx.Rollback()
		return ErrFailedToSeedDB.Wrap(err)
	}
	if err := tx.Commit(); err != nil {
		return ErrFailedToSeedDB.Wrap(err)
	}
	return nil
}

func DropAllTable(db *sqlx.DB) error {
	if _, err := db.Exec(`
DROP Table IF EXISTS account, field, project, session, usr, verification_token, submission, submitted_field;
DROP Type IF EXISTS user_role, role, project_type, type;
	`); err != nil {
		return err
	}
	return nil
}

func getMigrationString() (string, error) {
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
