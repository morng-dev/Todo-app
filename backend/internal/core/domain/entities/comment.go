package entities

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Comment struct {
	ID        uuid.UUID      `json:"id"`
	Content   string         `json:"content"`
	TodoID    uuid.UUID      `json:"todo_id"`
	Todo      *Todo          `json:"todo,omitempty"`
	UserID    uuid.UUID      `json:"user_id,omitempty"`
	User      *User          `json:"user"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at,omitempty"`
}
