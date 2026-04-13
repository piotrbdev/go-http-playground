package middlewares

import (
	"gin-app/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ErrorHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()

		if len(c.Errors) == 0 {
			return
		}

		err := c.Errors.Last().Err

		if apiErr, ok := err.(*models.ApiError); ok {
			c.JSON(apiErr.Status, models.Response{
				Message: apiErr.Message,
			})
			c.Abort()
			return
		}

		c.JSON(http.StatusInternalServerError, models.Response{
			Message: "internal server error",
		})
	}
}
