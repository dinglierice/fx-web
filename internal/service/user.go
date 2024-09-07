package service

import (
	"context"
	"fx-web/internal/domain"
)

type userService struct {
	repo domain.UserRepository
}

func NewUserService(repo domain.UserRepository) domain.UserService {
	return &userService{repo: repo}
}

func (s *userService) GetUser(ctx context.Context, id string) (*domain.User, error) {
	return s.repo.GetByID(ctx, id)
}

func (s *userService) CreateUser(ctx context.Context, user *domain.User) error {
	// 可以在这里添加业务逻辑,比如验证
	return s.repo.Create(ctx, user)
}

func (s *userService) UpdateUser(ctx context.Context, user *domain.User) error {
	// 可以在这里添加业务逻辑,比如验证
	return s.repo.Update(ctx, user)
}

func (s *userService) DeleteUser(ctx context.Context, id string) error {
	return s.repo.Delete(ctx, id)
}
