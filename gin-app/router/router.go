package router

import (
	"gin-app/handlers"
	"gin-app/middlewares"

	"github.com/gin-gonic/gin"
)

func SetupRouter(userHandler *handlers.UserHandler, todoHandler *handlers.TodosHandler) *gin.Engine {
	r := gin.New()

	r.Use(gin.Recovery())
	r.Use(middlewares.LogPathMiddleware)
	r.Use(middlewares.AppNameMiddleware)
	r.Use(middlewares.ErrorHandler())

	// r.POST("/signup", userHandler.SignUp)
	// r.POST("/login", userHandler.Login)

	api := r.Group("/api")

	private := api.Group("/private")
	private.Use(middlewares.AuthMiddleware())

	// RegisterUserRouters(private, userHandler)
	RegisterTodoRouters(private, todoHandler)

	return r
}
