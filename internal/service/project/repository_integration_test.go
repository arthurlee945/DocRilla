package project_test

import (
	"context"
	"reflect"
	"testing"
	"time"

	"github.com/arthurlee945/Docrilla/internal/model"
	"github.com/arthurlee945/Docrilla/internal/model/mock"
	"github.com/arthurlee945/Docrilla/internal/service/project"
	"github.com/arthurlee945/Docrilla/internal/util"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

const DSN = "postgresql://public_user:Qwer1234@localhost:5432/docrilla?sslmode=disable"

func TestGetOverviewById(t *testing.T) {
	testProj := &model.Project{
		UUID:        mock.Project.UUID,
		Title:       mock.Project.Title,
		Description: mock.Project.Description,
		CreatedAt:   mock.Project.CreatedAt,
		VisitedAt:   mock.Project.VisitedAt,
	}
	dbConn, repo := repoPrep(t)
	defer dbConn.Close()

	ctx := context.Background()

	if proj, err := repo.GetOverviewById(ctx, "wrongid"); err == nil {
		t.Errorf("Expected GetOverviewById to return Error, but got err = %+v; proj = %+v", err, proj)
	}
	proj, err := repo.GetOverviewById(ctx, *mock.Project.UUID)
	if err != nil {
		t.Errorf("Expected GetOverviewById to return *model.Project. got = %+v", err)
	}
	if *proj.UUID == "" || *proj.Title == "" || *proj.Archived != false || proj.CreatedAt.IsZero() {
		t.Errorf("Expected GetOverviewById to return all values specified, but got = %+v", proj)
	}
	if reflect.DeepEqual(proj, testProj) {
		t.Errorf("Expected GetOverviewById project to equal to test project expected = %+v, got = %+v", testProj, proj)
	}
}

func TestGetDetailById(t *testing.T) {
	dbConn, repo := repoPrep(t)
	defer dbConn.Close()
	ctx := context.Background()
	if proj, err := repo.GetDetailById(ctx, "wrongid"); err == nil {
		t.Errorf("Expected GetDetailById to return Error, but got err = %+v; proj = %+v", err, proj)
	}
	proj, err := repo.GetDetailById(ctx, *mock.Project.UUID)
	if err != nil {
		t.Errorf("Expected GetOverviewById to return *model.Project. got = %+v", err)
	}
	if *proj.Title == "" || *proj.DocumentUrl == "" || proj.CreatedAt.IsZero() || proj.UpdatedAt.IsZero() {
		t.Errorf("Expected GetDetailById to contain proj base values, but got = %+v", proj)
	}

	if len(proj.Fields) != 2 {
		t.Errorf("Expected GetDetailById to contain 2 fields, but got = %d", len(proj.Fields))
	}
}

func TestCreateUpdateDeleteProject(t *testing.T) {
	dbConn, repo := repoPrep(t)
	defer dbConn.Close()
	ctx := context.Background()

	title, desc, docURL := "TEST TITLE", "TEST DESCRIPTION", "TEST DOC URL"

	mockProj := &model.Project{
		UserID:      mock.User.ID,
		Title:       util.ToPointer(title),
		Description: util.ToPointer(desc),
		DocumentUrl: util.ToPointer(docURL),
	}

	//CREATE
	newProj, err := repo.Create(ctx, mockProj)
	if err != nil {
		t.Errorf("Expected Create to return Project, but got err = %+v", err)
	}
	if *newProj.Title != title || *newProj.Description != desc || *newProj.DocumentUrl != docURL {
		t.Errorf("Expected Created New Project to contain correct values but got = %+v", newProj)
	}

	// UPDATE
	field1, field2 := mock.Field1, mock.Field2
	field1.ID, field2.ID = nil, nil
	field1.ProjectID, field2.ProjectID = newProj.ID, newProj.ID

	newTitle, newDesc, newDocURL := "NEW TEST TITLE", "NEW TEST DESCRIPTION", "NEW TEST DOC URL"
	newProj.Fields = []model.Field{field1, field2}
	newProj.Title = util.ToPointer(newTitle)
	newProj.Description = util.ToPointer(newDesc)
	newProj.DocumentUrl = util.ToPointer(newDocURL)
	newProj.VisitedAt = util.ToPointer(time.Now())

	if err := repo.Update(ctx, newProj); err != nil {
		t.Errorf("Expected Update to not throw but got err = %+v", err)
	}

	updatedProj, err := repo.GetDetailById(ctx, *newProj.UUID)
	if err != nil {
		t.Errorf("Expected GetDetailById after update project to not throw but got err = %+v", err)
	}
	if *updatedProj.Title != newTitle || *updatedProj.Description != newDesc || *updatedProj.DocumentUrl != newDocURL {
		t.Errorf("Expected Updated Project to contain correct values but got = %+v", newProj)
	}

	// DELETE
	if err := repo.Delete(ctx, *updatedProj.UUID); err != nil {
		t.Errorf("Expected Delete to not throw but got err = %+v", err)
	}
}

func repoPrep(t *testing.T) (*sqlx.DB, project.Repository) {
	db, err := sqlx.Open("postgres", DSN)

	if err != nil {
		t.Fatalf("Failed to initialize Test DB connection err=%+v", err)
	}

	if err = db.Ping(); err != nil {
		t.Fatalf("Failed to initialize Test DB connection err=%+v", err)
	}

	return db, project.NewRepository(db)
}
