package entities

// Payment represents a payment transaction record
type Payment struct {
	ID              string
	OrderID         string
	Amount          float64
	PaymentMethod   string
	Status          string
	TransactionID   string // External payment processor transaction ID
	PaymentDetails  map[string]interface{}  // Store payment method specific details (masked)
	GatewayResponse map[string]interface{} // Store payment gateway response
}