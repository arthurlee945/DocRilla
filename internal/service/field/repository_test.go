package field_test

import (
	"context"
	"testing"

	"github.com/arthurlee945/Docrilla/internal/model"
	fieldEnum "github.com/arthurlee945/Docrilla/internal/model/enum/field"
	"github.com/arthurlee945/Docrilla/internal/model/mock"
	"github.com/arthurlee945/Docrilla/internal/service/field"
	"github.com/arthurlee945/Docrilla/internal/util"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func TestFieldRepositoryGetById(t *testing.T) {
	dbConn, repo := repoPrep(t)
	defer dbConn.Close()

	ctx := context.Background()

	if field, err := repo.GetById(ctx, "wrongid"); err == nil {
		t.Errorf("Expected Get to return Error, but got err = %+v; field = %+v", err, field)
	}
	field, err := repo.GetById(ctx, *mock.Field1.UUID)
	if err != nil {
		t.Errorf("Expected Get to return *model.Field. got = %+v", err)
	}
	if *field.UUID != *mock.Field1.UUID ||
		*field.X1 != *mock.Field1.X1 ||
		*field.Y1 != *mock.Field1.Y1 ||
		*field.X2 != *mock.Field1.X2 ||
		*field.Y2 != *mock.Field1.Y2 ||
		*field.Page != *mock.Field1.Page ||
		*field.Type != *mock.Field1.Type {
		t.Errorf("Expected Get to return all values specified, but got = %+v", map[string]any{
			"uuid":     *field.UUID,
			"mockUUID": *mock.Field1.UUID,
		})
	}

}

func TestFieldRepositoryCreateUpdateDeleteProject(t *testing.T) {
	dbConn, repo := repoPrep(t)
	defer dbConn.Close()
	ctx := context.Background()

	x1, y1, x2, y2, uType := 8.0, 88.0, 88.8, 88.88, fieldEnum.TEXT

	mockField := &model.Field{
		ProjectID: mock.Project.ID,
		X1:        util.ToPointer(x1),
		Y1:        util.ToPointer(y1),
		X2:        util.ToPointer(x2),
		Y2:        util.ToPointer(y2),
		Page:      util.ToPointer[uint32](2),
		Type:      util.ToPointer(uType),
	}
	//CREATE
	newField, err := repo.Create(ctx, mockField)
	if err != nil {
		t.Errorf("Expected Create to return Project, but got err = %+v", err)
	}
	if *newField.X1 != x1 || *newField.Y1 != y1 || *newField.X2 != x2 || *newField.Y2 != y2 || *newField.Type != uType {
		t.Errorf("Expected Created New Project to contain correct values but got = %+v", newField)
	}

	// UPDATE

	newPos := [4]float64{1, 2, 3, 4}
	newField.X1 = util.ToPointer(newPos[0])
	newField.Y1 = util.ToPointer(newPos[1])
	newField.X2 = util.ToPointer(newPos[2])
	newField.Y2 = util.ToPointer(newPos[3])

	updatedField, err := repo.Update(ctx, newField)
	if err != nil {
		t.Errorf("Expected Update to not throw but got err = %+v", err)
	}

	if *updatedField.X1 != newPos[0] || *updatedField.Y1 != newPos[1] || *updatedField.X2 != newPos[2] || *updatedField.Y2 != newPos[3] {
		t.Errorf("Expected Updated Project to contain correct values but got = %+v", updatedField)
	}

	// DELETE
	if err := repo.Delete(ctx, *updatedField.UUID); err != nil {
		t.Errorf("Expected Delete to not throw but got err = %+v", err)
	}
}

func repoPrep(t *testing.T) (*sqlx.DB, field.Repository) {
	db, err := sqlx.Open("postgres", testDSN)

	if err != nil {
		t.Fatalf("Failed to initialize Test DB connection err=%+v", err)
	}

	if err = db.Ping(); err != nil {
		t.Fatalf("Failed to initialize Test DB connection err=%+v", err)
	}

	return db, field.NewRepository(db)
}
