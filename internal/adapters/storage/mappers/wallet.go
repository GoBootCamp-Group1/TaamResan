package mappers

import (
	"TaamResan/internal/adapters/storage/entities"
	"TaamResan/internal/wallet"
)

func WalletEntityToDomain(entity *entities.Wallet) *wallet.Wallet {
	return &wallet.Wallet{
		ID:     entity.ID,
		Credit: entity.Credit,
	}
}

func DomainToWalletEntity(model *wallet.Wallet) *entities.Wallet {
	return &entities.Wallet{
		ID:     model.ID,
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
