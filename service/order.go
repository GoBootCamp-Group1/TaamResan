package service

import (
	"TaamResan/internal/cart"
	"TaamResan/internal/food"
	"TaamResan/internal/order"
	"TaamResan/internal/wallet"
	"TaamResan/pkg/jwt"
	"context"
	"errors"
	"fmt"
)

type OrderService struct {
	orderOps  *order.Ops
	cartOps   *cart.Ops
	foodOps   *food.Ops
	walletOps *wallet.Ops
}

func NewOrderService(orderOps *order.Ops, cartOps *cart.Ops, foodOps *food.Ops, walletOps *wallet.Ops) *OrderService {
	return &OrderService{
		orderOps:  orderOps,
		cartOps:   cartOps,
		foodOps:   foodOps,
		walletOps: walletOps,
	}
}

var (
	ErrNoCartItem          = errors.New("there is no items in cart")
	ErrInsufficientCredits = errors.New("insufficient credits in your wallet, please top-up to make order")
)

func (s *OrderService) Create(ctx context.Context, data *order.InputData) (*order.Order, error) {
	//TODO: begin transaction

	//user id
	userID := ctx.Value(jwt.UserClaimKey).(*jwt.UserClaims).UserID
	//get cart items
	cartModel, err := s.cartOps.GetById(ctx, data.CartID)
	if err != nil {
		return nil, fmt.Errorf("can not fetch user cart")
	}

	//CONDITIONS
	//check if cart has items -> TODO: extract
	if len(cartModel.Items) == 0 {
		return nil, ErrNoCartItem
	}

	//calculate total amount
	totalAmount := cartModel.CalculateItemsAmount()

	//TODO: check if foods price is equal to cart item price

	//check if wallet has enough credit -> TODO: extract
	userWallet, err := s.walletOps.GetWalletByUserId(ctx, userID)

	if err != nil {
		return nil, err
	}

	if userWallet.Credit < totalAmount {
		return nil, ErrInsufficientCredits
	}

	//create an order
	newOrder, err := s.orderOps.Create(ctx, data, cartModel)
	if err != nil {
		return nil, err
	}

	err = s.walletOps.Expense(ctx, userWallet, totalAmount)

	if err != nil {
		return nil, err
	}

	return newOrder, nil
}

func (s *OrderService) CancelByCustomer(ctx context.Context, o *order.Order) (*order.Order, float64, error) {
	//TODO:start transaction
	//check if order is already cancelled
	orderInfo, err := s.orderOps.GetOrderByID(ctx, o.ID)
	if orderInfo.Status == order.STATUS_CANCELLED_BY_CUSTOMER {
		return nil, 0, fmt.Errorf("order is already cancelled by customer")
	}
	//update status of order
	orderInfo.Status = order.STATUS_CANCELLED_BY_CUSTOMER
	updatedOrder, err := s.orderOps.Update(ctx, orderInfo)
	if err != nil {
		return nil, 0, fmt.Errorf("can not update order: %w", err)
	}
	//calculate food cancellation price
	cancellationAmount, err := s.orderOps.GetItemsCancellationFee(ctx, o)
	if err != nil {
		return nil, 0, fmt.Errorf("can not get order cancellation amount: %w", err)
	}
	totalAmount, err := s.orderOps.GetItemsFee(ctx, o)
	if err != nil {
		return nil, 0, fmt.Errorf("can not get order items fee: %w", err)
	}
	//add to wallet
	userID := ctx.Value(jwt.UserClaimKey).(*jwt.UserClaims).UserID
	userWallet, err := s.walletOps.GetWalletByUserId(ctx, userID)
	refundAmount := totalAmount - cancellationAmount
	err = s.walletOps.Refund(ctx, userWallet, refundAmount)
	if err != nil {
		return nil, 0, fmt.Errorf("can not refund user wallet: %w", err)
	}
	//TODO:commit transaction
	return updatedOrder, refundAmount, nil
}
