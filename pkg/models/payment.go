package models

import "gorm.io/gorm"

type Payment struct {
	gorm.Model
	PaymentMethod string
}
