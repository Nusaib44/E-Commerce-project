package models

import "gorm.io/gorm"

type Address struct {
	gorm.Model
	UserId      int
	HouseName   string
	Street      string
	AddressLine string
	City        string
	State       string
	Pincode     int
	Country     string
	IsDefault   bool
}
