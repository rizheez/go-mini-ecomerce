package models

import (
	"time"
)

// OrderItem represents an item within an order
type OrderItem struct {
	ID              int     `gorm:"primaryKey;autoIncrement" json:"id"`
	OrderID         string  `gorm:"not null;type:varchar(50);index" json:"order_id"`
	ProductID       int     `gorm:"not null;index" json:"product_id"`
	ProductName     string  `gorm:"not null;type:varchar(255)" json:"product_name"` // Snapshot at time of order
	Quantity        int     `gorm:"not null" json:"quantity"`
	UnitPrice       float64 `gorm:"not null;type:decimal(10,2)" json:"unit_price"`
	TotalPrice      float64 `gorm:"not null;type:decimal(10,2)" json:"total_price"`
	ProductSnapshot JSONB   `gorm:"type:jsonb" json:"product_snapshot"` // Store product details at time of order
	CreatedAt       time.Time `gorm:"default:now()" json:"created_at"`
}