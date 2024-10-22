package handler

import (
	"fx-web/internal/domain"
	"fx-web/internal/utils"
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

// TODO ASE加密算法这里可以再学习下
// TODO 用户注册这部分,起初为了维护service涉及实体的干净, 把过多的逻辑放在了handler中, 待探索和修改, 目前阶段暂时按照service兼职实体转换的方式去做

// @BasePath /api/v1

// UserRegister godoc
// @Summary Register a new user
// @Description Register a new user with the provided details
// @Tags 用户
// @Accept json
// @Produce json
// @Param userDto body domain.UserDTO true "User DTO"
// @Success 200 {object} middleware.CommonResponse{data=string} "success"
// @Failure 400 {object} middleware.Response
// @Failure 500 {object} middleware.Response
// @Router /user/register [post]
func (u *UserHandler) UserRegister(c *gin.Context) {
	userDto := &domain.UserDTO{}
	if err := c.ShouldBind(userDto); err != nil {
		u.logger.Error("Failed to bind user DTO", zap.Error(err))
		c.Status(http.StatusBadRequest)
		c.Set("errCode", 400)
		c.Set("errMessage", err.Error())
		return
	}

	if userDto.Key == "" || len(userDto.Key) != 16 {
		u.logger.Error("密钥长度不够")
		c.Status(http.StatusBadRequest)
		c.Set("errCode", 400)
		c.Set("errMessage", "密钥长度不够")
		return
	}

	money, err := utils.AesEncoding("10000", []byte(userDto.Key))
	if err != nil {
		u.logger.Error("Failed to encrypt initial money", zap.Error(err))
		c.Status(http.StatusBadRequest)
		c.Set("errCode", 400)
		c.Set("errMessage", err.Error())
		return
	}

	// Convert UserDTO to User
	user := userDto.ToUser(userDto.Password)
	user.Money = money
	// Create the user
	err = u.service.CreateUser(c, user)
	if err != nil {
		u.logger.Error("Failed to create user", zap.Error(err))
		c.Status(http.StatusInternalServerError)
		c.Set("errCode", 400)
		c.Set("errMessage", err.Error())
		return
	}

	c.Status(http.StatusOK)
	c.Set("data", "注册成功")
}

// @BasePath /api/v1

// UserLogin godoc
// @Summary 用户登录
// @Description 用户登录并获取JWT token
// @Tags 用户
// @Accept  json
// @Produce  json
// @Param   user  body      domain.UserDTO  true  "用户信息"
// @Success 200   {object}  middleware.CommonResponse{data=string}  "{"token": "jwt_token"}"
// @Failure 400   {object}  middleware.Response  "{"errCode": 400, "errMessage": "错误信息"}"
// @Router /user/login [post]
func (u *UserHandler) UserLogin(c *gin.Context) {
	user := &domain.UserDTO{}
	if err := c.ShouldBind(user); err == nil {
		token, loginErr := u.service.Login(c.Request.Context(), user)
		if nil != loginErr {
			c.Status(http.StatusBadRequest)
			c.Set("errCode", 400)
			c.Set("errMessage", loginErr.Error())
		} else {
			c.Status(http.StatusOK)
			c.Set("data", token)
		}
	} else {
		u.logger.Info("UserLogin", zap.Error(err))
		c.Status(http.StatusBadRequest)
		c.Set("errCode", 400)
		c.Set("errMessage", err.Error())
	}
}

// @BasePath /api/v1

// UserQueryTest @Summary 用于测试获取用户信息
// @Description 获取用户信息接口
// @Tags 用户
// @Accept json
// @Produce json
// @Param id path uint64 true "用户ID"
// @Success 200 {object} middleware.CommonResponse{data=domain.User} "success"
// @Failure 400 {object} middleware.Response "error"
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
