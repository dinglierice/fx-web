package apis

import (
	"fx-web/internal/routes/handler"
	"github.com/gin-gonic/gin"
)

func RegisterHealthRoutes(r *gin.Engine) {
	r.GET("/health", handler.HealthCheck)
}
