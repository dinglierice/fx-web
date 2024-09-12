package domain

import (
	"context"
	"fx-web/internal/ent"
	"time"
)

type PsConfig struct {
	PsID       uint64
	PsScene    string
	PsFilter   uint64
	PsMessage  *uint64
	PsEvent    *uint64
	PsFeature  *uint64
	PsStrategy *uint64
	OwnerID    *uint64
	Managers   *string
	UpdateUser *uint64
	CreatedAt  time.Time
	UpdatedAt  time.Time
}

type PsConfigRepository interface {
	GetByID(ctx context.Context, id uint64) (*ent.PsConfig, error)
	GetByScene(ctx context.Context, scene string) (*ent.PsConfig, error)
	Create(ctx context.Context, config *ent.PsConfig) error
	Update(ctx context.Context, config *ent.PsConfig) error
	Delete(ctx context.Context, id uint64) error
	List(ctx context.Context, limit, offset int) ([]*ent.PsConfig, error)
	// 可以根据需求添加其他方法
}

type PsConfigService interface {
	GetConfigByID(ctx context.Context, id uint64) (*PsConfig, error)
	GetConfigByScene(ctx context.Context, scene string) (*PsConfig, error)
	CreateConfig(ctx context.Context, config *PsConfig) error
	UpdateConfig(ctx context.Context, config *PsConfig) error
	DeleteConfig(ctx context.Context, id uint64) error
	ListConfigs(ctx context.Context, limit, offset int) ([]*PsConfig, error)
}
