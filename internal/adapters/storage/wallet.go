package storage

import (
	"TaamResan/internal/adapters/storage/entities"
	"TaamResan/internal/adapters/storage/mappers"
	"TaamResan/internal/wallet"
	"TaamResan/pkg/jwt"
	"context"
	"gorm.io/gorm"
)

type walletRepo struct {
	db *gorm.DB
}

func NewWalletRepo(db *gorm.DB) wallet.Repo {
	return &walletRepo{
		db: db,
	}
}

func (w *walletRepo) Create(ctx context.Context) error {

	userID := ctx.Value(jwt.UserClaimKey).(*jwt.UserClaims).UserID

	walletEntity := entities.Wallet{
		Credit: 0.0,
		UserID: userID,
	}

	if err := w.db.Create(&walletEntity).Error; err != nil {
		return err
	}

	return nil
}

func (w *walletRepo) CreateForUserID(ctx context.Context, userId uint) error {

	walletEntity := entities.Wallet{
		Credit: 0.0,
		UserID: userId,
	}

	if err := w.db.Create(&walletEntity).Error; err != nil {
		return err
	}

	return nil
}

func (w *walletRepo) Update(ctx context.Context, wallet *wallet.Wallet) error {
	//TODO implement me
	panic("implement me")
}

func (w *walletRepo) Delete(ctx context.Context, wallet *wallet.Wallet) error {
	//TODO implement me
	panic("implement me")
}

func (w *walletRepo) TopUp(ctx context.Context, wallet *wallet.Wallet, amount float64) error {
	//TODO implement me
	panic("implement me")
}

func (w *walletRepo) Expense(ctx context.Context, wallet *wallet.Wallet, amount float64) error {
	//TODO implement me
	panic("implement me")
}

func (w *walletRepo) StoreWalletCard(ctx context.Context, card *wallet.WalletCard) error {
	//get wallet id
	userID := ctx.Value(jwt.UserClaimKey).(*jwt.UserClaims).UserID

	var walletEntity entities.Wallet
	if err := w.db.Debug().Model(&entities.Wallet{}).Where("user_id = ?", userID).Find(&walletEntity).Error; err != nil {
		return err
	}

	walletCardEntity := mappers.DomainToWalletCardEntity(card)

	walletCardEntity.WalletID = walletEntity.ID

	//store in database
	if err := w.db.Create(&walletCardEntity).Error; err != nil {
		return err
	}

	return nil
}
