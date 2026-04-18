package models

type Response struct {
	Status  string `json:"status,omitempty"`
	Message string `json:"message,omitempty"`
}

type PaginatedUsersResponse struct {
	Data  []User `json:"data"`
	Page  int    `json:"page"`
	Limit int    `json:"limit"`
	Total int    `json:"total"`
}
