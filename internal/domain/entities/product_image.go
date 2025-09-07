package entities

// ProductImage represents an image for a product
type ProductImage struct {
	ID        int
	ProductID int
	URL       string
	AltText   string
	IsPrimary bool
	SortOrder int
}