package field_test

import (
	"context"
	"testing"

	"github.com/arthurlee945/Docrilla/internal/model/mock"
	"github.com/arthurlee945/Docrilla/internal/service/field"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func TestFieldService_CRUDValidation(t *testing.T) {
	dbConn, service := servicePrep(t)
	defer dbConn.Close()
	ctx := context.Background()

	//GetById
	correctFieldUUID, invalidFieldUUID := *mock.Field1.UUID, "test-id"

	_, getOverviewErr := service.GetById(ctx, invalidFieldUUID)
	if getOverviewErr == nil {
		t.Errorf("Expected GetById with incorrect ID Request to return Invalid ID error but got=nil")
	}
	if !field.ErrInvalidUUID.Is(getOverviewErr) {
		t.Errorf("Expected GetById Err to be invalid UUID but got=%+v", getOverviewErr)
	}
	if _, err := service.GetById(ctx, correctFieldUUID); err != nil {
		t.Errorf("Expected GetById to not return erro but got=%+v", err)
	}

	//GetDetailById
	_, getDetailErr := service.GetById(ctx, invalidFieldUUID)
	if getDetailErr == nil {
		t.Errorf("Expected GetDetailById with incorrect ID Request to return Invalid ID error but got=nil")
	}
	if !field.ErrInvalidUUID.Is(getDetailErr) {
		t.Errorf("Expected GetDetailById Err to be invalid UUID but got=%+v", getDetailErr)
	}
	if _, err := service.GetById(ctx, correctFieldUUID); err != nil {
		t.Errorf("Expected GetDetailById to not return erro but got=%+v", err)
	}

	//Create
	invalidCreateReq := field.CreateRequest{}
	createProj, createErr := service.Create(ctx, invalidCreateReq)
	if createErr == nil {
		t.Errorf("Expected Create to return error but got=%+v", createProj)
	}
	if !field.ErrInvalidReqObj.Is(createErr) {
		t.Errorf("Expected Create to retuirn ErrInvalidReqObj but got=%+v", createErr)
	}

	//Update
	invalidUpdateReq := field.UpdateRequest{}
	updateProj, udpateReqErr := service.Update(ctx, invalidUpdateReq)
	if udpateReqErr == nil {
		t.Errorf("Expected Update to return error but got=%+v", updateProj)
	}
	if !field.ErrInvalidReqObj.Is(udpateReqErr) {
		t.Errorf("Expected Update to retuirn ErrInvalidReqObj but got=%+v", udpateReqErr)
	}

	invalidUpdateReq.UUID = invalidFieldUUID
	invalidUpdateReq.ProjectID = invalidFieldUUID
	_, updateUUIDErr := service.Update(ctx, invalidUpdateReq)
	if getDetailErr == nil {
		t.Errorf("Expected Update with incorrect ID Request to return Invalid ID error but got=nil")
	}
	if !field.ErrInvalidUUID.Is(updateUUIDErr) {
		t.Errorf("Expected Update Err to be invalid UUID but got=%+v", updateUUIDErr)
	}

	//DELETE
	if err := service.Delete(ctx, invalidFieldUUID); err == nil {
		t.Errorf("Expected Update with incorrect ID Request to return Invalid ID error but got=nil")
	}
}

func servicePrep(t *testing.T) (*sqlx.DB, field.Service) {
	db, err := sqlx.Open("postgres", testDSN)
	if err != nil {
		t.Fatalf("Failed to initialize Test DB connection err=%+v", err)
	}

	if err = db.Ping(); err != nil {
		t.Fatalf("Failed to initialize Test DB connection err=%+v", err)
	}

	return db, field.NewService(field.NewRepository(db))
}
