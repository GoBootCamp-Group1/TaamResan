package entities

import "time"

type UserAddress struct {
	UserID    uint `gorm:"primaryKey"`
	AddressID uint `gorm:"primaryKey"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
