package repository

import (
	"context"
	"fx-web/internal/domain"
	"fx-web/internal/ent"
)

// TODO 待实现
type userRepository struct {
	db *ent.Client
}

func NewUserRepository(db *ent.Client) domain.UserRepository {
	return &userRepository{db: db}
}

func (r *userRepository) GetByID(ctx context.Context, id uint64) (*ent.User, error) {
	// 实现从数据库获取用户的逻辑
	return r.db.User.Get(ctx, id)
}

func (r *userRepository) Create(ctx context.Context, user *ent.User) error {
	// 实现创建用户的逻辑
	return nil
}

func (r *userRepository) Update(ctx context.Context, user *ent.User) error {
	// 实现更新用户的逻辑
	return nil
}

func (r *userRepository) Delete(ctx context.Context, id string) error {
	// 实现删除用户的逻辑
	return nil
}
