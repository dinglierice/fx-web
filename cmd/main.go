package main

import (
	"fx-web/internal/app"
	"fx-web/internal/conf"
	"fx-web/internal/db"
	"fx-web/internal/logger"
	"fx-web/internal/repository"
	"fx-web/internal/routes"
	"fx-web/internal/routes/handler"
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
			db.InitEntDbConnection,
			// 用户MVC
			repository.NewUserRepository,
			service.NewUserService,
			handler.NewUserHandler,
			// 配置MVN
			repository.NewPsConfigRepository,
			service.NewPsConfigService,
			handler.NewPsConfigHandler,
			// 注入路由
			routes.ProvideRoutes,
		),
		fx.Invoke(
			app.StartApp,
		)).Run()
}
