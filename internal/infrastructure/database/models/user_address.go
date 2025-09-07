package models

import (
	"time"
)

// UserAddress represents a user's shipping address
type UserAddress struct {
	ID            int       `gorm:"primaryKey;autoIncrement" json:"id"`
	UserID        int       `gorm:"not null;index" json:"user_id"`
	Label         string    `gorm:"not null;type:varchar(50)" json:"label"` // 'Home', 'Work', 'Other'
	RecipientName string    `gorm:"not null;type:varchar(255)" json:"recipient_name"`
	Phone         string    `gorm:"not null;type:varchar(20)" json:"phone"`
	AddressLine1  string    `gorm:"not null;type:varchar(255)" json:"address_line_1"`
	AddressLine2  string    `gorm:"type:varchar(255)" json:"address_line_2"`
	City          string    `gorm:"not null;type:varchar(100)" json:"city"`
	State         string    `gorm:"not null;type:varchar(100)" json:"state"`
	PostalCode    string    `gorm:"not null;type:varchar(20)" json:"postal_code"`
	Country       string    `gorm:"not null;type:varchar(100);default:'USA'" json:"country"`
	IsDefault     bool      `gorm:"default:false" json:"is_default"`
	CreatedAt     time.Time `gorm:"default:now()" json:"created_at"`
	UpdatedAt     time.Time `gorm:"default:now()" json:"updated_at"`
}