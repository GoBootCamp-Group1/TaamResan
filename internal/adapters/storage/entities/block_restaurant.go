package entities

import "gorm.io/gorm"

type BlockRestaurant struct {
	gorm.Model
	UserId       uint
	RestaurantId uint
}
