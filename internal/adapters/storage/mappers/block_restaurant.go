package mappers

import (
	"TaamResan/internal/adapters/storage/entities"
	"TaamResan/internal/block_restaurant"
	"gorm.io/gorm"
)

func BlockRestaurantEntityToDomain(entity *entities.BlockRestaurant) *block_restaurant.BlockRestaurant {
	return &block_restaurant.BlockRestaurant{
		ID:           entity.ID,
		UserId:       entity.UserId,
		RestaurantId: entity.RestaurantId,
	}
}

func DomainToBlockRestaurantEntity(model *block_restaurant.BlockRestaurant) *entities.BlockRestaurant {
	return &entities.BlockRestaurant{
		Model:        gorm.Model{ID: model.ID},
		UserId:       model.UserId,
		RestaurantId: model.RestaurantId,
	}
}
