package store_test

import (
	"context"
	"reflect"
	"testing"

	"github.com/arthurlee945/Docrilla/internal/model"
	"github.com/arthurlee945/Docrilla/internal/model/test"
	"github.com/arthurlee945/Docrilla/internal/service/project/store"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

const DSN = "postgresql://public_user:Qwer1234@localhost:5432/docrilla?sslmode=disable"

func TestGetProjectOverview(t *testing.T) {
	sampleProj := &model.Project{
		UUID:        test.Project.UUID,
		Title:       test.Project.Title,
		Description: test.Project.Description,
		CreatedAt:   test.Project.CreatedAt,
		VisitedAt:   test.Project.VisitedAt,
	}
	dbConn, store := storePrep(t)
	defer dbConn.Close()

	ctx := context.Background()
	proj, err := store.GetProjectOverview(ctx, test.User, test.Project.UUID)
	if err != nil {
		t.Errorf("Expected GetProjectOverview to return *model.Project. got = %+v", err)
	}

	if reflect.DeepEqual(proj, sampleProj) {
		t.Errorf("Expected GetProjectOverview project to equal to test project expected = %+v, got = %+v", sampleProj, proj)
	}
}

func storePrep(t *testing.T) (*sqlx.DB, *store.Store) {
	db, err := sqlx.Open("postgres", DSN)

	if err != nil {
		t.Fatalf("Failed to initialize Test DB connection err=%+v", err)
	}

	if err = db.Ping(); err != nil {
		t.Fatalf("Failed to initialize Test DB connection err=%+v", err)
	}

	return db, store.NewStore(db)
}
