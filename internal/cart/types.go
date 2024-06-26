package cart

import (
	"TaamResan/internal/cart_item"
	"context"
)

type Repo interface {
	Delete(ctx context.Context, id uint) error
	GetByUserId(ctx context.Context, userId uint) (*Cart, error)
}

type Cart struct {
	ID           uint
	UserId       uint
	RestaurantId *uint
	Items        []*cart_item.CartItem
}
