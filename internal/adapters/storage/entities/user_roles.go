package entities

import (
	"gorm.io/gorm"
)

type UserRoles struct {
	gorm.Model
	ID     uint
	Uuid   string
	UserId uint
	RoleId uint
}
