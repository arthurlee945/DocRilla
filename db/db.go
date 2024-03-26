package db

import (
	"database/sql"

	"github.com/arthurlee945/doc-rilla/config"
	_ "github.com/lib/pq"
)

var cfg = config.New()

func NewConnection() (*sql.DB, error) {
	db, err := sql.Open("postgres", cfg.DatabaseUrl)

	if err != nil {
		return nil, err
	}

	if err = db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}
