package entities

import (
	"gorm.io/gorm"
)

type Restaurant struct {
	gorm.Model
	Name           string
	OwnedBy        uint
	ApprovalStatus uint
	AddressId      uint
	CourierSpeed   float64
}
