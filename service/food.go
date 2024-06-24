package service

import (
	"TaamResan/internal/category"
	"TaamResan/internal/food"
	"context"
)

type FoodService struct {
	foodOps     *food.Ops
	categoryOps *category.Ops
}

func NewFoodService(foodOps *food.Ops, categoryOps *category.Ops) *FoodService {
	return &FoodService{
		foodOps:     foodOps,
		categoryOps: categoryOps,
	}
}

func (s *FoodService) Create(ctx context.Context, food *food.Food) (uint, error) {
	if len(food.Categories) > 0 {
		for i, c := range food.Categories {
			loadedCategory, err := s.categoryOps.GetByName(ctx, c.Name, food.RestaurantId)
			if err != nil {
				return 0, err
			}
			food.Categories[i].ID = loadedCategory.ID
		}
	}
	return s.foodOps.Create(ctx, food)
}

func (s *FoodService) Update(ctx context.Context, food *food.Food) error {
	return s.foodOps.Update(ctx, food)
}

func (s *FoodService) Delete(ctx context.Context, id uint) error {
	return s.foodOps.Delete(ctx, id)
}

func (s *FoodService) GetById(ctx context.Context, id uint) (*food.Food, error) {
	return s.foodOps.GetById(ctx, id)
}

func (s *FoodService) GetAll(ctx context.Context, restaurantId uint) ([]*food.Food, error) {
	return s.foodOps.GetAll(ctx, restaurantId)
}
