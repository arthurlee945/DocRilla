package project_test

import (
	"context"
	"testing"

	"github.com/arthurlee945/Docrilla/internal/model/mock"
	"github.com/arthurlee945/Docrilla/internal/service/project"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func TestServiceCRUDValidation(t *testing.T) {
	dbConn, service := servicePrep(t)
	defer dbConn.Close()
	ctx := context.Background()

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
	if !project.ErrInvalidUUID.Is(getOverviewErr) {
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
	if !project.ErrInvalidUUID.Is(getDetailErr) {
		t.Errorf("Expected GetDetailById Err to be invalid UUID but got=%+v", getDetailErr)
	}
	if _, err := service.GetDetailById(ctx, correctProjUUID); err != nil {
		t.Errorf("Expected GetDetailById to not return erro but got=%+v", err)
	}

	//Create
	invalidCreateReq := project.CreateRequest{
		UserID: *mock.User.ID,
	}
	createProj, createErr := service.Create(ctx, invalidCreateReq)
	if createErr == nil {
		t.Errorf("Expected Create to return error but got=%+v", createProj)
	}
	if !project.ErrInvalidReqObj.Is(createErr) {
		t.Errorf("Expected Create to retuirn ErrInvalidReqObj but got=%+v", createErr)
	}

	//Update
	invalidUpdateReq := project.UpdateRequest{}
	updateProj, udpateReqErr := service.Update(ctx, invalidUpdateReq)
	if udpateReqErr == nil {
		t.Errorf("Expected Update to return error but got=%+v", updateProj)
	}
	if !project.ErrInvalidReqObj.Is(udpateReqErr) {
		t.Errorf("Expected Update to retuirn ErrInvalidReqObj but got=%+v", udpateReqErr)
	}

	invalidUpdateReq.UUID = invalidProjUUID
	_, updateUUIDErr := service.Update(ctx, invalidUpdateReq)
	if getDetailErr == nil {
		t.Errorf("Expected Update with incorrect ID Request to return Invalid ID error but got=nil")
	}
	if !project.ErrInvalidUUID.Is(updateUUIDErr) {
		t.Errorf("Expected Update Err to be invalid UUID but got=%+v", updateUUIDErr)
	}

	//DELETE
	if err := service.Delete(ctx, invalidProjUUID); err == nil {
		t.Errorf("Expected Update with incorrect ID Request to return Invalid ID error but got=nil")
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
