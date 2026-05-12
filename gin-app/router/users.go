package router

import (
	"gin-app/handlers"

	"github.com/gin-gonic/gin"
)

func RegisterUserRouters(rg *gin.RouterGroup, h *handlers.UserHandler) {
	users := rg.Group("/users")
	users.GET("", h.GetUsers)
	users.GET("/:id", h.GetUser)
	users.PUT("/:id", h.UpdateUser)
	users.DELETE("/:id", h.DeleteUser)
}
