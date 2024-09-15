package domain

import (
	"context"
	"fx-web/internal/ent"
)

type User struct {
	Base
	UserName       string
	Email          string
	PasswordDigest string
	NickName       string
	Status         string
	Avatar         string
	Money          string
}

type UserRepository interface {
	GetByID(ctx context.Context, id uint64) (*ent.User, error)
	Create(ctx context.Context, user *ent.User) error
	Update(ctx context.Context, user *ent.User) error
	Delete(ctx context.Context, id string) error
	// 可以根据需求添加其他方法
}

type UserService interface {
	GetUser(ctx context.Context, id uint64) (*User, error)
	CreateUser(ctx context.Context, user *User) error
	UpdateUser(ctx context.Context, user *User) error
	DeleteUser(ctx context.Context, id string) error
	Login(ctx context.Context, user *User) error
}

type UserDTO struct {
	NickName string `form:"nick_name" json:"nick_name"`
	UserName string `form:"user_name" json:"user_name"`
	Password string `form:"password" json:"password"`
	Key      string `form:"key" json:"key"`
}

func (dto *UserDTO) ToUser(passwordDigest string) *User {
	return &User{
		UserName:       dto.UserName,
		Email:          "", // Email should be set separately if needed
		PasswordDigest: passwordDigest,
		NickName:       dto.NickName,
		Status:         "", // Status should be set separately if needed
		Avatar:         "", // Avatar should be set separately if needed
		Money:          "", // Money should be set separately if needed
	}
}
