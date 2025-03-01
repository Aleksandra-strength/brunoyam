package models

import "github.com/google/uuid"

type Task struct {
	ID          uuid.UUID `json:"id"`
	Title       string    `json:"title" binding:"required"`
	Description string    `json:"description"`
	Status      bool      `json:"status"`
	Priority    int       `json:"priority"`
}
