package main

import (
	"fmt"
	"gin-app/handlers"
	"gin-app/router"
	"gin-app/storage"

	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()
	db, err := storage.NewPostgresDB()
	if err != nil {
		panic(err)
	}
	defer db.Close()

	err = storage.CreateUserTable(db)
	if err != nil {
		panic(err)
	}

	store := storage.NewPostgresStorage(db)

	users, err := store.GetUsers()
	if err != nil {
		panic(err)
	}
	fmt.Println(users)

	userHandler := handlers.NewUserHandler(store)
	r := router.SetupRouter(userHandler)
	r.Run(":8080")
}
