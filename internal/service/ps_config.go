package service

import (
	"context"
	"errors"
	"fx-web/internal/domain"
	"fx-web/internal/ent"
	"sync"
)

// psConfigService is a concrete implementation of PsConfigService.
type psConfigService struct {
	repo domain.PsConfigRepository
}

// NewPsConfigService creates a new PsConfigService.
func NewPsConfigService(repo domain.PsConfigRepository) domain.PsConfigService {
	return &psConfigService{repo: repo}
}

func (s *psConfigService) GetConfigByID(ctx context.Context, id uint64) (*domain.PsConfig, error) {
	entConfig, err := s.repo.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}
	return convertEntToDomain(entConfig), nil
}

func (s *psConfigService) GetConfigByScene(ctx context.Context, scene string) (*domain.PsConfig, error) {
	entConfig, err := s.repo.GetByScene(ctx, scene)
	if err != nil {
		return nil, err
	}
	return convertEntToDomain(entConfig), nil
}

func (s *psConfigService) CreateConfig(ctx context.Context, config *domain.PsConfig) error {
	if config.PsScene == "" {
		return errors.New("ps_scene is required")
	}
	entConfig := convertDomainToEnt(config)
	return s.repo.Create(ctx, entConfig)
}

func (s *psConfigService) UpdateConfig(ctx context.Context, config *domain.PsConfig) error {
	if config.PsID == 0 {
		return errors.New("ps_id is required")
	}
	entConfig := convertDomainToEnt(config)
	return s.repo.Update(ctx, entConfig)
}

func (s *psConfigService) DeleteConfig(ctx context.Context, id uint64) error {
	return s.repo.Delete(ctx, id)
}

// ListConfigs lists all configs.
// TODO 学习go并发的核心概念和使用
func (s *psConfigService) ListConfigs(ctx context.Context, limit, offset int) ([]*domain.PsConfig, error) {
	var wg sync.WaitGroup
	var mu sync.Mutex
	var domainConfigs []*domain.PsConfig
	errChan := make(chan error, 1)
	doneChan := make(chan struct{})

	wg.Add(1)
	go func() {
		defer wg.Done()
		entConfigs, err := s.repo.List(ctx, limit, offset)
		if err != nil {
			errChan <- err
			return
		}
		mu.Lock()
		defer mu.Unlock()
		for _, entConfig := range entConfigs {
			domainConfigs = append(domainConfigs, convertEntToDomain(entConfig))
		}
	}()

	go func() {
		wg.Wait()
		close(doneChan)
		close(errChan)
	}()

	select {
	case err := <-errChan:
		if err != nil {
			return nil, err
		}
	case <-doneChan:
		// All goroutines have completed
	}
	return domainConfigs, nil
}

// convertEntToDomain converts an ent.PsConfig to a domain.PsConfig.
func convertEntToDomain(entConfig *ent.PsConfig) *domain.PsConfig {
	return &domain.PsConfig{
		PsID:       entConfig.ID,
		PsScene:    entConfig.PsScene,
		PsFilter:   entConfig.PsFilter,
		PsMessage:  entConfig.PsMessage,
		PsEvent:    entConfig.PsEvent,
		PsFeature:  entConfig.PsFeature,
		PsStrategy: entConfig.PsStrategy,
		OwnerID:    entConfig.OwnerID,
		Managers:   entConfig.Managers,
		UpdateUser: entConfig.UpdateUser,
		CreatedAt:  entConfig.CreatedAt,
		UpdatedAt:  entConfig.UpdatedAt,
	}
}

// convertDomainToEnt converts a domain.PsConfig to an ent.PsConfig.
func convertDomainToEnt(domainConfig *domain.PsConfig) *ent.PsConfig {
	return &ent.PsConfig{
		ID:         domainConfig.PsID,
		PsScene:    domainConfig.PsScene,
		PsFilter:   domainConfig.PsFilter,
		PsMessage:  domainConfig.PsMessage,
		PsEvent:    domainConfig.PsEvent,
		PsFeature:  domainConfig.PsFeature,
		PsStrategy: domainConfig.PsStrategy,
		OwnerID:    domainConfig.OwnerID,
		Managers:   domainConfig.Managers,
		UpdateUser: domainConfig.UpdateUser,
		CreatedAt:  domainConfig.CreatedAt,
		UpdatedAt:  domainConfig.UpdatedAt,
	}
}
