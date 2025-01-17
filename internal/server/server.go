package server

import (
	"fx-web/internal/conf"
	"fx-web/internal/middleware"
	"fx-web/internal/routes"
	"github.com/gin-gonic/gin"
)

func ProvideGinEngine(cfg *conf.Config, routes routes.Routes) *gin.Engine {
	if cfg.Environment == "production" {
		gin.SetMode(gin.ReleaseMode)
	}

	r := gin.Default()
	r.Use(gin.Recovery())
	r.Use(middleware.ResponseMiddleware())
	routes.SetupRoutes(r)
	return r
}
