package mappers

import (
	"TaamResan/internal/adapters/storage/entities"
	"TaamResan/internal/category"
	"TaamResan/internal/food"
	"gorm.io/gorm"
)

func FoodEntityToDomain(entity *entities.Food) *food.Food {
	var categories []*category.Category

	for _, c := range entity.Categories {
		categories = append(categories, CategoryEntityToDomain(c))
	}

	return &food.Food{
		ID:                 entity.ID,
		RestaurantId:       entity.RestaurantId,
		CreatedBy:          entity.CreatedBy,
		Name:               entity.Name,
		Price:              entity.Price,
		CancelRate:         entity.CancelRate,
		PreparationMinutes: entity.PreparationMinutes,
		Restaurant:         RestaurantEntityToDomain(entity.Restaurant),
		Categories:         categories,
	}
}

func DomainToFoodEntity(model *food.Food) *entities.Food {
	return &entities.Food{
		Model:              gorm.Model{ID: model.ID},
		RestaurantId:       model.RestaurantId,
		CreatedBy:          model.CreatedBy,
		Name:               model.Name,
		Price:              model.Price,
		CancelRate:         model.CancelRate,
		PreparationMinutes: model.PreparationMinutes,
	}
}
