package service

import (
	"context"
	"errors"
	"fx-web/internal/domain"
	"fx-web/internal/ent"
	"go.uber.org/zap"
	"golang.org/x/crypto/bcrypt"
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
	if user.UserName == "" || user.PasswordDigest == "" {
		return errors.New("missing required fields")
	}

	// 用户是否已经存在
	existedUser, err := s.repo.GetByName(ctx, user.UserName)
	if err != nil {
		s.logger.Info("GetUser repo execute error", zap.Error(err))
		return err
	}
	if existedUser != nil {
		return errors.New("user already exists")
	}

	// 密码加密
	password, err := genPassword(user.PasswordDigest)
	if err != nil {
		return err
	}

	// 将用户保存到数据库
	userDo := &ent.User{
		NickName:       user.NickName,
		UserName:       user.UserName,
		Status:         domain.Active,
		PasswordDigest: password,
		Money:          user.Money, // 初始金额
	}
	err = s.repo.Create(ctx, userDo)
	if err != nil {
		return err
	}

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

func genPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), domain.PassWordCost)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}
