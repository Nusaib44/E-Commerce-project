package models

import (
	"gorm.io/gorm"
)

type Walet struct {
	gorm.Model
	UserId  int
	Balance int
}
