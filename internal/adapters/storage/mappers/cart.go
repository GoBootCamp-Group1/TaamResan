package mappers

import (
	"TaamResan/internal/adapters/storage/entities"
	"TaamResan/internal/cart"
	"gorm.io/gorm"
)

func CartEntityToDomain(entity *entities.Cart) *cart.Cart {
	return &cart.Cart{
		ID:           entity.ID,
		UserId:       entity.UserId,
		RestaurantId: entity.RestaurantId,
	}
}

func DomainToCartEntity(model *cart.Cart) *entities.Cart {
	return &entities.Cart{
		Model:        gorm.Model{ID: model.ID},
		UserId:       model.UserId,
		RestaurantId: model.RestaurantId,
	}
}
