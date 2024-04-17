package store_test

import (
	"context"
	"database/sql"
	"reflect"
	"testing"
	"time"

	"github.com/arthurlee945/Docrilla/internal/model"
	"github.com/arthurlee945/Docrilla/internal/model/mock"
	"github.com/arthurlee945/Docrilla/internal/model/null"
	"github.com/arthurlee945/Docrilla/internal/service/project/store"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

const DSN = "postgresql://public_user:Qwer1234@localhost:5432/docrilla?sslmode=disable"

func TestGetProjectOverview(t *testing.T) {
	testProj := &model.Project{
		UUID:        mock.Project.UUID,
		Title:       mock.Project.Title,
		Description: mock.Project.Description,
		CreatedAt:   mock.Project.CreatedAt,
		VisitedAt:   mock.Project.VisitedAt,
	}
	dbConn, store := storePrep(t)
	defer dbConn.Close()

	ctx := context.Background()

	if proj, err := store.GetProjectOverview(ctx, &mock.User, "wrongid"); err == nil {
		t.Errorf("Expected GetProjectOverview to return Error, but got err = %+v; proj = %+v", err, proj)
	}
	proj, err := store.GetProjectOverview(ctx, &mock.User, mock.Project.UUID)
	if err != nil {
		t.Errorf("Expected GetProjectOverview to return *model.Project. got = %+v", err)
	}
	if proj.UUID == "" || proj.Title == "" || proj.Archived != false || proj.CreatedAt.IsZero() {
		t.Errorf("Expected GetProjectOverview to return all values specified, but got = %+v", proj)
	}
	if reflect.DeepEqual(proj, testProj) {
		t.Errorf("Expected GetProjectOverview project to equal to test project expected = %+v, got = %+v", testProj, proj)
	}
}

func TestGetProjectDetail(t *testing.T) {
	dbConn, store := storePrep(t)
	defer dbConn.Close()
	ctx := context.Background()
	if proj, err := store.GetProjectDetail(ctx, &mock.User, "wrongid"); err == nil {
		t.Errorf("Expected GetProjectDetail to return Error, but got err = %+v; proj = %+v", err, proj)
	}
	proj, err := store.GetProjectDetail(ctx, &mock.User, mock.Project.UUID)
	if err != nil {
		t.Errorf("Expected GetProjectOverview to return *model.Project. got = %+v", err)
	}
	if proj.Title == "" || proj.DocumentUrl == "" || proj.CreatedAt.IsZero() || proj.UpdatedAt.IsZero() {
		t.Errorf("Expected GetProjectDetail to contain proj base values, but got = %+v", proj)
	}

	if len(proj.Fields) != 2 {
		t.Errorf("Expected GetProjectDetail to contain 2 fields, but got = %d", len(proj.Fields))
	}
}

func TestCreateUpdateDeleteProject(t *testing.T) {
	dbConn, store := storePrep(t)
	defer dbConn.Close()
	ctx := context.Background()

	title, desc, docURL := "TEST TITLE", "TEST DESCRIPTION", "TEST DOC URL"

	mockProj := &model.Project{
		UserID:      mock.User.ID,
		Title:       title,
		Description: null.String{NullString: sql.NullString{String: desc}},
	}

	// CREATE
	// if proj, err := store.CreateProject(ctx, &mock.User, mockProj); err == nil {
	// 	t.Errorf("Expected CreateProject to return Error, but got err = %+v; proj = %+v", err, proj)
	// }

	mockProj.DocumentUrl = docURL
	newProj, err := store.CreateProject(ctx, &mock.User, mockProj)
	if err != nil {
		t.Errorf("Expected CreateProject to return Project, but got err = %+v", err)
	}
	if newProj.Title != title || newProj.Description.String == desc || newProj.DocumentUrl != docURL {
		t.Errorf("Expected Created New Project to contain correct values but got = %+v", newProj)
	}

	// UPDATE
	field1, field2 := mock.Field1, mock.Field2
	field1.ProjectID = newProj.ID
	field2.ProjectID = newProj.ID

	newTitle, newDesc, newDocURL := "NEW TEST TITLE", "NEW TEST DESCRIPTION", "NEW TEST DOC URL"
	newProj.Fields = []model.Field{field1, field2}
	newProj.Title = newTitle
	newProj.Description = null.String{NullString: sql.NullString{String: newDesc}}
	newProj.DocumentUrl = newDocURL
	newProj.VisitedAt = null.Time{NullTime: sql.NullTime{Time: time.Now()}}

	if err := store.UpdateProject(ctx, &mock.User, newProj); err != nil {
		t.Errorf("Expected UpdateProject to not throw but got err = %+v", err)
	}
	updatedProj, err := store.GetProjectDetail(ctx, &mock.User, newProj.UUID)
	if err != nil {
		t.Errorf("Expected GetProjectDetail after update project to not throw but got err = %+v", err)
	}
	if updatedProj.Title != newTitle || updatedProj.Description.String != newDesc || updatedProj.DocumentUrl != newDocURL {
		t.Errorf("Expected Updated Project to contain correct values but got = %+v", newProj)
	}

	// DELETE
	if err := store.DeleteProject(ctx, &mock.User, updatedProj); err != nil {
		t.Errorf("Expected DeleteProject to not throw but got err = %+v", err)
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
