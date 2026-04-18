package storage

import "gin-app/models"

type UserStore interface {
	AddUser(name, email string) (models.User, error)
	GetUser(id int) (models.User, error)
	GetUsers(limit, offset int) ([]models.User, error)
	UpdateUser(id int, name, email string) (models.User, error)
	DeleteUser(id int) error
}
