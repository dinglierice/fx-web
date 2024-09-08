package initializer

import (
	"context"
	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	"fmt"
	"fx-web/internal/conf"
	"fx-web/internal/ent"
	"go.uber.org/zap"
	"time"
)

// Init TODO 这里还要记得开启ent的debug模式
func Init(logger *zap.Logger) func(*conf.Config) (*ent.Client, error) {
	return func(cfg *conf.Config) (*ent.Client, error) {
		logger.Info("Initializing database connection")

		driverName := cfg.DB.Driver
		dataSourceName := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
			cfg.DB.Host, cfg.DB.Port, cfg.DB.User, cfg.DB.Password, cfg.DB.Name)

		drv, err := sql.Open(driverName, dataSourceName)
		if err != nil {
			logger.Error("Failed to open database connection", zap.Error(err))
			return nil, fmt.Errorf("failed to open database connection: %w", err)
		}

		db := drv.DB()
		db.SetMaxIdleConns(cfg.DB.MaxIdleConns)
		db.SetMaxOpenConns(cfg.DB.MaxOpenConns)
		db.SetConnMaxLifetime(time.Duration(cfg.DB.ConnMaxLifetime) * time.Second)

		client := ent.NewClient(ent.Driver(drv))

		if cfg.DB.AutoMigrate {
			if err := client.Schema.Create(context.Background()); err != nil {
				logger.Error("Failed to create schema resources", zap.Error(err))
				return nil, fmt.Errorf("failed to create schema resources: %w", err)
			}
		}

		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		if err := client.Debug().Schema.Diff(ctx, dialect.MySQL); err != nil {
			logger.Error("Failed to verify database schema", zap.Error(err))
			return nil, fmt.Errorf("failed to verify database schema: %w", err)
		}

		logger.Info("Database connection and schema verification completed successfully")

		return client, nil
	}
}
