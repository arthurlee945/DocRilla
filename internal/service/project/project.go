package project

import (
	"context"

	"github.com/arthurlee945/Docrilla/internal/model"
)

type Repository interface {
	GetProjectOverview(ctx context.Context, user *model.User, uuid string) (*model.Project, error)
	GetProjectDetail(ctx context.Context, user *model.User, uuid string) (*model.Project, error)
	CreateProject(ctx context.Context, user *model.User, proj *model.Project) (string, error)
	UpdateProject(ctx context.Context, user *model.User, proj *model.Project) error
}

type Project struct {
	repository Repository
}

func New(r Repository) *Project {
	return &Project{
		repository: r,
	}
}
