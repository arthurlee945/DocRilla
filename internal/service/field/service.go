package field

import (
	"context"

	"github.com/arthurlee945/Docrilla/internal/model"
	fieldEnum "github.com/arthurlee945/Docrilla/internal/model/enum/field"
	"github.com/arthurlee945/Docrilla/internal/util"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
)

var validate = validator.New(validator.WithRequiredStructEnabled())

type Service interface {
	Get(ctx context.Context, id string) (*model.Field, error)
}

type service struct {
	repo Repository
}

func NewService(repo Repository) Service {
	return &service{repo}
}

func (s *service) Get(ctx context.Context, id string) (*model.Field, error) {
	if err := uuid.Validate(id); err != nil {
		return nil, ErrInvalidUUID.Wrap(err)
	}
	field, err := s.repo.GetById(ctx, id)
	if err != nil {
		return nil, ErrServiceGet.Wrap(err)
	}
	return field, nil
}

type CreateRequest struct {
	ProjectId string `validate:"required"`
	X         float64
	Y         float64
	Width     float64        `validate:"required"`
	Height    float64        `validate:"required"`
	Page      uint32         `validate:"required"`
	Type      fieldEnum.Type `validate:"required"`
}

func (s *service) Create(ctx context.Context, req CreateRequest) (*model.Field, error) {
	if err := validate.Struct(req); err != nil {
		return nil, ErrInvalidReqObj.Wrap(err)
	}
	if err := uuid.Validate(req.ProjectId); err != nil {
		return nil, ErrInvalidUUID.Wrap(err)
	}
	field, err := s.repo.Create(ctx, &model.Field{
		ProjectID: util.ToPointer(req.ProjectId),
		X:         util.ToPointer(req.X),
		Y:         util.ToPointer(req.Y),
		Width:     util.ToPointer(req.Width),
		Height:    util.ToPointer(req.Height),
		Page:      util.ToPointer(req.Page),
		Type:      util.ToPointer(req.Type),
	})
	if err != nil {
		return nil, ErrServiceCreate.Wrap(err)
	}
	return field, nil
}

type UpdateRequest struct {
	UUID      string `validate:"required"`
	ProjectID string `validate:"required"`
}

func (s *service) Update(ctx context.Context, req UpdateRequest) (*model.Field, error) {
	return nil, nil
}
