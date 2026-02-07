package middlewares

import (
	"gin-app/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware(c *gin.Context) {
	token := c.GetHeader("Authorization")
	if token == "" {
		c.JSON(http.StatusUnauthorized, models.Response{Message: "missing auth token"})
		c.Abort()
		return
	}
	c.Next()
}
