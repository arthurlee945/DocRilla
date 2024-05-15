package mock

import (
	"github.com/arthurlee945/Docrilla/internal/model"
	"github.com/arthurlee945/Docrilla/internal/model/enum/field"
	"github.com/arthurlee945/Docrilla/internal/model/enum/user"
	"github.com/arthurlee945/Docrilla/internal/util/ptr"
)

var (
	User = model.User{
		ID:       ptr.ToPointer[uint64](1),
		Name:     ptr.ToPointer("admin"),
		Email:    ptr.ToPointer("admin@admin.com"),
		Password: ptr.ToPointer("$2a$14$Up0V3zSFmLCtuYMrVnwbSOi/Z76dAja6ZqTU22akvbNYZQFVBmn7G"), //qwer1234
		Role:     ptr.ToPointer(user.MOCK),
	}

	Account = model.Account{
		UserID:  User.ID,
		Type:    ptr.ToPointer("SEED"),
		Provide: ptr.ToPointer("SEED"),
	}

	Project = model.Project{
		ID:          ptr.ToPointer[uint64](1),
		UUID:        ptr.ToPointer("6be6167d-d25d-4ca4-9b6d-bfdc4e150f3d"),
		UserID:      User.ID,
		Route:       ptr.ToPointer("TEST ROUTE"),
		Token:       ptr.ToPointer("TEST TOKEN"),
		Title:       ptr.ToPointer("TEST TITLE"),
		Description: ptr.ToPointer("TEST DESCRIPTION"),
		DocumentUrl: ptr.ToPointer("NO URL"),
	}

	Field1 = model.Field{
		ID:        ptr.ToPointer[uint64](1),
		UUID:      ptr.ToPointer("4f836019-e3d0-4ddf-8d4d-93d41eb2c01b"),
		ProjectID: Project.UUID,
		X:         ptr.ToPointer[float64](0),
		Y:         ptr.ToPointer[float64](0),
		Width:     ptr.ToPointer[float64](24),
		Height:    ptr.ToPointer[float64](24),
		Page:      ptr.ToPointer[uint32](1),
		Type:      ptr.ToPointer(field.TEXT),
	}

	Field2 = model.Field{
		ID:        ptr.ToPointer[uint64](2),
		UUID:      ptr.ToPointer("8d566b8a-10be-4610-be14-316d0313fce0"),
		ProjectID: Project.UUID,
		X:         ptr.ToPointer[float64](34),
		Y:         ptr.ToPointer[float64](50),
		Width:     ptr.ToPointer[float64](524),
		Height:    ptr.ToPointer[float64](124),
		Page:      ptr.ToPointer[uint32](1),
		Type:      ptr.ToPointer(field.TEXT),
	}
)
