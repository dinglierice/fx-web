package db

import (
	"context"
	"entgo.io/ent/dialect/sql"
	"fmt"
	"fx-web/internal/conf"
	"fx-web/internal/ent"
	"go.uber.org/zap"
	"time"
)

func InitEntDbConnection(logger *zap.Logger) func(*conf.Config) (*ent.Client, error) {
	return func(cfg *conf.Config) (*ent.Client, error) {
		logger.Info("Initializing database connection")

		driverName := cfg.DB.Driver
		dataSourceName := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?parseTime=True",
			cfg.DB.User, cfg.DB.Password, cfg.DB.Host, cfg.DB.Port, cfg.DB.Name)

		drv, err := sql.Open(driverName, dataSourceName)
		if err != nil {
			logger.Error("Failed to open database connection", zap.Error(err))
			return nil, fmt.Errorf("failed to open database connection: %w", err)
		}

		db := drv.DB()
		db.SetMaxIdleConns(cfg.DB.MaxIdleConns)
		db.SetMaxOpenConns(cfg.DB.MaxOpenConns)
		db.SetConnMaxLifetime(time.Duration(cfg.DB.ConnMaxLifetime) * time.Second)

		var client *ent.Client
		if cfg.Environment == "development" {
			client = ent.NewClient(ent.Driver(drv)).Debug()
		}
		client = ent.NewClient(ent.Driver(drv))

		if cfg.DB.AutoMigrate {
			if err = client.Schema.Create(context.Background()); err != nil {
				logger.Error("Failed to create schema resources", zap.Error(err))
				return nil, fmt.Errorf("failed to create schema resources: %w", err)
			}
		}

		logger.Info("Database connection and schema verification completed successfully")

		return client, nil
	}
}
