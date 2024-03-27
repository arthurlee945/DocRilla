package db

import (
	"context"
	"database/sql"
	"time"

	"github.com/arthurlee945/Docrilla/config"
	_ "github.com/lib/pq"
)

var cfg = config.New()

func NewConnection() (*sql.DB, error) {
	db, err := sql.Open("postgres", cfg.DatabaseUrl)

	if err != nil {
		return nil, err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if _, err := db.Conn(ctx); err != nil {
		return nil, err
	}

	if err = db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}
