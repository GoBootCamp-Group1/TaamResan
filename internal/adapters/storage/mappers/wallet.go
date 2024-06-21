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
