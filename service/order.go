package service

import (
	"TaamResan/internal/cart"
	"TaamResan/internal/food"
	"TaamResan/internal/order"
	"TaamResan/pkg/jwt"
	"context"
	"fmt"
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
	//user id
	userID := ctx.Value(jwt.UserClaimKey).(*jwt.UserClaims).UserID
	//get cart items
	userCart, err := s.cartOps.GetByUserId(ctx, userID)
	if err != nil {
		return nil, fmt.Errorf("can not fetch user cart")
	}

	//create an draft order
	draftOrder, err := s.orderOps.Create(ctx, data)

	if len(userCart.Items) > 0 {
		for _, item := range userCart.Items {
			//add into order item
		}
	}

	return s.orderOps.Create(ctx, data)
}
