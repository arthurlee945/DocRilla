package mock

import (
	"database/sql"

	"github.com/arthurlee945/Docrilla/internal/model"
	"github.com/arthurlee945/Docrilla/internal/model/enum/field"
	"github.com/arthurlee945/Docrilla/internal/model/enum/user"
	"github.com/arthurlee945/Docrilla/internal/model/null"
)

var (
	User = model.User{
		ID:       1,
		Name:     "admin",
		Email:    "admin@admin.com",
		Password: null.String{NullString: sql.NullString{String: "qwer1234"}},
		Role:     user.MOCK,
	}

	Account = model.Account{
		UserID:  User.ID,
		Type:    "SEED",
		Provide: "SEED",
	}

	Project = model.Project{
		ID:          1,
		UUID:        "6be6167d-d25d-4ca4-9b6d-bfdc4e150f3d",
		UserID:      User.ID,
		Title:       "TEST TITLE",
		Description: null.String{NullString: sql.NullString{String: "TEST DESCRIPTION"}},
		DocumentUrl: "NO URL",
	}

	Field1 = model.Field{
		UUID:      "4f836019-e3d0-4ddf-8d4d-93d41eb2c01b",
		ProjectID: Project.ID,
		X1:        0,
		Y1:        0,
		X2:        24,
		Y2:        24,
		Page:      1,
		Type:      field.TEXT,
	}

	Field2 = model.Field{
		UUID:      "8d566b8a-10be-4610-be14-316d0313fce0",
		ProjectID: Project.ID,
		X1:        50,
		Y1:        50,
		X2:        524,
		Y2:        124,
		Page:      1,
		Type:      field.TEXT,
	}
)
