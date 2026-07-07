package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	ID        uuid.UUID      `gorm:"type:uuid;primary_key;default:gen_random_uuid()" json:"id"`
	Email     string         `gorm:"type:varchar(100);unique_index" json:"email" validate:"required,email"`
	Name      string         `gorm:"varchar(100)" json:"name"`
	Password  string         `gorm:"type:varchar(100)" json:"-" validate:"required"`
	Todo      []Todo         `gorm:"foreignKey:UserID" json:"todos,omitempty"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
