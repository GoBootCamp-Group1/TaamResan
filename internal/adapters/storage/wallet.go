package storage

import (
	"TaamResan/internal/adapters/storage/entities"
	"TaamResan/internal/adapters/storage/mappers"
	"TaamResan/internal/wallet"
	"TaamResan/pkg/jwt"
	"context"
	"fmt"
	"gorm.io/gorm"
)

type walletRepo struct {
	db       *gorm.DB
	userRepo userRepo
}

func NewWalletRepo(db *gorm.DB) wallet.Repo {
	return &walletRepo{
		db: db,
	}
}

func (w *walletRepo) GetWalletByUserId(ctx context.Context, userId uint) (*wallet.Wallet, error) {

	var walletEntity *entities.Wallet
	if err := w.db.WithContext(ctx).Model(&entities.Wallet{}).Where("user_id = ?", userId).Find(&walletEntity).Error; err != nil {
		return nil, err
	}

	return mappers.WalletEntityToDomain(walletEntity), nil
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

	if err := w.db.WithContext(ctx).Create(&walletEntity).Error; err != nil {
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

func (w *walletRepo) TopUp(ctx context.Context, walletTopUp *wallet.WalletTopUp) error {
	//check for card existence
	var walletCard entities.WalletCard
	if err := w.db.WithContext(ctx).Model(&entities.WalletCard{}).Where("number = ?", walletTopUp.CardNumber).Find(&walletCard).Error; err != nil {
		return err
	}

	if walletCard.ID == 0 {
		return fmt.Errorf("card number is not registered in our system")
	}

	return w.db.Transaction(func(tx *gorm.DB) error {
		//fetch wallet
		userID := ctx.Value(jwt.UserClaimKey).(*jwt.UserClaims).UserID
		walletEntity, walletFetchErr := w.GetUserActiveWallet(ctx, userID)
		if walletFetchErr != nil {
			return walletFetchErr
		}

		//charge wallet
		walletEntity.Credit += walletTopUp.Amount
		if err := tx.Save(&walletEntity).Error; err != nil {
			return err
		}
		//store transaction
		transaction := entities.WalletTransaction{
			WalletID: walletEntity.ID,
			Type:     wallet.TRANSACTION_TYPE_TOPUP,
			Status:   wallet.TRANSACTION_STATUS_DONE, // as we don't have real payment gateway
			Amount:   walletTopUp.Amount,
			Additional: map[string]interface{}{
				"card_number": walletTopUp.CardNumber,
			},
		}
		if err := tx.Save(&transaction).Error; err != nil {
			return err
		}

		return nil
	})
}

func (w *walletRepo) Withdraw(ctx context.Context, walletWithdraw *wallet.WalletWithdraw) error {
	//check for card existence
	var walletCard entities.WalletCard
	if err := w.db.WithContext(ctx).Model(&entities.WalletCard{}).Where("id = ?", walletWithdraw.CardID).Find(&walletCard).Error; err != nil {
		return err
	}

	if walletCard.ID == 0 {
		return fmt.Errorf("card not found")
	}

	//fetch wallet
	userID := ctx.Value(jwt.UserClaimKey).(*jwt.UserClaims).UserID
	walletEntity, walletFetchErr := w.GetUserActiveWallet(ctx, userID)
	if walletFetchErr != nil {
		return walletFetchErr
	}

	if walletEntity.ID != walletCard.WalletID {
		return fmt.Errorf("card is not belongs to you")
	}

	//check for available credits
	if walletWithdraw.Amount > walletEntity.Credit {
		return fmt.Errorf("insufficient credits")
	}

	return w.db.Transaction(func(tx *gorm.DB) error {

		//deduct from wallet
		walletEntity.Credit -= walletWithdraw.Amount
		if err := tx.Save(&walletEntity).Error; err != nil {
			return err
		}

		//store transaction
		transaction := entities.WalletTransaction{
			WalletID: walletEntity.ID,
			Type:     wallet.TRANSACTION_TYPE_WITHDRAW,
			Status:   wallet.TRANSACTION_STATUS_DONE, // as we don't have real payment gateway
			Amount:   walletWithdraw.Amount,
		}
		if err := tx.Save(&transaction).Error; err != nil {
			return err
		}

		return nil
	})
}

func (w *walletRepo) Expense(ctx context.Context, targetWallet *wallet.Wallet, amount float64) error {
	walletEntity := mappers.DomainToWalletEntity(targetWallet)

	return w.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {

		//deduct from wallet
		walletEntity.Credit -= amount
		if err := tx.Save(&walletEntity).Error; err != nil {
			return err
		}

		//store transaction
		transaction := entities.WalletTransaction{
			WalletID: walletEntity.ID,
			Type:     wallet.TRANSACTION_TYPE_EXPENSE,
			Status:   wallet.TRANSACTION_STATUS_DONE,
			Amount:   amount,
		}
		if err := tx.Save(&transaction).Error; err != nil {
			return err
		}

		return nil
	})
}

func (w *walletRepo) GetUserActiveWallet(ctx context.Context, userId uint) (entities.Wallet, error) {
	var walletEntity entities.Wallet
	if err := w.db.WithContext(ctx).Model(&entities.Wallet{}).Where("user_id = ?", userId).Find(&walletEntity).Error; err != nil {
		return entities.Wallet{}, err
	}

	if walletEntity.ID == 0 {
		return entities.Wallet{}, fmt.Errorf("wallet not found for this user")
	}

	return walletEntity, nil
}

func (w *walletRepo) StoreWalletCard(ctx context.Context, card *wallet.WalletCard) error {
	//get wallet id
	userID := ctx.Value(jwt.UserClaimKey).(*jwt.UserClaims).UserID

	var walletEntity entities.Wallet
	if err := w.db.WithContext(ctx).Model(&entities.Wallet{}).Where("user_id = ?", userID).Find(&walletEntity).Error; err != nil {
		return err
	}

	//check if already exists in database?
	var existedCard entities.WalletCard
	if err := w.db.WithContext(ctx).Model(&entities.WalletCard{}).Where("number = ?", card.Number).Find(&existedCard).Error; err != nil {
		return err
	}

	if existedCard.ID != 0 {
		return fmt.Errorf("this card number is already exists")
	}

	walletCardEntity := mappers.DomainToWalletCardEntity(card)

	walletCardEntity.WalletID = walletEntity.ID

	//store in database
	if err := w.db.WithContext(ctx).Create(&walletCardEntity).Error; err != nil {
		return err
	}

	return nil
}

func (w *walletRepo) DeleteWalletCard(ctx context.Context, card *wallet.WalletCard) error {
	//get wallet id
	userID := ctx.Value(jwt.UserClaimKey).(*jwt.UserClaims).UserID

	var walletEntity entities.Wallet
	if err := w.db.WithContext(ctx).Model(&entities.Wallet{}).Where("user_id = ?", userID).Find(&walletEntity).Error; err != nil {
		return err
	}

	//authorize
	if walletEntity.UserID != userID {
		return fmt.Errorf("this wallet card is not belongs to you")
	}

	walletCardEntity := mappers.DomainToWalletCardEntity(card)

	//store in database
	if err := w.db.WithContext(ctx).Delete(&walletCardEntity).Error; err != nil {
		return err
	}

	return nil
}
