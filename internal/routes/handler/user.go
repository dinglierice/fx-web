package handler

import (
	"fx-web/internal/domain"
	"github.com/gin-gonic/gin"
	"strconv"
)

type UserHandler struct {
	service domain.UserService
}

func NewUserHandler(service domain.UserService) *UserHandler {
	return &UserHandler{service: service}
}

func (u *UserHandler) UserLogin(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "login",
	})
}

func (u *UserHandler) UserQueryTest(c *gin.Context) {
	idString := c.Param("id")
	pId, _ := strconv.ParseUint(idString, 10, 64)
	user, err := u.service.GetUser(c, pId)
	if err != nil {
		return
	}
	c.JSON(200, gin.H{
		"message": "query",
		"data":    user,
	})
}
