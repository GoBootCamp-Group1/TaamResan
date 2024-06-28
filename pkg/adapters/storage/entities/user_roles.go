package entities

import (
	"gorm.io/gorm"
)

type UserRoles struct {
	gorm.Model
	ID     uint
	UserId uint
	RoleId uint
}
