package routes

import (
	"fx-web/internal/routes/handler"
	"github.com/gin-gonic/gin"
)

type Routes interface {
	SetupRoutes(r *gin.Engine)
}

type GinRoutes struct {
	userHandler *handler.UserHandler
}

func (ginRoutes *GinRoutes) SetupRoutes(r *gin.Engine) {
	// 健康检查路由
	r.GET("/health", handler.HealthCheck)
	// 用户路由
	r.POST("/user/login", ginRoutes.userHandler.UserLogin)
	r.GET("/user/queryTest/:id", ginRoutes.userHandler.UserQueryTest)
}

func ProvideRoutes(userHandler *handler.UserHandler) Routes {
	return &GinRoutes{
		userHandler: userHandler,
	}
}
