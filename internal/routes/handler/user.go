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

func NewUserHandler(service domain.UserService, logger *zap.Logger) *UserHandler {
	return &UserHandler{
		service: service,
		logger:  logger,
	}
}

func (u *UserHandler) UserRegister(c *gin.Context) {
	user := &domain.User{}
	if err := c.ShouldBind(user); err == nil {
		_ = u.service.CreateUser(c.Request.Context(), user)
		c.Status(200)
		c.Set("data", "注册成功")
	} else {
		c.Set("errCode", 400)
		c.Set("errMessage", err.Error())
		c.Status(http.StatusBadRequest)
	}
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

// @BasePath /api/v1

// UserQueryTest @Summary 用于测试获取用户信息
// @Description 获取用户信息接口
// @Tags user
// @Accept json
// @Produce json
// @Param id path uint64 true "用户ID"
// @Success 200 {object} domain.User "用户信息"
// @Failure 400 {object} any "请求错误"
// @Router /user/queryTest/{id} [get]
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
