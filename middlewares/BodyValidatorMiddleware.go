package middlewares

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func ValidateBodyForPOSTRequest() gin.HandlerFunc {
	return func(c *gin.Context) {
		if c.Request.Method != "POST" {
			c.Next()
			return
		}
		if c.Request.ContentLength == 0 {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"status":  http.StatusBadRequest,
				"message": "POST request must have body",
			})
			return
		}
	}
}
