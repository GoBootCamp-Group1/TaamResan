package wallet

import (
	"TaamResan/internal/user"
	"context"
)

type Wallet struct {
	ID     uint
	UserID uint
	Credit float64
	User   user.User
}

type WalletCard struct {
	ID       uint
	Wallet   Wallet
	BankName string
	Title    string
	Number   string
}

type WalletTopUp struct {
	Amount     float64
	CardNumber string
}

type WalletTransaction struct {
	WalletID   uint
	Type       uint
	Status     uint
	Amount     float64
	Additional map[string]interface{}
}

type WalletWithdraw struct {
	CardID uint
	Amount float64
	Status uint
}

const (
	TRANSACTION_TYPE_TOPUP uint = iota + 1
	TRANSACTION_TYPE_EXPENSE
	TRANSACTION_TYPE_WITHDRAW
	TRANSACTION_TYPE_UNKNOWN
)

const (
	TRANSACTION_STATUS_PENDING uint = iota + 1
	TRANSACTION_STATUS_DONE
	TRANSACTION_STATUS_FAILED
	TRANSACTION_STATUS_UNKNOWN
)

type Repo interface {
	Create(ctx context.Context) error
	CreateForUserID(ctx context.Context, userId uint) error
	Update(ctx context.Context, wallet *Wallet) error
	Delete(ctx context.Context, wallet *Wallet) error
	TopUp(ctx context.Context, w *WalletTopUp) error
	Expense(ctx context.Context, wallet *Wallet, amount float64) error

	StoreWalletCard(ctx context.Context, card *WalletCard) error
	DeleteWalletCard(ctx context.Context, card *WalletCard) error
	Withdraw(ctx context.Context, w *WalletWithdraw) error
	GetWalletByUserId(ctx context.Context, userId uint) (*Wallet, error)
}
