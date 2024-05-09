package project_test

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/arthurlee945/Docrilla/internal/middleware"
	"github.com/arthurlee945/Docrilla/internal/service/field"
	"github.com/arthurlee945/Docrilla/internal/service/project"
	"github.com/jmoiron/sqlx"
)

func TestHandlers(t *testing.T) {
	handler, dbConn := handlerPrep(t)
	defer dbConn.Close()

	type args struct {
		req *http.Request
	}
	tests := []struct {
		name         string
		args         func(t *testing.T) args
		expectedCode int
		expectedBody string
	}{
		{name: "must return http.StatusOK for GetOverviewById to valid uuid",
			args: func(*testing.T) args {
				req, err := http.NewRequest("GET", "/projects/6be6167d-d25d-4ca4-9b6d-bfdc4e150f3d/overview", nil)
				if err != nil {
					t.Fatalf("failed to GetOverviewById: %s", err.Error())
				}
				return args{
					req,
				}
			},
			expectedCode: http.StatusOK,
			expectedBody: `{"uuid":"6be6167d-d25d-4ca4-9b6d-bfdc4e150f3d","userID":null,"route":"8daa1c06-c532-45e3-ad41-0a8b600dfa96","token":"TEST TOKEN","title":"TEST TITLE","description":"TEST DESCRIPTION","documentUrl":null,"archived":false,"visitedAt":null,"createdAt":"2024-05-02T20:33:40.234Z","updatedAt":null,"fields":null}`,
		},
		{name: "must return http.StatusOK for GetDetailById to valid uuid",
			args: func(*testing.T) args {
				req, err := http.NewRequest("GET", "/projects/8fd01895-37a7-4eaf-9b44-a73240e78eb9/detail", nil)
				if err != nil {
					t.Fatalf("failed to GetDetailById: %s", err.Error())
				}
				return args{
					req,
				}
			},
			expectedCode: http.StatusOK,
			expectedBody: `{"uuid":"8fd01895-37a7-4eaf-9b44-a73240e78eb9","userID":1,"route":"38388d5f-98b1-4364-8d42-39464b84722c","token":null,"title":"Postman says hi222","description":null,"documentUrl":"https://www.google.com/","archived":false,"visitedAt":null,"createdAt":"2024-05-08T20:25:39.787Z","updatedAt":"2024-05-08T20:25:39.787Z","fields":[]}`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tArgs := tt.args(t)
			resp := httptest.NewRecorder()
			handler.ServeHTTP(resp, tArgs.req)
			if resp.Result().StatusCode != tt.expectedCode {
				t.Fatalf("expected status code to be %d but got=%d", tt.expectedCode, resp.Result().StatusCode)
			}
			if strings.TrimSpace(resp.Body.String()) != tt.expectedBody {
				t.Fatalf("expected body to be %s but got=%s", tt.expectedBody, resp.Body.String())
			}
		})
	}
}

func handlerPrep(t *testing.T) (http.Handler, *sqlx.DB) {
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
	middleware.Auth(testMux)
	return middleware.Auth(testMux), db
}
