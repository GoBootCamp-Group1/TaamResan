package category_food

import "context"

type Repo interface {
	Create(ctx context.Context, categoryFood *CategoryFood) (uint, error)
	Update(ctx context.Context, categoryFood *CategoryFood) error
	Delete(ctx context.Context, id uint) error
	GetById(ctx context.Context, id uint) (*CategoryFood, error)
	GetAll(ctx context.Context, restaurantId uint) ([]*CategoryFood, error)
}

type CategoryFood struct {
	ID         uint
	CategoryId uint
	FoodId     uint
}
