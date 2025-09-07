package models

import (
	"time"
)

// OrderStatusHistory tracks order status changes for audit trail
type OrderStatusHistory struct {
	ID        int       `gorm:"primaryKey;autoIncrement" json:"id"`
	OrderID   string    `gorm:"not null;type:varchar(50);index" json:"order_id"`
	FromStatus string   `gorm:"type:varchar(20)" json:"from_status"`
	ToStatus  string    `gorm:"not null;type:varchar(20)" json:"to_status"`
	Note      string    `gorm:"type:text" json:"note"`
	ChangedBy *int      `gorm:"type:integer" json:"changed_by"` // References users(id) ON DELETE SET NULL
	CreatedAt time.Time `gorm:"default:now()" json:"created_at"`
}