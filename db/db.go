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
