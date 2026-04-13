package storage

import (
	"gin-app/models"
	"sync"
)

type MemoryStorage struct {
	users  []models.User
	mutex  sync.Mutex
	nextID int
}

func NewMemoryStorage() *MemoryStorage {
	return &MemoryStorage{
		users:  make([]models.User, 0),
		nextID: 1,
	}
}

func (m *MemoryStorage) AddUser(name, email string) (models.User, error) {
	m.mutex.Lock()
	defer m.mutex.Unlock()
	user := models.User{
		ID:    m.nextID,
		Name:  name,
		Email: email,
	}
	m.nextID++
	m.users = append(m.users, user)
	return user, nil
}

func (m *MemoryStorage) GetUser(id int) (models.User, bool) {
	m.mutex.Lock()
	defer m.mutex.Unlock()
	for i, u := range m.users {
		if u.ID == id {
			return m.users[i], true
		}
	}
	return models.User{}, false
}

func (m *MemoryStorage) GetUsers() []models.User {
	m.mutex.Lock()
	defer m.mutex.Unlock()

	return m.users
}

func (m *MemoryStorage) UpdateUser(id int, name, email string) (models.User, bool) {
	m.mutex.Lock()
	defer m.mutex.Unlock()

	for i, u := range m.users {
		if u.ID == id {
			m.users[i].Name = name
			m.users[i].Email = email
			return m.users[i], true
		}
	}
	return models.User{}, false
}

func (m *MemoryStorage) DeleteUser(id int) bool {
	m.mutex.Lock()
	defer m.mutex.Unlock()

	for i, u := range m.users {
		if u.ID == id {
			m.users = append(m.users[:i], m.users[i+1:]...)
			return true
		}
	}
	return false
}
