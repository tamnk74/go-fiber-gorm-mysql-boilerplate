package models

import (
	"time"

	"gorm.io/gorm"
)

// Item type
type Item struct {
	ID          uint           `gorm:"primary_key" json:"id"`
	Name        string         `gorm:"type:varchar(255);unique,not null" json:"name"`
	Description string         `json:"description"`
	Status      int            `json:"status"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `json:"deleted_at"`
}
