package models

type CreateUserRequest struct {
	Name  string `json:"name" binding:"required,min=2"`
	Email string `json:"email" binding:"required,email"`
}
