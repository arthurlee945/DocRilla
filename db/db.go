package db

import (
	"github.com/arthurlee945/Docrilla/config"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

var cfg = config.New()

func NewConnection() (*sqlx.DB, error) {
	db, err := sqlx.Open("postgres", cfg.DatabaseUrl)

	if err != nil {
		return nil, err
	}

	if err = db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}
