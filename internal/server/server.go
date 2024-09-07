package server

import (
	"fx-web/internal/conf"
	"github.com/gin-gonic/gin"
)

func ProvideGinEngine(cfg *conf.Config) *gin.Engine {
	if cfg.Environment == "production" {
		gin.SetMode(gin.ReleaseMode)
	}

	r := gin.Default()
	r.Use(gin.Recovery())
	setupRoutes(r)
	return r
}

func setupRoutes(r *gin.Engine) {
	// 定义路由
	r.GET("/health", healthCheck)
}

func healthCheck(c *gin.Context) {
	c.JSON(200, gin.H{
		"status": "ok",
	})
}
