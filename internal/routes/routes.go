package routes

import (
	docs "fx-web/docs"
	"fx-web/internal/routes/handler"
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
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
	v1 := r.Group("/api/v1")
	{
		// 健康检查路由
		v1.GET("health", handler.HealthCheck)
		// 用户路由
		v1.POST("user/register", ginRoutes.userHandler.UserRegister)
		v1.POST("user/login", ginRoutes.userHandler.UserLogin)
		v1.GET("user/queryTest/:id", ginRoutes.userHandler.UserQueryTest)

		// 策略路由
		v1.GET("configs/list", ginRoutes.configHandler.List)
	}

	docs.SwaggerInfo.BasePath = "/api/v1"
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
}

func ProvideRoutes(userHandler *handler.UserHandler, configHandler *handler.ConfigHandler) Routes {
	return &GinRoutes{
		userHandler:   userHandler,
		configHandler: configHandler,
	}
}
