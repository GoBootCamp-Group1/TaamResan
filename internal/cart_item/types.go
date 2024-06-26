package cart_item

import (
	"TaamResan/internal/food"
	"context"
)

type Repo interface {
	Create(ctx context.Context, cartItem *CartItem) (uint, error)
	Update(ctx context.Context, cartItem *CartItem) error
	Delete(ctx context.Context, id uint) error
	GetAllByCartId(ctx context.Context, cartId uint) ([]*CartItem, error)
}

type CartItem struct {
	ID     uint
	CartId uint
	FoodId uint
	Amount float64
	Note   string
	Food   *food.Food
}
