package entities

import "gorm.io/gorm"

type CartItem struct {
	gorm.Model
	CartId uint
	FoodId uint
	Amount uint
	Note   string
}
