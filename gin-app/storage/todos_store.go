package storage

import "gin-app/models"

type TodosStore interface {
	CreateTodo(title string, userID string) (models.Todo, error)
	GetTodos(userID string) ([]models.TodoResponse, error)
}
