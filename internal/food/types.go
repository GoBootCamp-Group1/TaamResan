package food

import (
	"TaamResan/internal/category"
	"TaamResan/internal/restaurant"
	"context"
)

type Repo interface {
	Create(ctx context.Context, food *Food) (uint, error)
	Update(ctx context.Context, food *Food) error
	Delete(ctx context.Context, id uint) error
	GetById(ctx context.Context, id uint) (*Food, error)
	GetAll(ctx context.Context, restaurantId uint) ([]*Food, error)
	SearchFoods(ctx context.Context, searchData *FoodSearch) ([]*Food, error)
}

type Food struct {
	ID                 uint                   `json:"id"`
	RestaurantId       uint                   `json:"restaurant_id"`
	CreatedBy          uint                   `json:"createdBy"`
	Name               string                 `json:"name"`
	Price              float64                `json:"price"`
	CancelRate         float64                `json:"cancel_rate"`
	PreparationMinutes uint                   `json:"preparation_minutes"`
	Restaurant         *restaurant.Restaurant `json:"restaurant"`
	Categories         []*category.Category   `json:"categories"`
}

type FoodSearch struct {
	Name       string
	CategoryID *uint
	Lat        *float64
	Lng        *float64
}
