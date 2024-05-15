package project_test

import (
	"context"
	"testing"
	"time"

	"github.com/arthurlee945/Docrilla/internal/model"
	"github.com/arthurlee945/Docrilla/internal/model/mock"
	"github.com/arthurlee945/Docrilla/internal/service/project"
	"github.com/arthurlee945/Docrilla/internal/util/ptr"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func TestProjectRepository_GetAll(t *testing.T) {
	dbConn, repo := repoPrep(t)
	defer dbConn.Close()

	ctx := context.Background()
	projs, nextCursor, err := repo.GetAll(ctx, 10, "", *mock.User.ID)
	if err != nil {
		t.Errorf("Expected GetAll to return Error, but got err = %+v; projs = %+v; nextCursor = %s", err, projs, nextCursor)
	}
	if len(projs) < 1 {
		t.Errorf("Expected existing project length to be 1 or more but got = %d", len(projs))
	}
}

func TestProjectRepository_GetOverviewById(t *testing.T) {
	dbConn, repo := repoPrep(t)
	defer dbConn.Close()

	ctx := context.Background()

	if proj, err := repo.GetOverviewById(ctx, "wrongid", *mock.User.ID); err == nil {
		t.Errorf("Expected GetOverviewById to return Error, but got err = %+v; proj = %+v", err, proj)
	}
	proj, err := repo.GetOverviewById(ctx, *mock.Project.UUID, *mock.User.ID)
	if err != nil {
		t.Errorf("Expected GetOverviewById to return *model.Project. got = %+v", err)
	}
	if *proj.UUID == "" || *proj.Title == "" || *proj.Archived != false || proj.CreatedAt.IsZero() {
		t.Errorf("Expected GetOverviewById to return all values specified, but got = %+v", proj)
	}
}

func TestProjectRepository_GetDetailById(t *testing.T) {
	dbConn, repo := repoPrep(t)
	defer dbConn.Close()
	ctx := context.Background()
	if proj, err := repo.GetDetailById(ctx, "wrongid", *mock.User.ID); err == nil {
		t.Errorf("Expected GetDetailById to return Error, but got err = %+v; proj = %+v", err, proj)
	}
	proj, err := repo.GetDetailById(ctx, *mock.Project.UUID, *mock.User.ID)
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

func TestProjectRepository_CreateUpdateDeleteProject(t *testing.T) {
	dbConn, repo := repoPrep(t)
	defer dbConn.Close()
	ctx := context.Background()

	title, desc, docURL, testRoute := "TEST TITLE", "TEST DESCRIPTION", "TEST DOC URL", "test-route"+time.Now().String()

	mockProj := &model.Project{
		UserID:      mock.User.ID,
		Title:       ptr.ToPointer(title),
		Description: ptr.ToPointer(desc),
		DocumentUrl: ptr.ToPointer(docURL),
		Route:       ptr.ToPointer(testRoute),
	}

	//CREATE
	newProj, err := repo.Create(ctx, mockProj)
	if err != nil {
		t.Errorf("Expected Create to return Project, but got err = %+v", err)
	}
	if *newProj.Title != title || *newProj.Description != desc || *newProj.DocumentUrl != docURL || *newProj.Route != testRoute {
		t.Errorf("Expected Created New Project to contain correct values but got = %+v", newProj)
	}

	// UPDATE

	newTitle, newDesc, newDocURL := "NEW TEST TITLE", "NEW TEST DESCRIPTION", "NEW TEST DOC URL"
	newProj.Title = ptr.ToPointer(newTitle)
	newProj.Description = ptr.ToPointer(newDesc)
	newProj.DocumentUrl = ptr.ToPointer(newDocURL)
	newProj.VisitedAt = ptr.ToPointer(time.Now())
	updatedProj, err := repo.Update(ctx, newProj)
	if err != nil {
		t.Errorf("Expected Update to not throw but got err = %+v", err)
	}

	if *updatedProj.Title != newTitle || *updatedProj.Description != newDesc || *updatedProj.DocumentUrl != newDocURL {
		t.Errorf("Expected Updated Project to contain correct values but got = %+v", newProj)
	}

	// DELETE
	if err := repo.Delete(ctx, *updatedProj.UUID, *mock.User.ID); err != nil {
		t.Errorf("Expected Delete to not throw but got err = %+v", err)
	}
}

func repoPrep(t *testing.T) (*sqlx.DB, project.Repository) {
	db, err := sqlx.Open("postgres", testDSN)

	if err != nil {
		t.Fatalf("Failed to initialize Test DB connection err=%+v", err)
	}

	if err = db.Ping(); err != nil {
		t.Fatalf("Failed to initialize Test DB connection err=%+v", err)
	}

	return db, project.NewRepository(db)
}
