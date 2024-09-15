package handler

import (
	"fx-web/internal/domain"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
	"strconv"
)

type ConfigHandler struct {
	service domain.PsConfigService
	logger  *zap.Logger
}

// @BasePath /api/v1

// List @Summary 列出配置
// @Description 分页获取配置列表
// @Tags config
// @Accept json
// @Produce json
// @Param pageNo path int true "页码"
// @Param pageSize path int true "每页数量"
// @Success 200 {array} domain.PsConfig "配置列表"
// @Failure 400 {object} any "请求错误"
// @Router /configs/list [get]
func (h ConfigHandler) List(c *gin.Context) {
	pageNoString := c.Param("pageNo")
	pageSizString := c.Param("pageSize")
	pageNo, _ := strconv.Atoi(pageNoString)
	pageSize, _ := strconv.Atoi(pageSizString)
	configs, err := h.service.ListConfigs(c, pageSize, pageNo)
	if err != nil {
		c.Set("errCode", 400)
		c.Set("errMessage", err.Error())
		c.Status(http.StatusBadRequest)
	} else {
		c.Set("data", configs)
		c.Status(http.StatusOK)
	}
}

func NewPsConfigHandler(service domain.PsConfigService, logger *zap.Logger) *ConfigHandler {
	return &ConfigHandler{
		service: service,
		logger:  logger,
	}
}
