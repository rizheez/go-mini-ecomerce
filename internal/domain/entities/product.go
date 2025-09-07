package entities

// Product represents a product in the catalog
type Product struct {
	ID              int
	Name            string
	Description     string
	Price           float64
	StockQuantity   int
	CategoryID      int
	SKU             string
	Specifications  map[string]interface{}
	IsActive        bool
	Weight          float64
	Dimensions      map[string]interface{}
}