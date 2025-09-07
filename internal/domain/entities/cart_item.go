package entities

// CartItem represents an item in a shopping cart
type CartItem struct {
	ID        int
	CartID    int
	ProductID int
	Quantity  int
	UnitPrice float64
}