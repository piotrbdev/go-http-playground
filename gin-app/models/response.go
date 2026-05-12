package models

type Response struct {
	Status  string `json:"status,omitempty"`
	Message string `json:"message,omitempty"`
}

type TokenResponse struct {
	Token string `json:"token"`
}

type PaginatedUsersResponse struct {
	Data  []UserResponse `json:"data"`
	Page  int            `json:"page"`
	Limit int            `json:"limit"`
	Total int            `json:"total"`
}
