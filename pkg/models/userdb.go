package models

import "gorm.io/gorm"

type Userdata struct {
	gorm.Model
	Username    string
	Email       string `gorm:"unique"`
	PhoneNumber string
	Password    string
	Status      bool
	Isadmin     bool
	Walet       int
}
