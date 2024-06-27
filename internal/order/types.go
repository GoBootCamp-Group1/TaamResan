package order

import (
	"TaamResan/internal/address"
	"TaamResan/internal/cart"
	"TaamResan/internal/cart_item"
	"TaamResan/internal/restaurant"
	"TaamResan/internal/user"
	"context"
	"time"
)

type Order struct {
	ID                 uint
	RestaurantID       uint
	UserID             uint
	AddressID          uint
	CreatedAt          time.Time
	CustomerApprovedAt time.Time
	Status             uint
	Note               string

	Restaurant *restaurant.Restaurant `json:"restaurant"`
	Address    *address.Address       `json:"address"`
	User       *user.User             `json:"user"`
}

type InputData struct {
	CartID    uint
	AddressID uint
	Note      *string
}

const (
	STATUS_DRAFT uint = iota + 1
	STATUS_UNPAID
	STATUS_PAID
	STATUS_CANCELLED_BY_CUSTOMER
	STATUS_CANCELLED_BY_RESTAURANT
	STATUS_WAITING_RESTAURANT_APPROVE
	STATUS_PREPARING
	STATUS_COURIER_TO_RESTAURANT
	STATUS_COURIER_TO_CUSTOMER
	STATUS_DELIVERED
)

type Repo interface {
	Create(ctx context.Context, data *InputData, cartModel *cart.Cart) (*Order, error)
	AddCartItemToOrder(ctx context.Context, order *Order, item *cart_item.CartItem) error
	Update(ctx context.Context, order *Order) (*Order, error)
	GetItemsCancellationFee(ctx context.Context, order *Order) (float64, error)
	GetItemsFee(ctx context.Context, order *Order) (float64, error)
	GetOrderByID(ctx context.Context, id uint) (*Order, error)
}
