package entities

// Category represents a product category
type Category struct {
	ID          int
	Name        string
	Description string
	ImageURL    string
	IsActive    bool
	SortOrder   int
}