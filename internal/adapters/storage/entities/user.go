package entities

import (
	"gorm.io/gorm"
	"time"
)

type User struct {
	gorm.Model
	ID        uint
	Uuid      string
	Name      string
	Email     string
	Mobile    string
	BirthDate time.Time
	Password  string
}
