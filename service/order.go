package service

import (
	"TaamResan/internal/cart"
	"TaamResan/internal/food"
	"TaamResan/internal/order"
	"context"
)

type OrderService struct {
	orderOps *order.Ops
	cartOps  *cart.Ops
	foodOps  *food.Ops
}

func NewOrderService(orderOps *order.Ops, cartOps *cart.Ops, foodOps *food.Ops) *OrderService {
	return &OrderService{
		orderOps: orderOps,
		cartOps:  cartOps,
		foodOps:  foodOps,
	}
}

func (s *OrderService) Create(ctx context.Context, data *order.InputData) (*order.Order, error) {
	return s.orderOps.Create(ctx, data)
}
