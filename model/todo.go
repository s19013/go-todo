package model

import (
	"time"
)

type Todo struct {
	ID          int       `json:"id"`
	Subject     string    `json:"subject"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
}

type CreateTodoRequest struct {
	Subject     string `json:"subject"`
	Description string `json:"description"`
}

type CreateTodoResponse struct {
	ID      int    `json:"id"`
	Message string `json:"message"`
}

type GetTodoRequest struct {
	ID int `json:"id"`
}

type GetTodoResponse struct {
	Todo Todo `json:"todo"`
}

type UpdateTodoRequest struct {
	ID          int    `json:"id"`
	Subject     string `json:"subject"`
	Description string `json:"description"`
}

type UpdateTodoResponse struct {
	Message string `json:"message"`
}
