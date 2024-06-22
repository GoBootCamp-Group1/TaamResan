package wallet

import (
	"TaamResan/internal/user"
	"context"
)

type Wallet struct {
	ID     uint
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

type Repo interface {
	Create(ctx context.Context) error
	CreateForUserID(ctx context.Context, userId uint) error
	Update(ctx context.Context, wallet *Wallet) error
	Delete(ctx context.Context, wallet *Wallet) error
	TopUp(ctx context.Context, wallet *Wallet, amount float64) error
	Expense(ctx context.Context, wallet *Wallet, amount float64) error

	StoreWalletCard(ctx context.Context, card *WalletCard) error
}
