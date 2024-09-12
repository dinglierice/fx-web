package repository

import (
	"context"
	"errors"
	"fx-web/internal/domain"
	"fx-web/internal/ent"
	"fx-web/internal/ent/psconfig"
)

// TODO 在psConfig场景中去理解值传递和引用传递
// entPsConfigRepository is a concrete implementation of PsConfigRepository using Ent.
type entPsConfigRepository struct {
	client *ent.Client
}

// NewPsConfigRepository creates a new entPsConfigRepository.
func NewPsConfigRepository(client *ent.Client) domain.PsConfigRepository {
	return &entPsConfigRepository{client: client}
}

func (r *entPsConfigRepository) GetByID(ctx context.Context, id uint64) (*ent.PsConfig, error) {
	config, err := r.client.PsConfig.Get(ctx, id)
	if err != nil {
		return nil, err
	}
	return config, nil
}

func (r *entPsConfigRepository) GetByScene(ctx context.Context, scene string) (*ent.PsConfig, error) {
	config, err := r.client.PsConfig.Query().Where(psconfig.PsScene(scene)).Only(ctx)
	if err != nil {
		return nil, err
	}
	return config, nil
}

func (r *entPsConfigRepository) Create(ctx context.Context, config *ent.PsConfig) error {
	_, err := r.client.PsConfig.Create().
		SetPsScene(config.PsScene).
		SetPsFilter(config.PsFilter).
		SetNillablePsMessage(config.PsMessage).
		SetNillablePsEvent(config.PsEvent).
		SetNillablePsFeature(config.PsFeature).
		SetNillablePsStrategy(config.PsStrategy).
		SetNillableOwnerID(config.OwnerID).
		SetNillableManagers(config.Managers).
		SetNillableUpdateUser(config.UpdateUser).
		SetCreatedAt(config.CreatedAt).
		SetCreatedAt(config.UpdatedAt).
		Save(ctx)
	return err
}

func (r *entPsConfigRepository) Update(ctx context.Context, config *ent.PsConfig) error {
	if config.ID == 0 {
		return errors.New("ps_id is required")
	}
	_, err := r.client.PsConfig.UpdateOneID(config.ID).
		SetPsScene(config.PsScene).
		SetPsFilter(config.PsFilter).
		SetNillablePsMessage(config.PsMessage).
		SetNillablePsEvent(config.PsEvent).
		SetNillablePsFeature(config.PsFeature).
		SetNillablePsStrategy(config.PsStrategy).
		SetNillableOwnerID(config.OwnerID).
		SetNillableManagers(config.Managers).
		SetNillableUpdateUser(config.UpdateUser).
		SetUpdatedAt(config.UpdatedAt).
		Save(ctx)
	return err
}

func (r *entPsConfigRepository) Delete(ctx context.Context, id uint64) error {
	return r.client.PsConfig.DeleteOneID(id).Exec(ctx)
}

func (r *entPsConfigRepository) List(ctx context.Context, limit, offset int) ([]*ent.PsConfig, error) {
	configs, err := r.client.PsConfig.Query().
		Limit(limit).
		Offset(offset).
		All(ctx)
	if err != nil {
		return nil, err
	}

	var result []*ent.PsConfig
	for _, config := range configs {
		result = append(result, config)
	}
	return result, nil
}
