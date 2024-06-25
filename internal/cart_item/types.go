package cart_item

import "context"

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
	Amount uint
	Note   string
}
