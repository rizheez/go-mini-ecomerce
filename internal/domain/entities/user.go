package entities

// User represents a user in the system
type User struct {
	ID           int
	Email        string
	PasswordHash string
	Name         string
	Phone        string
	Role         string
	EmailVerified bool
	IsActive     bool
}