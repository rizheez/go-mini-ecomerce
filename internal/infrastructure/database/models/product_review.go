package models

import (
	"time"
)

// ProductReview represents a product review (Optional - for future enhancement)
type ProductReview struct {
	ID                  int       `gorm:"primaryKey;autoIncrement" json:"id"`
	ProductID           int       `gorm:"not null;index" json:"product_id"`
	UserID              int       `gorm:"not null;index" json:"user_id"`
	OrderItemID         *int      `gorm:"type:integer" json:"order_item_id"` // References order_items(id) ON DELETE SET NULL
	Rating              int       `gorm:"not null" json:"rating"`
	Title               string    `gorm:"type:varchar(255)" json:"title"`
	ReviewText          string    `gorm:"type:text" json:"review_text"`
	IsVerifiedPurchase  bool      `gorm:"default:false" json:"is_verified_purchase"`
	IsPublished         bool      `gorm:"default:true" json:"is_published"`
	HelpfulCount        int       `gorm:"default:0" json:"helpful_count"`
	CreatedAt           time.Time `gorm:"default:now()" json:"created_at"`
	UpdatedAt           time.Time `gorm:"default:now()" json:"updated_at"`
}