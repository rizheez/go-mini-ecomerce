package entities

// ProductReview represents a product review (Optional - for future enhancement)
type ProductReview struct {
	ID                  int
	ProductID           int
	UserID              int
	OrderItemID         *int
	Rating              int
	Title               string
	ReviewText          string
	IsVerifiedPurchase  bool
	IsPublished         bool
	HelpfulCount        int
}