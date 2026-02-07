package handlers

import (
	"gin-app/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func HealthCheck(c *gin.Context) {
	c.JSON(http.StatusOK, models.Response{Status: "ok"})
}
