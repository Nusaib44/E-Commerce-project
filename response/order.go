package response

type Update struct {
	ID        uint
	ProductId int
	Quantity  int
	Price     int
	Payment   string
	Address   interface{}
	Status    string
}
