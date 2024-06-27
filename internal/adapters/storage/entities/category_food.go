package entities

import "gorm.io/gorm"

type CategoryFood struct {
	gorm.Model
	CategoryId uint
	FoodId     uint
}
