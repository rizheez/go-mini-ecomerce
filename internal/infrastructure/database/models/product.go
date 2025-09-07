package models

import (
	"time"
)

// Product represents a product in the catalog
type Product struct {
	ID              int             `gorm:"primaryKey;autoIncrement" json:"id"`
	Name            string          `gorm:"not null;type:varchar(255)" json:"name"`
	Description     string          `gorm:"type:text" json:"description"`
	Price           float64         `gorm:"not null;type:decimal(10,2)" json:"price"`
	StockQuantity   int             `gorm:"not null;default:0" json:"stock_quantity"`
	CategoryID      int             `gorm:"not null;index" json:"category_id"`
	SKU             string          `gorm:"uniqueIndex;type:varchar(100)" json:"sku"`
	Specifications  JSONB           `gorm:"type:jsonb" json:"specifications"`
	IsActive        bool            `gorm:"default:true" json:"is_active"`
	Weight          float64         `gorm:"type:decimal(8,2)" json:"weight"`
	Dimensions      JSONB           `gorm:"type:jsonb" json:"dimensions"`
	CreatedAt       time.Time       `gorm:"default:now()" json:"created_at"`
	UpdatedAt       time.Time       `gorm:"default:now()" json:"updated_at"`
}