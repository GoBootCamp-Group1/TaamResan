package storage

import (
	"TaamResan/internal/category_food"
	"context"
	"gorm.io/gorm"
)

type categoryFoodRepo struct {
	db *gorm.DB
}

func NewCategoryFoodRepo(db *gorm.DB) category_food.Repo { return &categoryFoodRepo{db: db} }

func (r *categoryFoodRepo) Create(ctx context.Context, categoryFood *category_food.CategoryFood) (uint, error) {
	panic("implement me")
}

func (r *categoryFoodRepo) Update(ctx context.Context, categoryFood *category_food.CategoryFood) error {
	panic("implement me")
}

func (r *categoryFoodRepo) Delete(ctx context.Context, id uint) error {
	panic("implement me")
}

func (r *categoryFoodRepo) GetById(ctx context.Context, id uint) (*category_food.CategoryFood, error) {
	panic("implement me")
}

func (r *categoryFoodRepo) GetAll(ctx context.Context, restaurantId uint) ([]*category_food.CategoryFood, error) {
	panic("implement me")
}
