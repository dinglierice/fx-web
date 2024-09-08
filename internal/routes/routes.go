package routes

import (
	"fx-web/internal/routes/apis"
	"github.com/gin-gonic/gin"
)

type Routes interface {
	SetupRoutes(r *gin.Engine)
}

type GinRoutes struct {
}

func (ginRoutes *GinRoutes) SetupRoutes(r *gin.Engine) {
	// 健康检查路由
	apis.RegisterHealthRoutes(r)
	// 用户路由
	apis.RegisterUserRoutes(r)
}

func ProvideRoutes() Routes {
	return &GinRoutes{}
}
