package routes

import (
	"fx-web/internal/routes/handler"
	"github.com/gin-gonic/gin"
)

// 接口:Routes, 提供设置路由的SetupRoutes方法
// 实现:GinRoutes, 提供SetupRoutes方法的具体实现
// 注入:handler.UserHandler
// ProvideRoutes: 提供注入构造的方法

type Routes interface {
	SetupRoutes(r *gin.Engine)
}

type GinRoutes struct {
	userHandler   *handler.UserHandler
	configHandler *handler.ConfigHandler
}

func (ginRoutes *GinRoutes) SetupRoutes(r *gin.Engine) {
	// 健康检查路由
	r.GET("/health", handler.HealthCheck)
	// 用户路由
	r.POST("/user/login", ginRoutes.userHandler.UserLogin)
	r.GET("/user/queryTest/:id", ginRoutes.userHandler.UserQueryTest)

	// 策略路由
	r.GET("/configs/list", ginRoutes.configHandler.List)
}

func ProvideRoutes(userHandler *handler.UserHandler, configHandler *handler.ConfigHandler) Routes {
	return &GinRoutes{
		userHandler:   userHandler,
		configHandler: configHandler,
	}
}
