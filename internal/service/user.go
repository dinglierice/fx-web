package service

import (
	"context"
	"fx-web/internal/domain"
	"go.uber.org/zap"
)

type userService struct {
	repo   domain.UserRepository
	logger *zap.Logger
}

func NewUserService(repo domain.UserRepository, logger *zap.Logger) domain.UserService {
	return &userService{repo: repo, logger: logger}
}

func (s *userService) GetUser(ctx context.Context, id uint64) (*domain.User, error) {
	byID, err := s.repo.GetByID(ctx, id)
	if err != nil {
		s.logger.Info("GetUser repo execute error", zap.Error(err))
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

func (s *userService) Login(ctx context.Context, user *domain.User) error {
	// TODO 看似无需新增repo接口, 复现下参数校验逻辑即可
	return nil
}
