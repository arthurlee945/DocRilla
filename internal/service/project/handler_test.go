package project_test

import (
	"net/http"
	"testing"

	"github.com/arthurlee945/Docrilla/internal/service/field"
	"github.com/arthurlee945/Docrilla/internal/service/project"
	"github.com/jmoiron/sqlx"
)

func TestGetHandlers(t *testing.T) {
	_, dbConn := handlerPrep(t)
	defer dbConn.Close()
}

func handlerPrep(t *testing.T) (*http.ServeMux, *sqlx.DB) {
	db, err := sqlx.Open("postgres", testDSN)

	if err != nil {
		t.Fatalf("Failed to initialize Test DB connection err=%+v", err)
	}

	if err = db.Ping(); err != nil {
		t.Fatalf("Failed to initialize Test DB connection err=%+v", err)
	}
	testMux := http.NewServeMux()
	projService := project.NewService(project.NewRepository(db), field.NewRepository(db))
	project.RegisterHandler(testMux, projService)
	return testMux, db
}
