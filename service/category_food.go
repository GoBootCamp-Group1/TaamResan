package service

import (
	"TaamResan/internal/category_food"
	"context"
)

type CategoryFoodService struct {
	categoryFoodOps *category_food.Ops
}

func NewCategoryFoodService(categoryFoodOps *category_food.Ops) *CategoryFoodService {
	return &CategoryFoodService{categoryFoodOps: categoryFoodOps}
}

func (s *CategoryFoodService) Create(ctx context.Context, categoryFood *category_food.CategoryFood) (uint, error) {
	return s.categoryFoodOps.Create(ctx, categoryFood)
}

func (s *CategoryFoodService) Delete(ctx context.Context, id uint) error {
	return s.categoryFoodOps.Delete(ctx, id)
}

func (s *CategoryFoodService) GetById(ctx context.Context, id uint) (*category_food.CategoryFood, error) {
	return s.categoryFoodOps.GetById(ctx, id)
}

func (s *CategoryFoodService) GetAllByFoodId(ctx context.Context, foodId uint) ([]*category_food.CategoryFood, error) {
	return s.categoryFoodOps.GetAllByFoodId(ctx, foodId)
}

func (s *CategoryFoodService) GetAllByCategoryId(ctx context.Context, categoryId uint) ([]*category_food.CategoryFood, error) {
	return s.categoryFoodOps.GetAllByCategoryId(ctx, categoryId)
}
