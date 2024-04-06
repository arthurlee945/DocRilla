package project

import (
	"context"

	"github.com/arthurlee945/Docrilla/internal/model"
)

type Store interface {
	GetProjectOverview(ctx context.Context, user *model.User, uuid string) (*model.Project, error)
	GetProjectDetail(ctx context.Context, user *model.User, uuid string) (*model.Project, error)
	CreateProject(ctx context.Context, user *model.User, proj *model.Project) (string, error)
	UpdateProject(ctx context.Context, user *model.User, proj *model.Project) error
}

type Project struct {
	store Store
}

func New(s Store) *Project {
	return &Project{
		store: s,
	}
}
