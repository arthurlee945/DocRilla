package db

import (
	"os"

	"github.com/arthurlee945/Docrilla/internal/errors"
	"github.com/arthurlee945/Docrilla/internal/model"
	"github.com/arthurlee945/Docrilla/internal/model/mock"
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
	// USER GEN
	var userID uint64
	uRows, err := db.NamedQuery(`
	INSERT INTO usr (id, name, email, password, role)
	VALUES (:id, :name, :email, :password, :role)
	RETURNING id
	`, mock.User)
	if err != nil {
		return ErrFailedToSeedDB.Wrap(err)
	}
	for uRows.Next() {
		uRows.Scan(&userID)
	}
	if err := uRows.Close(); err != nil {
		return ErrFailedToSeedDB.Wrap(err)
	}

	// Account Gen
	if _, err := db.NamedExec(`INSERT INTO account (user_id, type, provider) VALUES (:user_id, :type, :provider)`, mock.Account); err != nil {
		return ErrFailedToSeedDB.Wrap(err)
	}

	// Project Gen
	var projID int
	pRows, err := db.NamedQuery(`
	INSERT INTO project (id, user_id, uuid, route, token, title, description, document_url) VALUES (:id, :user_id, :uuid, :route, :token, :title, :description, :document_url) RETURNING id
	`, mock.Project)
	if err != nil {
		return ErrFailedToSeedDB.Wrap(err)
	}
	for pRows.Next() {
		pRows.Scan(&projID)
	}
	if err := pRows.Close(); err != nil {
		return ErrFailedToSeedDB.Wrap(err)
	}

	// Field Gen
	for _, field := range []*model.Field{&mock.Field1, &mock.Field2} {
		if _, err := db.NamedExec(`
		INSERT INTO field (uuid, project_id, x, y, width, height, page, type)
		VALUES (:uuid, :project_id, :x, :y, :width, :height, :page, :type)
		`, field); err != nil {
			return ErrFailedToSeedDB.Wrap(err)
		}
	}
	return nil
}

func DropAllTable(db *sqlx.DB) error {
	if _, err := db.Exec(`
DROP Table IF EXISTS account, field, project, endpoint, session, usr, verification_token, submission, submitted_field;
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
