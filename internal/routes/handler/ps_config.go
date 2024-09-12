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
