package main

import (
	"database/sql"
	"fx-web/internal/app"
	"fx-web/internal/conf"
	"fx-web/internal/initializer"
	"fx-web/internal/logger"
	"fx-web/internal/repository"
	"fx-web/internal/server"
	"fx-web/internal/service"
	"go.uber.org/fx"
)

func main() {
	fx.New(
		fx.Provide(
			// 注入gin引擎
			server.ProvideGinEngine,
			conf.ProvideDevConfig,
			logger.ProvideLogger,
			repository.NewUserRepository,
			service.NewUserService,

			provideDatabaseConnection,

			initializer.NewInitializer,
		),
		fx.Invoke(
			app.StartApp,
		)).Run()
}

func provideDatabaseConnection(cfg *conf.Config) (*sql.DB, error) {
	// 实现数据库连接逻辑
	return nil, nil
}
