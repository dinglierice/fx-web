package handler

import (
	"fx-web/internal/domain"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
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

// UserQueryTest TODO 这种Set写法和直接c.JSON的区别是什么
func (u *UserHandler) UserQueryTest(c *gin.Context) {
	idString := c.Param("id")
	pId, _ := strconv.ParseUint(idString, 10, 64)
	user, err := u.service.GetUser(c, pId)
	if err != nil {
		c.Set("errCode", 400)
		c.Set("errMessage", err.Error())
		c.Status(http.StatusBadRequest)
	} else {
		c.Set("data", user)
		c.Status(http.StatusOK)
	}
}
