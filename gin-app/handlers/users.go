package handlers

import (
	"gin-app/models"
	"gin-app/storage"
	"net/http"
	"strconv"

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
		c.Error(models.NewBadRequest("invalid request body"))
		return
	}
	user := storage.AddUser(req.Name, req.Email)
	c.JSON(http.StatusCreated, user)
}

func GetUsers(c *gin.Context) {
	users := storage.GetUsers()
	c.JSON(http.StatusOK, users)
}

func UpdateUser(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		// c.JSON(http.StatusBadRequest, models.Response{Message: "invalid id"})
		c.Error(models.NewBadRequest("invalid id"))
		return
	}

	var req models.UpdateUserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		// c.JSON(http.StatusBadRequest, models.Response{Message: err.Error()})
		c.Error(models.NewBadRequest("invalid request body"))
		return
	}

	user, ok := storage.UpdateUser(id, req.Name, req.Email)
	if !ok {
		// c.JSON(http.StatusNotFound, models.Response{Message: "user not found"})
		c.Error(models.NewNotFound("user not found"))
		return
	}
	c.JSON(http.StatusOK, user)
}

func DeleteUser(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		// c.JSON(http.StatusBadRequest, models.Response{Message: "invalid id"})
		c.Error(models.NewBadRequest("invalid id"))
		return
	}

	if !storage.DeleteUser(id) {
		// c.JSON(http.StatusNotFound, models.Response{Message: "user not found"})
		c.Error(models.NewNotFound("user not found"))
	}

	c.JSON(http.StatusOK, models.Response{Message: "user deleted"})
}
