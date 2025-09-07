package entities

// UserAddress represents a user's shipping address
type UserAddress struct {
	ID            int
	UserID        int
	Label         string // 'Home', 'Work', 'Other'
	RecipientName string
	Phone         string
	AddressLine1  string
	AddressLine2  string
	City          string
	State         string
	PostalCode    string
	Country       string
	IsDefault     bool
}