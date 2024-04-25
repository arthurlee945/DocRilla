package project

import (
	"context"

	"github.com/go-playground/validator/v10"
	UUID "github.com/google/uuid"

	"github.com/arthurlee945/Docrilla/internal/model"
	"github.com/arthurlee945/Docrilla/internal/util"
)

var validate = validator.New(validator.WithRequiredStructEnabled())

type Service interface {
	GetAll(context.Context, GetAllRequest) ([]model.Project, string, error)
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
	limit  uint8
	cursor string
}

func (s *service) GetAll(ctx context.Context, req GetAllRequest) ([]model.Project, string, error) {
	projects, nextCursor, err := s.repo.GetAll(ctx, req.cursor, req.limit)
	if err != nil {
		return nil, "", ErrRepoGet.Wrap(err)
	}
	return projects, nextCursor, nil
}

func (s *service) GetOverviewById(ctx context.Context, uuid string) (*model.Project, error) {
	if err := UUID.Validate(uuid); err != nil {
		return nil, ErrInvalidUUID.Wrap(err)
	}
	project, err := s.repo.GetOverviewById(ctx, uuid)
	if err != nil {
		return nil, ErrRepoGet.Wrap(err)
	}
	return project, nil
}

func (s *service) GetDetailById(ctx context.Context, uuid string) (*model.Project, error) {
	if err := UUID.Validate(uuid); err != nil {
		return nil, ErrInvalidUUID.Wrap(err)
	}
	project, err := s.repo.GetDetailById(ctx, uuid)
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
	newProj := &model.Project{
		UserID:      util.ToPointer(req.UserID),
		Title:       util.ToPointer(req.Title),
		Description: req.Description,
		Route:       req.Route,
		Token:       req.Token,
		DocumentUrl: util.ToPointer(req.DocumentUrl),
	}
	createdProj, err := s.repo.Create(ctx, newProj)
	if err != nil {
		return nil, ErrServiceCreate.Wrap(err)
	}
	return createdProj, nil
}

// maybe need to seperate field update
type UpdateRequest struct {
	UUID         string `validate:"required"`
	Title        *string
	Description  *string
	Document_url *string
}

func (s *service) Update(ctx context.Context, uuid string) error {
	return nil
}

func (s *service) Delete(ctx context.Context, uuid string) error {
	return nil
}
