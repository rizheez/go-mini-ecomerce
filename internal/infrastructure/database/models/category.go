package models

import (
	"time"
)

// Category represents a product category
type Category struct {
	ID          int       `gorm:"primaryKey;autoIncrement" json:"id"`
	Name        string    `gorm:"uniqueIndex;not null;type:varchar(255)" json:"name"`
	Description string    `gorm:"type:text" json:"description"`
	ImageURL    string    `gorm:"type:varchar(500)" json:"image_url"`
	IsActive    bool      `gorm:"default:true" json:"is_active"`
	SortOrder   int       `gorm:"default:0" json:"sort_order"`
	CreatedAt   time.Time `gorm:"default:now()" json:"created_at"`
	UpdatedAt   time.Time `gorm:"default:now()" json:"updated_at"`
}