package entities

import (
	"gorm.io/gorm"
)

type Address struct {
	gorm.Model
	ID    uint
	Title string
	Lat   float64
	Lng   float64
	Users []*User `gorm:"many2many:user_addresses;"`
}
