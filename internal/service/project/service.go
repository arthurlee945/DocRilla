package project

import (
	"context"

	"github.com/arthurlee945/Docrilla/internal/model"
)

type Service interface {
}

type Project struct {
	model.Project
}

type service struct {
	repository Repository
}

func NewService(r Repository) Service {
	return &service{
		repository: r,
	}
}

func (s *service) GetAll(ctx context.Context, opts struct{ limit string }) (*Project, error) {
	return nil, nil
}

func (s *service) GetOverviewById(ctx context.Context, uuid string) (*Project, error) {
	return nil, nil
}

func (s *service) GetDetailById(ctx context.Context, uuid string) (*Project, error) {
	return nil, nil
}

func (s *service) Update(ctx context.Context, uuid string) error {
	return nil
}

func (s *service) Delete(ctx context.Context, uuid string) error {
	return nil
}
