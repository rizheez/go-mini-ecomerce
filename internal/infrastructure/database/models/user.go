package models

import (
	"time"
)

// User represents a user in the system
type User struct {
	ID            int       `gorm:"primaryKey;autoIncrement" json:"id"`
	Email         string    `gorm:"uniqueIndex;not null;type:varchar(255)" json:"email"`
	Password      string    `gorm:"not null;type:varchar(255)" json:"-"`
	Name          string    `gorm:"not null;type:varchar(255)" json:"name"`
	Phone         string    `gorm:"type:varchar(20)" json:"phone"`
	Role          string    `gorm:"type:varchar(20);default:'customer'" json:"role"`
	EmailVerified bool      `gorm:"default:true" json:"email_verified"`
	IsActive      bool      `gorm:"default:true" json:"is_active"`
	CreatedAt     time.Time `gorm:"default:now()" json:"created_at"`
	UpdatedAt     time.Time `gorm:"default:now()" json:"updated_at"`
}
