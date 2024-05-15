package project

import (
	"context"
	"sync"

	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"

	"github.com/arthurlee945/Docrilla/internal/errors"
	"github.com/arthurlee945/Docrilla/internal/model"
	"github.com/arthurlee945/Docrilla/internal/service/auth"
	"github.com/arthurlee945/Docrilla/internal/service/field"
	"github.com/arthurlee945/Docrilla/internal/util/ptr"
)

var validate = validator.New(validator.WithRequiredStructEnabled())

type Service interface {
	GetAll(context.Context, GetAllRequest) (projects []model.Project, nextCursor string, err error)
	GetOverviewById(context.Context, string) (*model.Project, error)
	GetDetailById(context.Context, string) (*model.Project, error)
	Create(context.Context, CreateRequest) (*model.Project, error)
	Update(context.Context, UpdateRequest) (*model.Project, error)
	Delete(context.Context, string) error
}

type service struct {
	projRepository  Repository
	fieldRepository field.Repository
}

func NewService(projRepository Repository, fieldReposity field.Repository) Service {
	return &service{projRepository, fieldReposity}
}

type GetAllRequest struct {
	Limit  uint8
	Cursor string
}

func (s *service) GetAll(ctx context.Context, req GetAllRequest) ([]model.Project, string, error) {
	userId, err := auth.GetUser(ctx)
	if err != nil {
		return nil, "", err
	}
	projects, nextCursor, err := s.projRepository.GetAll(ctx, req.Limit, req.Cursor, userId)
	if err != nil {
		return nil, "", err
	}
	return projects, nextCursor, nil
}

func (s *service) GetOverviewById(ctx context.Context, id string) (*model.Project, error) {
	userId, err := auth.GetUser(ctx)
	if err != nil {
		return nil, err
	}
	if err := uuid.Validate(id); err != nil {
		return nil, errors.ErrInvalidRequest.Wrap(err)
	}
	project, err := s.projRepository.GetOverviewById(ctx, id, userId)
	if err != nil {
		return nil, err
	}
	return project, nil
}

func (s *service) GetDetailById(ctx context.Context, id string) (*model.Project, error) {
	userId, err := auth.GetUser(ctx)

	if err != nil {
		return nil, err
	}
	if err := uuid.Validate(id); err != nil {
		return nil, errors.ErrInvalidRequest.Wrap(err)
	}
	project, err := s.projRepository.GetDetailById(ctx, id, userId)
	if err != nil {
		return nil, err
	}
	return project, nil
}

type CreateRequest struct {
	Title       string  `json:"title" validate:"required"`
	Description *string `json:"desciption"`
	Route       *string `json:"route"`
	Token       *string `json:"token"`
	DocumentUrl string  `json:"documentUrl" validate:"required"`
}

func (s *service) Create(ctx context.Context, req CreateRequest) (*model.Project, error) {
	userId, err := auth.GetUser(ctx)
	if err != nil {
		return nil, err
	}
	if err := validate.Struct(req); err != nil {
		return nil, errors.ErrValidation.Wrap(err)
	}

	if req.Route == nil {
		req.Route = ptr.ToPointer(uuid.NewString())
	}

	createdProj, err := s.projRepository.Create(ctx, &model.Project{
		UserID:      &userId,
		Title:       ptr.ToPointer(req.Title),
		Description: req.Description,
		Route:       req.Route,
		Token:       req.Token,
		DocumentUrl: ptr.ToPointer(req.DocumentUrl),
	})
	if err != nil {
		return nil, err
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
	Fields      []field.UpdateRequest
}

// ADD Field repo and update this
func (s *service) Update(ctx context.Context, req UpdateRequest) (*model.Project, error) {
	userId, err := auth.GetUser(ctx)
	if err != nil {
		return nil, err
	}
	if err := validate.Struct(req); err != nil {
		return nil, errors.ErrValidation.Wrap(err)
	}
	if err := uuid.Validate(req.UUID); err != nil {
		return nil, errors.ErrInvalidRequest.Wrap(err)
	}
	projChan, errChan := make(chan *model.Project), make(chan error)
	uCtx, uCancel := context.WithCancel(ctx)
	defer uCancel()

	go func() {
		wg := sync.WaitGroup{}
		wg.Add(len(req.Fields) + 1)
		var uProj *model.Project
		uFields := []model.Field{}

		go func() {
			defer wg.Done()
			updatedProj, err := s.projRepository.Update(uCtx, &model.Project{
				UserID:      &userId,
				UUID:        ptr.ToPointer(req.UUID),
				Title:       req.Title,
				Description: req.Description,
				DocumentUrl: req.DocumentUrl,
				Route:       req.Route,
				Token:       req.Token,
			})
			if err != nil {
				errChan <- err
				uCancel()
			}
			uProj = updatedProj
		}()

		for _, f := range req.Fields {
			go func(fieldUpReq *field.UpdateRequest) {
				defer wg.Done()
				uField, err := s.fieldRepository.Update(uCtx, &model.Field{
					UUID:      &fieldUpReq.UUID,
					ProjectID: &req.UUID,
					X:         fieldUpReq.X,
					Y:         fieldUpReq.Y,
					Width:     fieldUpReq.Width,
					Height:    fieldUpReq.Height,
					Page:      fieldUpReq.Page,
					Type:      fieldUpReq.Type,
				})
				if err != nil {
					errChan <- err
					uCancel()
				}
				uFields = append(uFields, *uField)
			}(&f)
		}

		wg.Wait()
		uProj.Fields = uFields
		projChan <- uProj
	}()
	select {
	case err := <-errChan:
		return nil, err
	case proj := <-projChan:
		return proj, nil
	}
}

func (s *service) Delete(ctx context.Context, id string) error {
	userId, err := auth.GetUser(ctx)
	if err != nil {
		return err
	}
	if err := uuid.Validate(id); err != nil {
		return errors.ErrInvalidRequest.Wrap(err)
	}
	if err := s.projRepository.Delete(ctx, id, userId); err != nil {
		return err
	}
	return nil
}
