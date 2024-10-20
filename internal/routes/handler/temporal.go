package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type TemporalHandler struct {
}

func (temporalHandler *TemporalHandler) TemporalWelcome(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Welcome to Temporal",
	})
}
