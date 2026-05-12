package storage

import (
	"gin-app/models"

	"gorm.io/gorm"
)

type TodosGormStorage struct {
	db *gorm.DB
}

func NewTodosGormStorage(db *gorm.DB) *TodosGormStorage {
	return &TodosGormStorage{db: db}
}

// TODO: implement this method then do GetTodos and add it to interface
func (g *TodosGormStorage) CreateTodo(title string, userID string) (models.Todo, error) {
	var todo models.Todo
	todo = models.Todo{
		Title:  title,
		UserID: userID,
	}
	err := g.db.Create(&todo).Error
	return todo, err
}

func (g *TodosGormStorage) GetTodos(userID string) ([]models.TodoResponse, error) {
	var todos []models.TodoResponse
	err := g.db.Model(&models.Todo{}).Where("user_id=?", userID).Find(&todos).Error
	return todos, err
}
