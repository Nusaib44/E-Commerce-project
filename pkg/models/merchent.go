package models

type Merchent struct {
	Merchantname     string
	Email            string `gorm:"unique"`
	PhoneNumber      string
	Password         string
	Status           bool
	Verified         bool
	Companyname      string
	PayableCommetion string
	Location         string
}
