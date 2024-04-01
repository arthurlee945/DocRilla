package db

import (
	"github.com/arthurlee945/Docrilla/model"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func Connect(connStr string) (*sqlx.DB, error) {
	db, err := sqlx.Open("postgres", connStr)

	if err != nil {
		return nil, err
	}

	if err = db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}

func InitializeTable(db *sqlx.DB) error {
	qs, err := model.GetQueryString()
	if err != nil {
		return err
	}
	if _, err := db.Exec(qs); err != nil {
		return err
	}
	return nil
}

func Seed(db *sqlx.DB) error {
	userInsert := `
	INSERT INTO usr (name, email, password, role)
	VALUES(?, ?, ?, ?)
	`
	user, uErr := db.Exec(userInsert, "admin", "admin@admin.com", "qwer1234", "ADMIN")
	if uErr != nil {
		return uErr
	}
	userID, uIdErr := user.LastInsertId()
	if uIdErr != nil {
		return uIdErr
	}
	_, aErr := db.Exec("INSERT INTO account (user_id, type, provider) VALUES(?, ?, ?)", userID, "SEED", "SEED")
	if aErr != nil {
		return aErr
	}
	projInsert := `
	INSERT INTO proj (user_id, title, description, document_url)
	`
	proj, pErr := db.Exec(projInsert, userID, "SAMPLE TITLE", "SAMPLE DESCRIPTION", "NO URL")
	if pErr != nil {
		return pErr
	}
	projID, pIDErr := proj.LastInsertId()
	if pIDErr != nil {
		return pIDErr
	}
	fieldInsert := `
	INSERT INTO field (project_id, x1, y1, x2, y2, page, type, value)
	VALUES(?, ?, ?, ?, ?, ?, ?, ?)
	`
	_, f1Err := db.Exec(fieldInsert, projID, 0, 0, 24, 24, 1, "TEXT", "TEST VALUE 1")
	if f1Err != nil {
		return f1Err
	}
	_, f2Err := db.Exec(fieldInsert, projID, 50, 50, 524, 124, 1, "TEXT", "TEST VALUE 2")
	if f2Err != nil {
		return f2Err
	}
	return nil
}
