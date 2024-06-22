package wallet

import (
	"context"
)

type Ops struct {
	repo Repo
}

func NewOps(repo Repo) *Ops {
	return &Ops{
		repo: repo,
	}
}

func (o *Ops) Create(ctx context.Context) error {
	return o.repo.Create(ctx)
}

func (o *Ops) CreateForUser(ctx context.Context, userId uint) error {
	return o.repo.CreateForUserID(ctx, userId)
}

func (o *Ops) Update(ctx context.Context, wallet *Wallet) error {
	return o.repo.Update(ctx, wallet)
}

func (o *Ops) Delete(ctx context.Context, wallet *Wallet) error {
	return o.repo.Delete(ctx, wallet)
}

func (o *Ops) TopUp(ctx context.Context, wallet *Wallet, amount float64) error {
	return o.repo.TopUp(ctx, wallet, amount)
}

func (o *Ops) Expense(ctx context.Context, wallet *Wallet, amount float64) error {
	return o.repo.Expense(ctx, wallet, amount)
}

func (o *Ops) CreateWalletCard(ctx context.Context, card *WalletCard) error {
	return o.repo.StoreWalletCard(ctx, card)
}

func (o *Ops) DeleteWalletCard(ctx context.Context, card *WalletCard) error {
	return o.repo.DeleteWalletCard(ctx, card)
}
