package models

import "gorm.io/gorm"

type Cart struct {
	gorm.Model
	UserId int `gorm:"foriegnkey"`
	Total  int
}

type ShoppingCartItem struct {
	gorm.Model
	CartId        int `gorm:"foreignkey"`
	ProductItemID int `gorm:"foreignkey"`
	ProductName   string
	Quantity      int
	Total         int
}
