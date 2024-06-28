package mappers

import (
	"TaamResan/internal/wallet"
	"TaamResan/pkg/adapters/storage/entities"
)

func WalletEntityToDomain(entity *entities.Wallet) *wallet.Wallet {
	return &wallet.Wallet{
		ID:     entity.ID,
		UserID: entity.UserID,
		Credit: entity.Credit,
	}
}

func DomainToWalletEntity(model *wallet.Wallet) *entities.Wallet {
	return &entities.Wallet{
		ID:     model.ID,
		UserID: model.UserID,
		Credit: model.Credit,
	}
}

func WalletCardEntityToDomain(entity *entities.WalletCard) *wallet.WalletCard {
	return &wallet.WalletCard{
		ID:       entity.ID,
		Title:    entity.Title,
		BankName: entity.BankName,
		Number:   entity.Number,
	}
}

func DomainToWalletCardEntity(model *wallet.WalletCard) *entities.WalletCard {
	return &entities.WalletCard{
		ID:       model.ID,
		Title:    model.Title,
		BankName: model.BankName,
		Number:   model.Number,
	}
}
