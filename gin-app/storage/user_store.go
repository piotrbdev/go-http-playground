package storage

import "gin-app/models"

type UserStore interface {
	AddUser(name, email string) models.User
	GetUser(id int) (models.User, bool)
	GetUsers() []models.User
	UpdateUser(id int, name, email string) (models.User, bool)
	DeleteUser(id int) bool
}
