package mappers

import (
	"TaamResan/internal/cart"
	"TaamResan/internal/cart_item"
	"TaamResan/pkg/adapters/storage/entities"
	"gorm.io/gorm"
)

func CartEntityToDomain(entity *entities.Cart) *cart.Cart {

	var items []*cart_item.CartItem
	for _, i := range entity.Items {
		items = append(items, CartItemEntityToDomain(i))
	}

	return &cart.Cart{
		ID:           entity.ID,
		UserId:       entity.UserId,
		RestaurantId: entity.RestaurantId,
		Items:        items,
	}
}

func DomainToCartEntity(model *cart.Cart) *entities.Cart {
	return &entities.Cart{
		Model:        gorm.Model{ID: model.ID},
		UserId:       model.UserId,
		RestaurantId: model.RestaurantId,
	}
}
