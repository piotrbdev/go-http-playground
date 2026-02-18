package main

import (
	"gin-app/handlers"
	"gin-app/router"
	"gin-app/storage"
)

func main() {
	store := storage.NewMemoryStorage()
	userHandler := handlers.NewUserHandler(store)
	r := router.SetupRouter(userHandler)
	r.Run(":8080")
}
