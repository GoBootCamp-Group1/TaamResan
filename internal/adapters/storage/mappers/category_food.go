package mappers

import (
	"TaamResan/internal/adapters/storage/entities"
	"TaamResan/internal/category_food"
	"gorm.io/gorm"
)

func CategoryFoodEntityToDomain(entity *entities.CategoryFood) *category_food.CategoryFood {
	return &category_food.CategoryFood{
		ID:         entity.ID,
		CategoryId: entity.CategoryId,
		FoodId:     entity.FoodId,
	}
}

func DomainToCategoryFoodEntity(model *category_food.CategoryFood) *entities.CategoryFood {
	return &entities.CategoryFood{
		Model:      gorm.Model{ID: model.ID},
		CategoryId: model.CategoryId,
		FoodId:     model.FoodId,
	}
}
