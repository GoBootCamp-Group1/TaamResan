package entities

import (
	"gorm.io/gorm"
)

type Wallet struct {
	gorm.Model
	ID     uint
	Credit float64
	UserID uint // Foreign key
	User   User // Associate
}
