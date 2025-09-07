package entities

// Order represents an order in the system
type Order struct {
	ID              string
	UserID          int
	Status          string
	Subtotal        float64
	ShippingCost    float64
	TaxAmount       float64
	TotalAmount     float64
	PaymentMethod   string
	PaymentStatus   string
	ShippingAddress map[string]interface{} // Store complete address snapshot
	TrackingNumber  string
	Notes           string
}