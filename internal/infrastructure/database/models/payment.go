package models

import (
	"time"
)

// Payment represents a payment transaction record
type Payment struct {
	ID              string    `gorm:"primaryKey;type:varchar(50)" json:"id"` // Format: PAY-YYYYMMDDNNNNN
	OrderID         string    `gorm:"not null;type:varchar(50);index" json:"order_id"`
	Amount          float64   `gorm:"not null;type:decimal(10,2)" json:"amount"`
	PaymentMethod   string    `gorm:"not null;type:varchar(20)" json:"payment_method"`
	Status          string    `gorm:"not null;type:varchar(20);default:'pending'" json:"status"`
	TransactionID   string    `gorm:"type:varchar(100);index" json:"transaction_id"` // External payment processor transaction ID
	PaymentDetails  JSONB     `gorm:"type:jsonb" json:"payment_details"`  // Store payment method specific details (masked)
	GatewayResponse JSONB     `gorm:"type:jsonb" json:"gateway_response"` // Store payment gateway response
	ProcessedAt     time.Time `gorm:"type:timestamp with time zone" json:"processed_at"`
	FailedAt        time.Time `gorm:"type:timestamp with time zone" json:"failed_at"`
	FailureReason   string    `gorm:"type:text" json:"failure_reason"`
	CreatedAt       time.Time `gorm:"default:now()" json:"created_at"`
	UpdatedAt       time.Time `gorm:"default:now()" json:"updated_at"`
}