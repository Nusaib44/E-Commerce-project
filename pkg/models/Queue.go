package models

import "gorm.io/gorm"

type ProductQueue struct {
	gorm.Model
	ProductName   string
	Category      int `gorm:"foreignkey"`
	Brand         string
	Price         int
	Quantity      int
	Image         string
	Description   string
	Total         int
}
