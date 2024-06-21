package entities

import (
	"gorm.io/gorm"
)

type Role struct {
	gorm.Model
	Uuid string
	Name string
}
