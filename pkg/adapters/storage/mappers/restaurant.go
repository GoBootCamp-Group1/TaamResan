package mappers

import (
	"TaamResan/internal/address"
	"TaamResan/internal/restaurant"
	"TaamResan/pkg/adapters/storage/entities"
	"gorm.io/gorm"
)

func RestaurantEntityToDomain(entity *entities.Restaurant) *restaurant.Restaurant {
	return &restaurant.Restaurant{
		ID:             entity.ID,
		Name:           entity.Name,
		OwnedBy:        entity.OwnedBy,
		ApprovalStatus: entity.ApprovalStatus,
		Address:        address.Address{ID: entity.AddressId},
		CourierSpeed:   entity.CourierSpeed,
	}
}

func DomainToRestaurantEntity(model *restaurant.Restaurant) *entities.Restaurant {
	return &entities.Restaurant{
		Model:          gorm.Model{ID: model.ID},
		Name:           model.Name,
		OwnedBy:        model.OwnedBy,
		ApprovalStatus: model.ApprovalStatus,
		AddressId:      model.Address.ID,
		CourierSpeed:   model.CourierSpeed,
	}
}
