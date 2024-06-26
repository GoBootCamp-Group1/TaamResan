package mappers

import (
	"TaamResan/internal/adapters/storage/entities"
	"TaamResan/internal/cart_item"
	"gorm.io/gorm"
)

func CartItemEntityToDomain(entity *entities.CartItem) *cart_item.CartItem {
	return &cart_item.CartItem{
		ID:     entity.ID,
		CartId: entity.CartId,
		FoodId: entity.FoodId,
		Amount: entity.Amount,
		Note:   entity.Note,
	}
}

func DomainToCartItemEntity(model *cart_item.CartItem) *entities.CartItem {

	return &entities.CartItem{
		Model:  gorm.Model{ID: model.ID},
		CartId: model.CartId,
		FoodId: model.FoodId,
		Amount: model.Amount,
		Note:   model.Note,
	}
}
