package response

type CartProduct struct {
	ID            int
	ProductName   string
	Category      int
	Brand         string
	Price         int
	ProductOffer  int
	CategoryOffer int
	Total         int
	Image         string
	Description   string
	Quantity      int
}
type Cart struct {
	Total int
	Walet int
	// Product interface{}
}
