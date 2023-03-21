package models

import "gorm.io/gorm"

type Order struct {
	gorm.Model
	UserID    int
	ProductId int
	Quantity  int
	Price     int
	Payment   string
	Address   int
	Status    string
}
type Return struct {
	gorm.Model
	UserID    int
	ProductID int
	Quantity  int
	Price     int
	Status    string
}
