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
