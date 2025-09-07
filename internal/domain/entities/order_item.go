package entities

// OrderItem represents an item within an order
type OrderItem struct {
	ID              int
	OrderID         string
	ProductID       int
	ProductName     string // Snapshot at time of order
	Quantity        int
	UnitPrice       float64
	TotalPrice      float64
	ProductSnapshot map[string]interface{} // Store product details at time of order
}