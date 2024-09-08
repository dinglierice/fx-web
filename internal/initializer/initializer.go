package initializer

import (
	"fx-web/internal/conf"
	"fx-web/internal/ent"
	"go.uber.org/zap"
)

type Initializer struct {
	dbInit    func(*conf.Config) (*ent.Client, error)
	cacheInit func(*conf.Config) error
	logger    *zap.Logger
}

func NewInitializer(dbInit func(*conf.Config) (*ent.Client, error),
	cacheInit func(*conf.Config) error) *Initializer {
	return &Initializer{
		dbInit:    dbInit,
		cacheInit: cacheInit,
	}
}

func (i *Initializer) Initialize(cfg *conf.Config) error {
	i.logger.Info("initializing...")
	if _, err := i.dbInit(cfg); err != nil {
		return err
	}
	i.logger.Info("initializing db successfully")
	if err := i.cacheInit(cfg); err != nil {
		return err
	}
	// 初始化其他子系统...
	i.logger.Info("initializing done")
	return nil
}
