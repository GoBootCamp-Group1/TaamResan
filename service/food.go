package service

import (
	"TaamResan/internal/food"
	"context"
)

type FoodService struct {
	foodOps *food.Ops
}

func NewFoodService(foodOps *food.Ops) *FoodService {
	return &FoodService{foodOps: foodOps}
}

func (s *FoodService) Create(ctx context.Context, food *food.Food) (uint, error) {
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
