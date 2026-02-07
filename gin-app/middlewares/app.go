package middlewares

import "github.com/gin-gonic/gin"

func AppNameMiddleware(c *gin.Context) {
	c.Header("X-App-Name", "go-http-playground")
	c.Next()
}
