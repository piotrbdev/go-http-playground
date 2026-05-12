package handlers

import (
	"gin-app/models"
	"gin-app/storage"
	"gin-app/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

type TodosHandler struct {
	store storage.TodosStore
}

func NewTodoHandler(store storage.TodosStore) *TodosHandler {
	return &TodosHandler{
		store: store,
	}
}

func (h *TodosHandler) CreateTodo(c *gin.Context) {
	var req models.TodoRequest
	userID, err := utils.GetUserID(c)
	if err != nil {
		c.Error(models.NewUnauthorized("invalid token"))
		return
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.Error(models.NewBadRequest("invalid request body"))
		return
	}
	todo, err := h.store.CreateTodo(req.Title, userID)
	if err != nil {
		c.Error(models.NewBadRequest(err.Error()))
		return
	}
	c.JSON(
		http.StatusCreated,
		models.TodoResponse{Title: todo.Title, Done: todo.Done},
	)
}

func (h *TodosHandler) GetTodos(c *gin.Context) {
	userID, err := utils.GetUserID(c)
	if err != nil {
		c.Error(models.NewUnauthorized("invalid token"))
		return
	}
	todos, err := h.store.GetTodos(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.Response{
			Message: "database error",
		})
		return
	}

	c.JSON(http.StatusOK, todos)
}

// TODO: do getbyid, put, delete for one user
