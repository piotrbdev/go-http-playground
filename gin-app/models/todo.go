package models

type Todo struct {
	ID     int    `gorm:"primaryKey"`
	Title  string `gorm:"not null"`
	Done   bool   `gorm:"default:false"`
	UserID string `gorm:"not null"`
}

type TodoRequest struct {
	Title string `json:"title" binding:"required"`
}

type TodoResponse struct {
	Title string `json:"title"`
	Done  bool   `json:"done"`
}
