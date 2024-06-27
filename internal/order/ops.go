package order

import (
	"TaamResan/internal/cart"
	"TaamResan/internal/cart_item"
	"context"
)

type Ops struct {
	repo Repo
}

func NewOps(repo Repo) *Ops { return &Ops{repo: repo} }

func (o Ops) Create(ctx context.Context, data *InputData, cartModel *cart.Cart) (*Order, error) {
	return o.repo.Create(ctx, data, cartModel)
}

func (o Ops) AddCartItemToOrder(ctx context.Context, order *Order, item *cart_item.CartItem) error {
	return o.repo.AddCartItemToOrder(ctx, order, item)
}

func (o Ops) Update(ctx context.Context, order *Order) (*Order, error) {
	return o.repo.Update(ctx, order)
}

func (o Ops) GetItemsCancellationFee(ctx context.Context, order *Order) (float64, error) {
	return o.repo.GetItemsCancellationFee(ctx, order)
}

func (o Ops) GetItemsFee(ctx context.Context, order *Order) (float64, error) {
	return o.repo.GetItemsFee(ctx, order)
}

func (o Ops) GetOrderByID(ctx context.Context, id uint) (*Order, error) {
	return o.repo.GetOrderByID(ctx, id)
}

func (o Ops) ChangeStatusByRestaurant(ctx context.Context, order *Order) error {
	return o.repo.ChangeStatusByRestaurant(ctx, order)
}
