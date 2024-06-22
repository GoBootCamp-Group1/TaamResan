package entities

import (
	"gorm.io/gorm"
	"time"
)

type Wallet struct {
	gorm.Model
	ID     uint
	Credit float64
	UserID uint // Foreign key
	User   User // Associate
}

type WalletCard struct {
	gorm.Model
	ID       uint
	Wallet   Wallet
	WalletID uint
	Title    string
	BankName string
	Number   string
}

type WalletTransaction struct {
	ID         uint `gorm:"primarykey"`
	CreatedAt  time.Time
	UpdatedAt  time.Time
	Wallet     Wallet
	WalletID   uint
	Type       uint
	Status     uint
	Amount     float64
	Additional map[string]interface{} `gorm:"serializer:json"`
}
