package store_test

import (
	"context"
	"database/sql"
	"log"
	"testing"

	"github.com/arthurlee945/Docrilla/internal/config"
	"github.com/arthurlee945/Docrilla/internal/db"
	"github.com/arthurlee945/Docrilla/internal/model"
	"github.com/jmoiron/sqlx"
)

var MockUser = model.User{
	Name:     "admin",
	Email:    "admin@admin.com",
	Password: sql.NullString{String: "qwer1234"},
	Role:     "ADMIN",
}

func TestGetProjectOverview(t *testing.T) {
	dbConn := getDBConn()
	defer dbConn.Close()
}

func getDBConn() *sqlx.DB {
	cfg, err := config.Load(context.Background())
	if err != nil {
		log.Fatalln(err)
	}
	dbConn, err := db.Connect(cfg.DSN)
	if err != nil {
		log.Fatalln(err)
	}
	return dbConn
}
