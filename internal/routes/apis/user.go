package apis

import (
	"fx-web/internal/routes/handler"
	"github.com/gin-gonic/gin"
)

func RegisterUserRoutes(r *gin.Engine) {
	r.POST("/user/login", handler.UserLogin)
}
