package handlers

import (
	"gin-app/models"
	"gin-app/storage"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/lib/pq"
	_ "github.com/lib/pq"
)

type UserHandler struct {
	store storage.UserStore
}

func NewUserHandler(store storage.UserStore) *UserHandler {
	return &UserHandler{
		store: store,
	}
}

func (h *UserHandler) GetUserByName(c *gin.Context) {
	name := c.Param("name")

	if name == "" {
		c.JSON(http.StatusBadRequest, models.Response{Message: "name is required"})
		return
	}

	c.JSON(http.StatusOK, models.Response{Message: "User " + name})
}

func (h *UserHandler) GetUser(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.Error(models.NewBadRequest("invalid id"))
		return
	}

	user, err := h.store.GetUser(id)
	if err != nil {
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
	user, err := h.store.AddUser(req.Name, req.Email)
	if err != nil {
		if pqErr, ok := err.(*pq.Error); ok {
			if pqErr.Code == "23505" {

				c.Error(models.NewBadRequest("email already exists"))
				return
			}
		}
		c.Error(models.NewBadRequest(err.Error()))
		return
	}
	c.JSON(http.StatusCreated, user)
}

func (h *UserHandler) GetUsers(c *gin.Context) {
	pageStr := c.Query("page")
	limitStr := c.Query("limit")

	page := 1
	limit := 10

	if p, err := strconv.Atoi(pageStr); err == nil && p > 0 {
		page = p
	}

	if l, err := strconv.Atoi(limitStr); err == nil && l > 0 {
		limit = l
	}

	offset := (page - 1) * limit

	users, err := h.store.GetUsers(limit, offset)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.Response{
			Message: "database error",
		})
		return
	}
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

	user, err := h.store.UpdateUser(id, req.Name, req.Email)
	if err != nil {
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
	err2 := h.store.DeleteUser(id)
	if err2 != nil {
		c.Error(models.NewNotFound("user not found"))
		return
	}

	c.JSON(http.StatusOK, models.Response{Message: "user deleted"})
}
