package food

import (
	"TaamResan/internal/category"
	"context"
)

type Repo interface {
	Create(ctx context.Context, food *Food) (uint, error)
	Update(ctx context.Context, food *Food) error
	Delete(ctx context.Context, id uint) error
	GetById(ctx context.Context, id uint) (*Food, error)
	GetAll(ctx context.Context, restaurantId uint) ([]*Food, error)
}

type Food struct {
	ID                 uint
	RestaurantId       uint
	CreatedBy          uint
	Name               string
	Price              float64
	CancelRate         float64
	PreparationMinutes uint
	Categories         []*category.Category
}
