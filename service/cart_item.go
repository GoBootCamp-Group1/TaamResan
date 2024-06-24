package service

import (
	"TaamResan/internal/cart_item"
	"context"
)

type CartItemService struct {
	cartItemOps *cart_item.Ops
}

func NewCartItemService(cartItemOps *cart_item.Ops) *CartItemService {
	return &CartItemService{cartItemOps: cartItemOps}
}

func (s *CartItemService) Create(ctx context.Context, cart *cart_item.CartItem) (uint, error) {
	return s.cartItemOps.Create(ctx, cart)
}

func (s *CartItemService) Update(ctx context.Context, cart *cart_item.CartItem) error {
	return s.cartItemOps.Update(ctx, cart)
}

func (s *CartItemService) Delete(ctx context.Context, id uint) error {
	return s.cartItemOps.Delete(ctx, id)
}

func (s *CartItemService) GetAllByCartId(ctx context.Context, cartId uint) ([]*cart_item.CartItem, error) {
	return s.cartItemOps.GetAllByCartId(ctx, cartId)
}
