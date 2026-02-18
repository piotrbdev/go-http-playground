package handlers

import (
	"gin-app/models"
	"gin-app/storage"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	store storage.UserStore
}

func NewUserHandler(store storage.UserStore) *UserHandler {
	return &UserHandler{
		store: store,
	}
}

func (h *UserHandler) GetUser(c *gin.Context) {
	name := c.Param("name")

	if name == "" {
		c.JSON(http.StatusBadRequest, models.Response{Message: "name is required"})
		return
	}

	c.JSON(http.StatusOK, models.Response{Message: "User " + name})
}

func (h *UserHandler) GetUserByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.Error(models.NewBadRequest("invalid id"))
		return
	}

	user, ok := h.store.GetUser(id)
	if !ok {
		c.Error(models.NewNotFound("user not found"))
		return
	}
	c.JSON(http.StatusOK, user)
}

func (h *UserHandler) CreateUser(c *gin.Context) {
	var req models.CreateUserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.Error(models.NewBadRequest("invalid request body"))
		return
	}
	user := h.store.AddUser(req.Name, req.Email)
	c.JSON(http.StatusCreated, user)
}

func (h *UserHandler) GetUsers(c *gin.Context) {
	users := h.store.GetUsers()
	c.JSON(http.StatusOK, users)
}

func (h *UserHandler) UpdateUser(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.Error(models.NewBadRequest("invalid id"))
		return
	}

	var req models.UpdateUserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.Error(models.NewBadRequest("invalid request body"))
		return
	}

	user, ok := h.store.UpdateUser(id, req.Name, req.Email)
	if !ok {
		c.Error(models.NewNotFound("user not found"))
		return
	}
	c.JSON(http.StatusOK, user)
}

func (h *UserHandler) DeleteUser(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.Error(models.NewBadRequest("invalid id"))
		return
	}

	if !h.store.DeleteUser(id) {
		c.Error(models.NewNotFound("user not found"))
	}

	c.JSON(http.StatusOK, models.Response{Message: "user deleted"})
}
