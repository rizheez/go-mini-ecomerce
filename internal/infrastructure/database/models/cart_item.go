package models

import (
	"time"
)

// CartItem represents an item in a shopping cart
type CartItem struct {
	ID        int       `gorm:"primaryKey;autoIncrement" json:"id"`
	CartID    int       `gorm:"not null;index" json:"cart_id"`
	ProductID int       `gorm:"not null;index" json:"product_id"`
	Quantity  int       `gorm:"not null" json:"quantity"`
	UnitPrice float64   `gorm:"not null;type:decimal(10,2)" json:"unit_price"`
	AddedAt   time.Time `gorm:"default:now()" json:"added_at"`
	UpdatedAt time.Time `gorm:"default:now()" json:"updated_at"`
}