package entities

import "gorm.io/gorm"

type Category struct {
	gorm.Model
	ParentId     *uint
	RestaurantId uint
	CreatedBy    uint
	Name         string
}
