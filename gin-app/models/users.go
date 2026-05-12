package models

type CreateUserRequest struct {
	Name     string `json:"name" binding:"required,min=2"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=8"`
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
	ID       int    `gorm:"primaryKey"`
	Name     string `gorm:"not null"`
	Email    string `gorm:"index;unique;not null"`
	Password string `gorm:"not null"`
}

type LoginRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

type UserResponse struct {
	Email string
	Name  string
}
