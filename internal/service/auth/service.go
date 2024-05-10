package auth

import (
	"context"

	"github.com/arthurlee945/Docrilla/internal/config"
	"github.com/arthurlee945/Docrilla/internal/errors"
	"github.com/arthurlee945/Docrilla/internal/model"
	"github.com/go-playground/validator/v10"
)

var validate = validator.New(validator.WithRequiredStructEnabled())

type Service interface {
	LogIn(ctx context.Context, req LogInReq) (jwt string, err error)
	SignUp(ctx context.Context, req SignUpRequest) (jwt string, err error)
	Delete(ctx context.Context) error
}

type service struct {
	cfg  *config.Config
	repo Repository
}

func NewService(cfg *config.Config, repository Repository) Service {
	return &service{cfg, repository}
}

type LogInReq struct {
	Email    string `json:"email" validate:"required"`
	Password string `json:"password" validate:"required"`
}

func (s *service) LogIn(ctx context.Context, req LogInReq) (jwt string, err error) {
	if err := validate.Struct(req); err != nil {
		return "", errors.ErrInvalidRequest.Wrap(err)
	}
	user, err := s.repo.GetByEmail(ctx, req.Email)
	if err != nil {
		return "", errors.AuthError(err)
	}

	if !compareHash(req.Password, *user.Password) {
		return "", errors.ErrUnauthorized
	}
	return generateToken(s.cfg.JwtSecret, *user.ID)
}

type SignUpRequest struct {
	Name     string
	Email    string `json:"email" validate:"required"`
	Password string `json:"password" validate:"required"`
}

func (s *service) SignUp(ctx context.Context, req SignUpRequest) (jwt string, err error) {
	if err := validate.Struct(req); err != nil {
		return "", errors.ErrInvalidRequest.Wrap(err)
	}
	hashedPassword, err := generateHash(req.Password)
	if err != nil {
		return "", errors.ErrUnknown.Wrap(err)
	}
	user, err := s.repo.Create(ctx, &model.User{
		Name:     &req.Name,
		Email:    &req.Email,
		Password: &hashedPassword,
	})
	if err != nil {
		return "", errors.AuthError(err)
	}
	return generateToken(s.cfg.JwtSecret, *user.ID)
}

func (s *service) Delete(ctx context.Context) error {
	userId, err := GetUser(ctx)
	if err != nil {
		return err
	}
	if err := s.repo.Delete(ctx, &model.User{ID: &userId}); err != nil {
		return errors.AuthError(err)
	}
	return nil
}
