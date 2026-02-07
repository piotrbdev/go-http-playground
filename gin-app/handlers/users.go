package handlers

import (
	"gin-app/models"
	"gin-app/storage"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetUser(c *gin.Context) {
	name := c.Param("name")

	if name == "" {
		c.JSON(http.StatusBadRequest, models.Response{Message: "name is required"})
		return
	}

	c.JSON(http.StatusOK, models.Response{Message: "User " + name})
}

func CreateUser(c *gin.Context) {
	var req models.CreateUserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, models.Response{Message: err.Error()})
		return
	}
	user := storage.AddUser(req.Name, req.Email)
	c.JSON(http.StatusCreated, user)
}

func GetUsers(c *gin.Context) {
	users := storage.GetUsers()
	c.JSON(http.StatusOK, users)
}
