package models

import "gorm.io/gorm"

type Product struct {
	gorm.Model
	ProductName   string
	Category      int `gorm:"foreignkey"`
	Brand         string
	Price         int
	ProductOffer  int
	CategoryOffer int
	Total         int
	Quantity      int
	Image         string
	Description   string
	Count         int
	Orderd        int
}
