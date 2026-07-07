package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Todo struct {
	ID          uuid.UUID      `gorm:"type:uuid;primary_key;default:gen_random_uuid()" json:"id"`
	Title       string         `gorm:"type:varchar(100)" json:"title" validate:"required"`
	Description string         `gorm:"type:text" json:"description"`
	Status      string         `gorm:"type:varchar(100);default:pendding" json:"status"`
	UserID      uuid.UUID      `json:"user_id"`
	User        *User          `gorm:"foreignKey:UserID" json:"user,omitempty"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"deleted_at,omitempty"`
}
