package entities

import "gorm.io/gorm"

type CartItem struct {
	gorm.Model
	CartId uint
	FoodId uint
	Amount float64
	Note   string

	Cart *Cart
	Food *Food
}
