package models

type CreateUserRequest struct {
	Name  string `json:"name" binding:"required,min=2"`
	Email string `json:"email" binding:"required,email"`
}

type UpdateUserRequest struct {
	Name  string `json:"name" binding:"required,min=2"`
	Email string `json:"email" binding:"required,email"`
}

// type User struct {
// 	ID    int    `json:"id"`
// 	Name  string `json:"name"`
// 	Email string `json:"email"`
// }

type User struct {
	ID    int    `gorm:"primaryKey"`
	Name  string `gorm:"not null"`
	Email string `gorm:"unique;not null"`
}
