package handler

import (
	"fx-web/internal/domain"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"strconv"
)

type UserHandler struct {
	service domain.UserService
	logger  *zap.Logger
}

func NewUserHandler(service domain.UserService) *UserHandler {
	return &UserHandler{service: service}
}

func (u *UserHandler) UserLogin(c *gin.Context) {
	user := &domain.User{}
	if err := c.ShouldBind(user); err == nil {
		res := u.service.Login(c.Request.Context(), user)
		c.JSON(200, res)
	} else {
		u.logger.Info("UserLogin", zap.Error(err))
		c.JSON(400, nil)
	}
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
