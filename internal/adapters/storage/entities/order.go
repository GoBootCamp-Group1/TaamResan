package entities

import (
	"gorm.io/gorm"
	"time"
)

type Order struct {
	gorm.Model
	RestaurantID       uint
	UserID             uint
	AddressID          uint
	CustomerApprovedAt time.Time
	Status             uint
	Note               string

	Restaurant *Restaurant
	User       *User
	Address    *Address
}

type OrderItem struct {
	gorm.Model
	OrderId uint
	FoodId  uint
	Amount  float64
	Note    string
}
