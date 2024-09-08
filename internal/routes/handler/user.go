package handler

import (
	"fx-web/internal/domain"
	"github.com/gin-gonic/gin"
)

func UserLogin(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "login",
	})
}

func UserQueryTest(c *gin.Context, userId int, service domain.UserService) {
	user, err := service.GetUser(c, userId)
	if err != nil {
		return
	}
	c.JSON(200, gin.H{
		"message": "query",
		"data":    user,
	})
}
