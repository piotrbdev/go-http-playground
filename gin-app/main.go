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
	// db, err := storage.NewPostgresDB()
	db, err := storage.NewGormDB()
	if err != nil {
		panic(err)
	}
	// defer db.Close()

	// err = storage.CreateUserTable(db)
	// if err != nil {
	// 	panic(err)
	// }

	// store := storage.NewPostgresStorage(db)
	store := storage.NewGormStorage(db)

	users, err := store.GetUsers(1, 10)
	if err != nil {
		panic(err)
	}
	fmt.Println(users)

	userHandler := handlers.NewUserHandler(store)
	r := router.SetupRouter(userHandler)
	r.Run(":8080")
}
