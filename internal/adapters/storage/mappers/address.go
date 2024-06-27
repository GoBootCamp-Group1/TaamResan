package mappers

import (
	"TaamResan/internal/adapters/storage/entities"
	"TaamResan/internal/address"
)

func AddressEntityToDomain(entity *entities.Address) *address.Address {
	return &address.Address{
		ID:    entity.ID,
		Title: entity.Title,
		Lat:   entity.Lat,
		Lng:   entity.Lng,
	}
}

func DomainToAddressEntity(model *address.Address) *entities.Address {
	return &entities.Address{
		ID:    model.ID,
		Title: model.Title,
		Lat:   model.Lat,
		Lng:   model.Lng,
	}
}
