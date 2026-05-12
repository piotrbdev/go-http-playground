package storage

import "gin-app/models"

type UsersStore interface {
	AddUser(name, email, password string) (models.User, error)
	GetUser(id int) (models.User, error)
	GetUsers(limit, offset int, filters map[string]string) ([]models.UserResponse, error)
	UpdateUser(id int, name, email string) (models.User, error)
	DeleteUser(id int) error
	CountUsers(filters map[string]string) (int64, error)
	IsEmailAvailable(email string) bool
	MatchUserPassword(email, password string) (*models.User, error)
}
