package mappers

import (
	"TaamResan/internal/address"
	"TaamResan/internal/order"
	"TaamResan/pkg/adapters/storage/entities"
	"gorm.io/gorm"
)

func OrderEntityToDomain(entity *entities.Order) *order.Order {
	var addr *address.Address
	if entity.Address != nil {
		addr = AddressEntityToDomain(entity.Address)
	}

	return &order.Order{
		ID:                 entity.ID,
		RestaurantID:       entity.RestaurantID,
		UserID:             entity.UserID,
		AddressID:          entity.AddressID,
		CreatedAt:          entity.CreatedAt,
		CustomerApprovedAt: entity.CustomerApprovedAt,
		Status:             entity.Status,
		Note:               entity.Note,

		Restaurant: RestaurantEntityToDomain(entity.Restaurant),
		Address:    addr,
	}
}

func DomainToOrderEntity(model *order.Order) *entities.Order {
	return &entities.Order{
		Model:              gorm.Model{ID: model.ID},
		RestaurantID:       model.RestaurantID,
		UserID:             model.UserID,
		AddressID:          model.AddressID,
		CustomerApprovedAt: model.CustomerApprovedAt,
		Status:             model.Status,
		Note:               model.Note,
		Restaurant:         DomainToRestaurantEntity(model.Restaurant),
		Address:            DomainToAddressEntity(model.Address),
	}
}
