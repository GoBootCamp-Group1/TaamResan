package entities

import "gorm.io/gorm"

type RestaurantStaff struct {
	gorm.Model
	UserId       uint
	RestaurantId uint
	Position     uint
}
