package mock

import (
	"github.com/arthurlee945/Docrilla/internal/model"
	"github.com/arthurlee945/Docrilla/internal/model/enum/field"
	"github.com/arthurlee945/Docrilla/internal/model/enum/user"
	"github.com/arthurlee945/Docrilla/internal/util"
)

var (
	User = model.User{
		ID:       util.ToPointer[uint64](1),
		Name:     util.ToPointer("admin"),
		Email:    util.ToPointer("admin@admin.com"),
		Password: util.ToPointer("qwer1234"),
		Role:     util.ToPointer(user.MOCK),
	}

	Account = model.Account{
		UserID:  User.ID,
		Type:    util.ToPointer("SEED"),
		Provide: util.ToPointer("SEED"),
	}

	Project = model.Project{
		ID:          util.ToPointer[uint64](1),
		UUID:        util.ToPointer("6be6167d-d25d-4ca4-9b6d-bfdc4e150f3d"),
		UserID:      User.ID,
		Route:       util.ToPointer("TEST ROUTE"),
		Token:       util.ToPointer("TEST TOKEN"),
		Title:       util.ToPointer("TEST TITLE"),
		Description: util.ToPointer("TEST DESCRIPTION"),
		DocumentUrl: util.ToPointer("NO URL"),
	}

	Field1 = model.Field{
		ID:        util.ToPointer[uint64](1),
		UUID:      util.ToPointer("4f836019-e3d0-4ddf-8d4d-93d41eb2c01b"),
		ProjectID: Project.UUID,
		X:         util.ToPointer[float64](0),
		Y:         util.ToPointer[float64](0),
		Width:     util.ToPointer[float64](24),
		Height:    util.ToPointer[float64](24),
		Page:      util.ToPointer[uint32](1),
		Type:      util.ToPointer(field.TEXT),
	}

	Field2 = model.Field{
		ID:        util.ToPointer[uint64](2),
		UUID:      util.ToPointer("8d566b8a-10be-4610-be14-316d0313fce0"),
		ProjectID: Project.UUID,
		X:         util.ToPointer[float64](34),
		Y:         util.ToPointer[float64](50),
		Width:     util.ToPointer[float64](524),
		Height:    util.ToPointer[float64](124),
		Page:      util.ToPointer[uint32](1),
		Type:      util.ToPointer(field.TEXT),
	}
)
