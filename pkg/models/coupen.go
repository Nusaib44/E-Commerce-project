package models

import "gorm.io/gorm"

type Coupen struct {
	gorm.Model
	Code   string
	Value  int
	Limit  int64
	Expire int64
	IsUsed bool
}
