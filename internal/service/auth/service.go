package auth

import (
	"context"

	"github.com/arthurlee945/Docrilla/internal/model"
)

type Service interface{}

type service struct {
	repo Repository
}

func NewService(repository Repository) Service {
	return &service{repository}
}

func (s *service) LogIn(ctx context.Context, email, password string) (string, error) {

	return "", nil
}

type SignUpRequest struct {
}

func (s *service) SignUp(ctx context.Context, req SignUpRequest) (*model.User, error) {
	return nil, nil
}

func (s *service) Delete(ctx context.Context, token string) error {
	return nil
}
