package router

import (
	"gin-app/handlers"
	"gin-app/middlewares"
	"gin-app/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func SetupRouter(userHandler *handlers.UserHandler) *gin.Engine {
	r := gin.New()

	r.Use(gin.Recovery())
	r.Use(middlewares.LogPathMiddleware)
	r.Use(middlewares.AppNameMiddleware)
	r.Use(middlewares.ErrorHandler())

	api := r.Group("/api")

	public := api.Group("/public")
	public.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, models.Response{Message: "public pong"})
	})

	private := api.Group("/private")
	// private.Use(middlewares.AuthMiddleware)
	private.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, models.Response{Message: "private pong"})
	})
	private_users := private.Group("/users")
	private_users.POST("", userHandler.CreateUser)
	private_users.GET("", userHandler.GetUsers)
	private_users.GET("/:id", userHandler.GetUser)
	private_users.PUT("/:id", userHandler.UpdateUser)
	private_users.DELETE("/:id", userHandler.DeleteUser)

	r.GET("/health", handlers.HealthCheck)
	r.GET("/users/:name", userHandler.GetUserByName)
	return r
}
