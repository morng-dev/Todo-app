package entities

import (
	"time"

	"github.com/google/uuid"
)

type Todo struct {
	ID          uuid.UUID `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Status      string    `json:"status"`
	UserID      uuid.UUID `json:"user_id"`
	User        *User     `json:"user,omitempty"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type TodoRequest struct {
	Title       string `json:"title"`
	Description string `json:"description"`
}

type TodoStatusRequest struct {
	Status string `json:"status" validate:"required"`
}
