package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Comment struct {
	ID        uuid.UUID      `gorm:"type:uuid;primary_key;default:gen_random_uuid()" json:"id"`
	Content   string         `gorm:"type:text" json:"content" validate:"required"`
	TodoID    uuid.UUID      `gorm:"type:uuid;index" json:"todo_id"`
	Todo      *Todo          `gorm:"foreignKey:TodoID" json:"todo,omitempty"`
	UserID    uuid.UUID      `gorm:"type:uuid;index" json:"user_id"`
	User      *User          `gorm:"foreignKey:UserID" json:"user,omitempty"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at,omitempty"`
}
