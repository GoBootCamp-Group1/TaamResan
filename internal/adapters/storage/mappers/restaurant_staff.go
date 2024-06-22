package mappers

import (
	"TaamResan/internal/adapters/storage/entities"
	"TaamResan/internal/restaurant_staff"
	"gorm.io/gorm"
)

func RestaurantStaffEntityToDomain(entity *entities.RestaurantStaff) *restaurant_staff.RestaurantStaff {
	return &restaurant_staff.RestaurantStaff{
		ID:           entity.ID,
		UserId:       entity.UserId,
		RestaurantId: entity.RestaurantId,
		Position:     entity.Position,
	}
}

func DomainToRestaurantStaffEntity(model *restaurant_staff.RestaurantStaff) *entities.RestaurantStaff {
	return &entities.RestaurantStaff{
		Model:        gorm.Model{ID: model.ID},
		UserId:       model.UserId,
		RestaurantId: model.RestaurantId,
		Position:     model.Position,
	}
}
