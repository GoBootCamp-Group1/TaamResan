package entities

import "gorm.io/gorm"

type Cart struct {
	gorm.Model
	UserId       uint
	RestaurantId uint
	Items        []*CartItem
}
