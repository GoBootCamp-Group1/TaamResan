package service

import (
	"TaamResan/internal/wallet"
	"context"
	"errors"
	"fmt"
)

type WalletService struct {
	walletOps *wallet.Ops
}

var (
	ErrFetchingWallet = errors.New("can not fetch wallet")
	ErrCreatingWallet = errors.New("can not create wallet")
	ErrUpdatingWallet = errors.New("can not update wallet")
	ErrDeletingWallet = errors.New("can not delete wallet")
)

func NewWalletService(walletOps *wallet.Ops) *WalletService {
	return &WalletService{
		walletOps: walletOps,
	}
}

func (s *WalletService) CreateWallet(ctx context.Context) error {
	err := s.walletOps.Create(ctx)
	if err != nil {
		return fmt.Errorf(ErrCreatingWallet.Error()+": %w", err)
	}

	return nil
}

func (s *WalletService) CreateWalletForUserID(ctx context.Context, userId uint) error {
	err := s.walletOps.CreateForUser(ctx, userId)
	if err != nil {
		return fmt.Errorf(ErrCreatingWallet.Error()+": %w", err)
	}

	return nil
}

func (s *WalletService) UpdateWallet(ctx context.Context, wallet *wallet.Wallet) error {
	err := s.walletOps.Update(ctx, wallet)
	if err != nil {
		return fmt.Errorf(ErrUpdatingWallet.Error()+": %w", err)
	}

	return nil
}

func (s *WalletService) CreateWalletCard(ctx context.Context, card *wallet.WalletCard) error {
	err := s.walletOps.CreateWalletCard(ctx, card)
	if err != nil {
		return fmt.Errorf(ErrCreatingWallet.Error()+": %w", err)
	}

	return nil
}

func (s *WalletService) DeleteWalletCard(ctx context.Context, card *wallet.WalletCard) error {
	err := s.walletOps.DeleteWalletCard(ctx, card)
	if err != nil {
		return fmt.Errorf(ErrCreatingWallet.Error()+": %w", err)
	}

	return nil
}
