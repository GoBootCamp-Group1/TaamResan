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
	ID           uint   `json:"id"`
	ParentId     *uint  `json:"parent_id"`
	RestaurantId uint   `json:"restaurant_id"`
	CreatedBy    uint   `json:"created_by"`
	Name         string `json:"name"`
}
