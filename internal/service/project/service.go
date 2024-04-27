package project

import (
	"context"

	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"

	"github.com/arthurlee945/Docrilla/internal/model"
	"github.com/arthurlee945/Docrilla/internal/util"
)

var validate = validator.New(validator.WithRequiredStructEnabled())

type Service interface {
	GetAll(context.Context, GetAllRequest) (projects []model.Project, nextCursor string, err error)
}

type Project struct {
	model.Project
}

type service struct {
	repo Repository
}

func NewService(r Repository) Service {
	return &service{
		repo: r,
	}
}

type GetAllRequest struct {
	Limit  uint8
	Cursor string
}

func (s *service) GetAll(ctx context.Context, req GetAllRequest) ([]model.Project, string, error) {
	projects, nextCursor, err := s.repo.GetAll(ctx, req.Limit, req.Cursor)
	if err != nil {
		return nil, "", ErrRepoGet.Wrap(err)
	}
	return projects, nextCursor, nil
}

func (s *service) GetOverviewById(ctx context.Context, id string) (*model.Project, error) {
	if err := uuid.Validate(id); err != nil {
		return nil, ErrInvalidUUID.Wrap(err)
	}
	project, err := s.repo.GetOverviewById(ctx, id)
	if err != nil {
		return nil, ErrRepoGet.Wrap(err)
	}
	return project, nil
}

func (s *service) GetDetailById(ctx context.Context, id string) (*model.Project, error) {
	if err := uuid.Validate(id); err != nil {
		return nil, ErrInvalidUUID.Wrap(err)
	}
	project, err := s.repo.GetDetailById(ctx, id)
	if err != nil {
		return nil, ErrRepoGet.Wrap(err)
	}
	return project, nil
}

type CreateRequest struct {
	UserID      uint64
	Title       string `validate:"required"`
	Description *string
	Route       *string
	Token       *string
	DocumentUrl string `validate:"required"`
}

func (s *service) Create(ctx context.Context, req CreateRequest) (*model.Project, error) {
	if err := validate.Struct(req); err != nil {
		return nil, ErrInvalidReqObj.Wrap(err)
	}
	createdProj, err := s.repo.Create(ctx, &model.Project{
		UserID:      util.ToPointer(req.UserID),
		Title:       util.ToPointer(req.Title),
		Description: req.Description,
		Route:       req.Route,
		Token:       req.Token,
		DocumentUrl: util.ToPointer(req.DocumentUrl),
	})
	if err != nil {
		return nil, ErrServiceCreate.Wrap(err)
	}
	return createdProj, nil
}

// maybe need to seperate field update
type UpdateRequest struct {
	UUID        string `validate:"required"`
	Title       *string
	Description *string
	DocumentUrl *string
	Route       *string
	Token       *string
}

// ADD Field repo and update this
func (s *service) Update(ctx context.Context, req UpdateRequest) (*model.Project, error) {
	if err := validate.Struct(req); err != nil {
		return nil, ErrInvalidReqObj.Wrap(err)
	}
	if err := uuid.Validate(req.UUID); err != nil {
		return nil, ErrInvalidUUID.Wrap(err)
	}
	updatedProj, err := s.repo.Update(ctx, &model.Project{
		UUID:        util.ToPointer(req.UUID),
		Title:       req.Title,
		Description: req.Description,
		DocumentUrl: req.DocumentUrl,
		Route:       req.Route,
		Token:       req.Token,
	})
	if err != nil {
		return nil, ErrServiceUpdate.Wrap(err)
	}
	return updatedProj, nil
}

func (s *service) Delete(ctx context.Context, id string) error {
	if err := uuid.Validate(id); err != nil {
		return ErrInvalidUUID.Wrap(err)
	}
	if err := s.repo.Delete(ctx, id); err != nil {
		return ErrServiceDelete.Wrap(err)
	}
	return nil
}
