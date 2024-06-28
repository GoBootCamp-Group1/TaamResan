package entities

import (
	"gorm.io/gorm"
	"time"
)

type User struct {
	gorm.Model
	Uuid      string
	Name      string
	Email     string
	Mobile    string
	BirthDate time.Time
	Password  string
	Addresses []*Address `gorm:"many2many:user_addresses;"`
}
