package category_food

import "context"

type Repo interface {
	Create(ctx context.Context, categoryFood *CategoryFood) (uint, error)
	Delete(ctx context.Context, id uint) error
	GetById(ctx context.Context, id uint) (*CategoryFood, error)
	GetAllByFoodId(ctx context.Context, foodId uint) ([]*CategoryFood, error)
	GetAllByCategoryId(ctx context.Context, categoryId uint) ([]*CategoryFood, error)
}

type CategoryFood struct {
	ID         uint
	CategoryId uint
	FoodId     uint
}
