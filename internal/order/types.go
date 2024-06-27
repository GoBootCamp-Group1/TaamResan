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
	ID                 uint      `json:"id"`
	RestaurantID       uint      `json:"restaurant_id"`
	UserID             uint      `json:"user_id"`
	AddressID          uint      `json:"address_id"`
	CreatedAt          time.Time `json:"created_at"`
	CustomerApprovedAt time.Time `json:"customer_approved_at"`
	Status             uint      `json:"status"`
	StatusTitle        string    `json:"status_title"`
	Note               string    `json:"note"`

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

func (o *Order) MapStatusToStr() string {
	switch o.Status {
	case STATUS_DRAFT:
		return "Draft"
	case STATUS_UNPAID:
		return "Unpaid"
	case STATUS_PAID:
		return "Paid"
	case STATUS_CANCELLED_BY_CUSTOMER:
		return "Cancelled by Customer"
	case STATUS_CANCELLED_BY_RESTAURANT:
		return "Cancelled by Restaurant"
	case STATUS_WAITING_RESTAURANT_APPROVE:
		return "Waiting for Restaurant Approval"
	case STATUS_PREPARING:
		return "Preparing"
	case STATUS_COURIER_TO_RESTAURANT:
		return "Courier to Restaurant"
	case STATUS_COURIER_TO_CUSTOMER:
		return "Courier to Customer"
	case STATUS_DELIVERED:
		return "Delivered"
	default:
		return "Unknown Status"
	}
}

type Repo interface {
	Create(ctx context.Context, data *InputData, cartModel *cart.Cart) (*Order, error)
	AddCartItemToOrder(ctx context.Context, order *Order, item *cart_item.CartItem) error
	Update(ctx context.Context, order *Order) (*Order, error)
	GetItemsCancellationFee(ctx context.Context, order *Order) (float64, error)
	GetItemsFee(ctx context.Context, order *Order) (float64, error)
	GetOrderByID(ctx context.Context, id uint) (*Order, error)
}
