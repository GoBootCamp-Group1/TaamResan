package entities

import "gorm.io/gorm"

type Food struct {
	gorm.Model
	RestaurantId       uint
	CreatedBy          uint
	Name               string
	Price              float64
	CancelRate         float64
	PreparationMinutes uint
	Restaurant         *Restaurant
	Categories         []*Category `gorm:"many2many:category_foods;"`
}
