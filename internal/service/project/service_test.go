package project_test

import (
	"context"
	"testing"

	"github.com/arthurlee945/Docrilla/internal/errors"
	"github.com/arthurlee945/Docrilla/internal/model/mock"
	"github.com/arthurlee945/Docrilla/internal/service/auth"
	"github.com/arthurlee945/Docrilla/internal/service/field"
	"github.com/arthurlee945/Docrilla/internal/service/project"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func TestProjectService_CRUDValidation(t *testing.T) {
	dbConn, service, ctx := servicePrep(t)
	defer func() {
		if err := dbConn.Close(); err != nil {
			t.Fatalf("Tried to close DB connection but got=%v", err)
		}
	}()

	//GetAll
	getAllRequest := project.GetAllRequest{10, ""}
	// TODO: add cursor test

	if _, _, err := service.GetAll(ctx, getAllRequest); err != nil {
		t.Errorf("Expected Service GetAll to not get error but got = %+v", err)
	}

	//GetOverviewById
	correctProjUUID, invalidProjUUID := *mock.Project.UUID, "test-id"

	_, getOverviewErr := service.GetOverviewById(ctx, invalidProjUUID)
	if getOverviewErr == nil {
		t.Errorf("Expected GetOverviewById with incorrect ID Request to return Invalid ID error but got=nil")
	}
	if !errors.ErrInvalidRequest.Is(getOverviewErr) {
		t.Errorf("Expected GetOverviewById Err to be invalid UUID but got=%+v", getOverviewErr)
	}
	if _, err := service.GetOverviewById(ctx, correctProjUUID); err != nil {
		t.Errorf("Expected GetOverviewById to not return erro but got=%+v", err)
	}

	//GetDetailById
	_, getDetailErr := service.GetDetailById(ctx, invalidProjUUID)
	if getDetailErr == nil {
		t.Errorf("Expected GetDetailById with incorrect ID Request to return Invalid ID error but got=nil")
	}
	if !errors.ErrInvalidRequest.Is(getDetailErr) {
		t.Errorf("Expected GetDetailById Err to be invalid UUID but got=%+v", getDetailErr)
	}
	if _, err := service.GetDetailById(ctx, correctProjUUID); err != nil {
		t.Errorf("Expected GetDetailById to not return erro but got=%+v", err)
	}

	//Create
	invalidCreateReq := project.CreateRequest{}
	createProj, createErr := service.Create(ctx, invalidCreateReq)
	if createErr == nil {
		t.Errorf("Expected Create to return error but got=%+v", createProj)
	}
	if !errors.ErrValidation.Is(createErr) {
		t.Errorf("Expected Create to retuirn ErrInvalidReqObj but got=%+v", createErr)
	}

	//Update
	invalidUpdateReq := project.UpdateRequest{}
	updateProj, udpateReqErr := service.Update(ctx, invalidUpdateReq)
	if udpateReqErr == nil {
		t.Errorf("Expected Update to return error but got=%+v", updateProj)
	}
	if !errors.ErrValidation.Is(udpateReqErr) {
		t.Errorf("Expected Update to retuirn ErrInvalidReqObj but got=%+v", udpateReqErr)
	}

	invalidUpdateReq.UUID = invalidProjUUID
	_, updateUUIDErr := service.Update(ctx, invalidUpdateReq)
	if getDetailErr == nil {
		t.Errorf("Expected Update with incorrect ID Request to return Invalid ID error but got=nil")
	}
	if !errors.ErrInvalidRequest.Is(updateUUIDErr) {
		t.Errorf("Expected Update Err to be invalid UUID but got=%+v", updateUUIDErr)
	}

	//DELETE
	if err := service.Delete(ctx, invalidProjUUID); err == nil {
		t.Errorf("Expected Update with incorrect ID Request to return Invalid ID error but got=nil")
	}
}

func TestProjectService_CreateUpdateDelete(t *testing.T) {
	dbConn, service, ctx := servicePrep(t)
	fieldService := field.NewService(field.NewRepository(dbConn))
	defer func() {
		if err := dbConn.Close(); err != nil {
			t.Fatalf("Tried to close DB connection but got=%v", err)
		}
	}()

	mockProj, err := service.Create(ctx, project.CreateRequest{
		Title:       *mock.Project.Title,
		DocumentUrl: *mock.Project.DocumentUrl,
	})
	if err != nil {
		t.Fatalf("Expected Project service CREATE to return proj but got=%+v", err)
	}
	mockField, err := fieldService.Create(ctx, field.CreateRequest{
		ProjectId: *mockProj.UUID,
		X:         *mock.Field1.X,
		Y:         *mock.Field1.Y,
		Width:     *mock.Field1.Width,
		Height:    *mock.Field1.Height,
		Type:      *mock.Field1.Type,
		Page:      *mock.Field1.Page,
	})
	if err != nil {
		t.Fatalf("Expected Field service CREATE to return field but got=%+v", err)
	}

	uTitle, uDescription, uToken := "TEST UPDATE TITLE", "TEST UPDATE DESCRIPTION", "TEST TOKEN TOKEN"
	uWidth, uHeight := 335.3, 293.5
	uProj, uErr := service.Update(ctx, project.UpdateRequest{
		UUID:        *mockProj.UUID,
		Title:       &uTitle,
		Description: &uDescription,
		Token:       &uToken,
		Fields: []field.UpdateRequest{
			{
				UUID:      *mockField.UUID,
				ProjectID: *mockProj.UUID,
				Width:     &uWidth,
				Height:    &uHeight,
			},
		},
	})
	if uErr != nil {
		t.Errorf("Expected Project Service UPDATE to return updated model.Project but got=%+v", uErr)
	}
	if *uProj.Title != uTitle || *uProj.Description != uDescription || *uProj.Token != uToken || len(uProj.Fields) != 1 {
		t.Errorf("Expected Project UPDATE to return correct value but got=%+v", map[string]any{
			"Title":      *uProj.Title,
			"Desc":       *uProj.Description,
			"Token":      *uProj.Token,
			"fieldLenth": len(uProj.Fields),
		})
	}
	uField := uProj.Fields[0]
	if *uField.Width != uWidth || *uField.Height != uHeight {
		t.Errorf("Expected Project UPDATE Field to return correct value but got=%+v", map[string]any{
			"Width":  *uField.Width,
			"Height": *uField.Height,
		})
	}

	if err := service.Delete(ctx, *uProj.UUID); err != nil {
		t.Errorf("Expected Project Service DELETE to success but got=%+v", err)
	}
}

func servicePrep(t *testing.T) (*sqlx.DB, project.Service, context.Context) {
	db, err := sqlx.Open("postgres", testDSN)
	ctx := context.Background()
	if err != nil {
		t.Fatalf("Failed to initialize Test DB connection err=%+v", err)
	}

	if err = db.Ping(); err != nil {
		t.Fatalf("Failed to initialize Test DB connection err=%+v", err)
	}

	return db, project.NewService(project.NewRepository(db), field.NewRepository(db)), context.WithValue(ctx, auth.AuthKey, *mock.User.ID)
}
