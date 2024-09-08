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

func (s *userService) GetUser(ctx context.Context, id int) (*domain.User, error) {
	byID, err := s.repo.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}
	return &domain.User{
		Base: domain.Base{
			ID:        byID.ID,
			CreatedAt: byID.CreatedAt,
			UpdatedAt: byID.UpdatedAt,
		},
		UserName: byID.UserName,
		Email:    byID.Email,
	}, nil
}

func (s *userService) CreateUser(ctx context.Context, user *domain.User) error {
	// 可以在这里添加业务逻辑,比如验证
	return nil
}

func (s *userService) UpdateUser(ctx context.Context, user *domain.User) error {
	// 可以在这里添加业务逻辑,比如验证
	return nil
}

func (s *userService) DeleteUser(ctx context.Context, id string) error {
	return nil
}
