package storage

import (
	"gin-app/models"
	"sync"
)

var (
	users  = make([]models.User, 0)
	mutex  = sync.Mutex{}
	nextID = 1
)

func AddUser(name, email string) models.User {
	mutex.Lock()
	defer mutex.Unlock()
	user := models.User{
		ID:    nextID,
		Name:  name,
		Email: email,
	}
	nextID++
	users = append(users, user)
	return user
}

func GetUsers() []models.User {
	mutex.Lock()
	defer mutex.Unlock()

	return users
}

func UpdateUser(id int, name, email string) (models.User, bool) {
	mutex.Lock()
	defer mutex.Unlock()

	for i, u := range users {
		if u.ID == id {
			users[i].Name = name
			users[i].Email = email
			return users[i], true
		}
	}
	return models.User{}, false
}

func DeleteUser(id int) bool {
	mutex.Lock()
	defer mutex.Unlock()

	for i, u := range users {
		if u.ID == id {
			users = append(users[:i], users[i+1:]...)
			return true
		}
	}
	return false
}
