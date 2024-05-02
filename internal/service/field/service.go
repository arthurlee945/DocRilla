package field

import (
	"context"
	"fmt"

	"github.com/arthurlee945/Docrilla/internal/model"
	fieldEnum "github.com/arthurlee945/Docrilla/internal/model/enum/field"
	"github.com/arthurlee945/Docrilla/internal/util"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
)

var validate = validator.New(validator.WithRequiredStructEnabled())

type Service interface {
	GetById(ctx context.Context, id string) (*model.Field, error)
	Create(ctx context.Context, req CreateRequest) (*model.Field, error)
	Update(ctx context.Context, req UpdateRequest) (*model.Field, error)
	Delete(ctx context.Context, id string) error
}

type service struct {
	repo Repository
}

func NewService(repo Repository) Service {
	return &service{repo}
}

func (s *service) GetById(ctx context.Context, id string) (*model.Field, error) {
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
	X         *float64
	Y         *float64
	Width     *float64
	Height    *float64
	Page      *uint32
	Type      *fieldEnum.Type
}

func (s *service) Update(ctx context.Context, req UpdateRequest) (*model.Field, error) {
	if err := validate.Struct(req); err != nil {
		return nil, ErrInvalidReqObj.Wrap(err)
	}
	if fErr, pErr := uuid.Validate(req.UUID), uuid.Validate(req.ProjectID); fErr != nil || pErr != nil {
		return nil, ErrInvalidUUID.Wrap(fmt.Errorf("field id err=%+v; project id err =%+v", fErr, pErr))
	}
	reqField := &model.Field{
		UUID:      util.ToPointer(req.UUID),
		ProjectID: util.ToPointer(req.ProjectID),
		X:         req.X,
		Y:         req.Y,
		Width:     req.Width,
		Height:    req.Height,
		Page:      req.Page,
		Type:      req.Type,
	}
	updatedField, err := s.repo.Update(ctx, reqField)
	if err != nil {
		return nil, ErrServiceUpdate.Wrap(err)
	}

	return updatedField, nil
}

func (s *service) Delete(ctx context.Context, id string) error {
	if err := uuid.Validate(id); err != nil {
		return ErrServiceDelete.Wrap(err)
	}
	if err := s.repo.Delete(ctx, id); err != nil {
		return err
	}
	return nil
}
