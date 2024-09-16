package repository

import (
	"context"
	"fx-web/internal/domain"
	"fx-web/internal/ent"
	"fx-web/internal/ent/user"
)

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
	_, err := r.db.User.Create().
		SetUserName(user.UserName).
		SetPasswordDigest(user.PasswordDigest).
		SetNickName(user.NickName).
		SetMoney(user.Money).
		SetStatus(user.Status).
		SetAvatar(user.Avatar).
		SetEmail(user.Email).
		Save(ctx)
	if err != nil {
		return err
	}
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

func (r *userRepository) GetByName(ctx context.Context, name string) (*ent.User, error) {
	u, err := r.db.User.
		Query().
		Where(user.UserNameEQ(name)).
		Only(ctx)
	if err != nil {
		if ent.IsNotFound(err) {
			return nil, nil
		}
		return nil, err
	}
	return u, nil
}
