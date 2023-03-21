package response

type Checkout struct {
	Message string
	Total   int
	Walet   int
	Address interface{}
	Items   interface{}
	Payment interface{}
}

type Address struct {
	HouseName   string
	Street      string
	AddressLine string
	City        string
	State       string
	Pincode     int
	Country     string
	IsDefault   bool
}
