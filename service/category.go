package service

import (
	"TaamResan/internal/category"
	"context"
)

type CategoryService struct {
	categoryOps *category.Ops
}

func NewCategoryService(categoryOps *category.Ops) *CategoryService {
	return &CategoryService{categoryOps: categoryOps}
}

func (s *CategoryService) Create(ctx context.Context, category *category.Category) (uint, error) {
	return s.categoryOps.Create(ctx, category)
}

func (s *CategoryService) Update(ctx context.Context, category *category.Category) error {
	return s.categoryOps.Update(ctx, category)
}

func (s *CategoryService) Delete(ctx context.Context, id uint) error {
	return s.categoryOps.Delete(ctx, id)
}

func (s *CategoryService) GetById(ctx context.Context, id uint) (*category.Category, error) {
	return s.categoryOps.GetById(ctx, id)
}

func (s *CategoryService) GetAll(ctx context.Context, restaurantId uint) ([]*category.Category, error) {
	return s.categoryOps.GetAll(ctx, restaurantId)
}
