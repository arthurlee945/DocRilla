package project_test

import (
	"context"
	"testing"

	"github.com/arthurlee945/Docrilla/internal/service/project"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func TestServiceGet(t *testing.T) {
	dbConn, service := servicePrep(t)
	defer dbConn.Close()
	ctx := context.Background()
	//GET ALL
	getAllRequest := project.GetAllRequest{10, ""}
	_, _, err := service.GetAll(ctx, getAllRequest)
	if err != nil {
		t.Errorf("Expected Service GetAll to not get error but got = %+v", err)
	}
}

func servicePrep(t *testing.T) (*sqlx.DB, project.Service) {
	db, err := sqlx.Open("postgres", testDSN)
	if err != nil {
		t.Fatalf("Failed to initialize Test DB connection err=%+v", err)
	}

	if err = db.Ping(); err != nil {
		t.Fatalf("Failed to initialize Test DB connection err=%+v", err)
	}

	return db, project.NewService(project.NewRepository(db))

}
