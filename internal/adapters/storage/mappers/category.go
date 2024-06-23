package mappers

import (
	"TaamResan/internal/adapters/storage/entities"
	"TaamResan/internal/category"
	"gorm.io/gorm"
)

func CategoryEntityToDomain(entity *entities.Category) *category.Category {
	return &category.Category{
		ID:           entity.ID,
		ParentId:     entity.ParentId,
		RestaurantId: entity.RestaurantId,
		CreatedBy:    entity.CreatedBy,
		Name:         entity.Name,
	}
}

func DomainToCategoryEntity(model *category.Category) *entities.Category {
	return &entities.Category{
		Model:        gorm.Model{ID: model.ID},
		ParentId:     model.ParentId,
		RestaurantId: model.RestaurantId,
		CreatedBy:    model.CreatedBy,
		Name:         model.Name,
	}
}
