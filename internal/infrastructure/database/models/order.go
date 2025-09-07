package models

import (
	"time"
)

// Order represents an order in the system
type Order struct {
	ID              string    `gorm:"primaryKey;type:varchar(50)" json:"id"` // Format: ORD-YYYYMMDDNNNNN
	UserID          int       `gorm:"not null" json:"user_id"`
	Status          string    `gorm:"not null;type:varchar(20);default:'pending'" json:"status"`
	Subtotal        float64   `gorm:"not null;type:decimal(10,2)" json:"subtotal"`
	ShippingCost    float64   `gorm:"not null;default:0;type:decimal(10,2)" json:"shipping_cost"`
	TaxAmount       float64   `gorm:"not null;default:0;type:decimal(10,2)" json:"tax_amount"`
	TotalAmount     float64   `gorm:"not null;type:decimal(10,2)" json:"total_amount"`
	PaymentMethod   string    `gorm:"not null;type:varchar(20)" json:"payment_method"`
	PaymentStatus   string    `gorm:"type:varchar(20);default:'pending'" json:"payment_status"`
	ShippingAddress JSONB     `gorm:"not null;type:jsonb" json:"shipping_address"` // Store complete address snapshot
	TrackingNumber  string    `gorm:"type:varchar(100)" json:"tracking_number"`
	Notes           string    `gorm:"type:text" json:"notes"`
	CancelledAt     time.Time `gorm:"type:timestamp with time zone" json:"cancelled_at"`
	ShippedAt       time.Time `gorm:"type:timestamp with time zone" json:"shipped_at"`
	DeliveredAt     time.Time `gorm:"type:timestamp with time zone" json:"delivered_at"`
	CreatedAt       time.Time `gorm:"default:now()" json:"created_at"`
	UpdatedAt       time.Time `gorm:"default:now()" json:"updated_at"`
}