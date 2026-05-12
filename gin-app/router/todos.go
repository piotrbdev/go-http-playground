package router

import (
	"gin-app/handlers"

	"github.com/gin-gonic/gin"
)

func RegisterTodoRouters(rg *gin.RouterGroup, h *handlers.TodosHandler) {
	todos := rg.Group("/todos")
	todos.POST("", h.CreateTodo)
	todos.GET("", h.GetTodos)
}
