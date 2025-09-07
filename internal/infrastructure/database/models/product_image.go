package models

import (
	"time"
)

// ProductImage represents an image for a product
type ProductImage struct {
	ID        int       `gorm:"primaryKey;autoIncrement" json:"id"`
	ProductID int       `gorm:"not null;index" json:"product_id"`
	URL       string    `gorm:"not null;type:varchar(500)" json:"url"`
	AltText   string    `gorm:"type:varchar(255)" json:"alt_text"`
	IsPrimary bool      `gorm:"default:false" json:"is_primary"`
	SortOrder int       `gorm:"default:0" json:"sort_order"`
	CreatedAt time.Time `gorm:"default:now()" json:"created_at"`
}