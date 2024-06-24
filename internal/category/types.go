package category

import "context"

type Repo interface {
	Create(ctx context.Context, category *Category) (uint, error)
	Update(ctx context.Context, category *Category) error
	Delete(ctx context.Context, id uint) error
	GetById(ctx context.Context, id uint) (*Category, error)
	GetByName(ctx context.Context, name string, restaurantId uint) (*Category, error)
	GetAll(ctx context.Context, restaurantId uint) ([]*Category, error)
}

type Category struct {
	ID           uint
	ParentId     *uint
	RestaurantId uint
	CreatedBy    uint
	Name         string
}
