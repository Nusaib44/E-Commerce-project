package models

import "gorm.io/gorm"

type ProductOffers struct {
	gorm.Model
	Product    int
	OfferPrice int
}

type CategoryOffers struct {
	gorm.Model
	Category   int
	OfferPrice int
}
