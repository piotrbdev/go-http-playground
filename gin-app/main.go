package main

import (
	"gin-app/handlers"
	"gin-app/middlewares"
	"gin-app/router"
	"gin-app/storage"

	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()
	db, err := storage.NewGormDB()
	if err != nil {
		panic(err)
	}

	err = middlewares.InitJWKS()
	if err != nil {
		panic(err)
	}

	todosStore := storage.NewTodosGormStorage(db)
	usersStore := storage.NewUserGormStorage(db)

	todoHandler := handlers.NewTodoHandler(todosStore)
	userHandler := handlers.NewUserHandler(usersStore)
	r := router.SetupRouter(userHandler, todoHandler)
	r.Run(":8080")
}
