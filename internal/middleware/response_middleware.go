package middleware

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// ResponseMiddleware TODO 待测试
func ResponseMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()

		// 获取响应数据
		if c.Writer.Status() == http.StatusOK {
			// TODO 待核实 res为nil时候Get的返回结果exists
			if data, exists := c.Get("data"); exists {
				switch v := data.(type) {
				case []any:
					c.JSON(http.StatusOK, NewMultiResponse(v))
				default:
					c.JSON(http.StatusOK, NewSingleResponse(v))
				}
			} else {
				c.JSON(http.StatusOK, NewSuccessResponse())
			}
		} else {
			if errCode, exists := c.Get("errCode"); exists {
				if errMessage, exists := c.Get("errMessage"); exists {
					c.JSON(c.Writer.Status(), NewErrorResponse(errCode.(string), errMessage.(string)))
				}
			}
		}
	}
}
