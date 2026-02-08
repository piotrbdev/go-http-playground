package main

import (
	"gin-app/handlers"
	"gin-app/middlewares"
	"gin-app/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.Use(middlewares.LogPathMiddleware)
	r.Use(middlewares.AppNameMiddleware)
	r.Use(middlewares.ErrorHandler())

	api := r.Group("/api")

	public := api.Group("/public")
	public.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, models.Response{Message: "public pong"})
	})

	private := api.Group("/private")
	private.Use(middlewares.AuthMiddleware)
	private.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, models.Response{Message: "private pong"})
	})

	private.POST("/users", handlers.CreateUser)
	private.GET("/users", handlers.GetUsers)
	private.PUT("/users/:id", handlers.UpdateUser)
	private.DELETE("/users/:id", handlers.DeleteUser)

	r.GET("/health", handlers.HealthCheck)
	r.GET("/users/:name", handlers.GetUser)

	r.Run(":8080")
}
