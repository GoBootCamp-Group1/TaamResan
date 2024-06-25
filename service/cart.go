package service

import (
	"TaamResan/internal/cart"
	"context"
)

type CartService struct {
	cartOps *cart.Ops
}

func NewCartService(cartOps *cart.Ops) *CartService {
	return &CartService{cartOps: cartOps}
}

func (s *CartService) Delete(ctx context.Context, id uint) error {
	return s.cartOps.Delete(ctx, id)
}
func (s *CartService) GetByUserId(ctx context.Context, userId uint) (*cart.Cart, error) {
	return s.cartOps.GetByUserId(ctx, userId)
}
