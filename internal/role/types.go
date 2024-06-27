package role

import "context"

type Repo interface {
	Create(ctx context.Context, role *Role) error
	Update(ctx context.Context, role *Role) error
	Delete(ctx context.Context, id uint) error
	Get(ctx context.Context, id uint) (*Role, error)
	GetAll(ctx context.Context) ([]*Role, error)
}

type Role struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
}

const (
	CUSTOMER            = "customer"
	ADMIN               = "admin"
	RESTAURANT_OWNER    = "restaurant owner"
	RESTAURANT_OPERATOR = "restaurant operator"
	UNKNOWN             = "unknown"
)

func (ur Role) String() string {
	return ur.Name
}

const (
	Customer uint = iota + 1
	Admin
	RestaurantOwner
	RestaurantOperator
)

var DefaultRole = Role{ID: Customer, Name: CUSTOMER}

const (
	LOG          = "log"
	ROLE         = "manage role"
	ADDRESS      = "manage address"
	CART         = "add to cart"
	ORDER        = "order"
	ORDER_STATUS = "order status"
	WALLET       = "manage wallet"
	RESTAURANT   = "manage restaurant"
	CATEGORY     = "manage category"
	FOOD         = "manage food"
)

var RolePermissions = map[uint]([]string){
	Customer:           {ADDRESS, CART, ORDER, WALLET},
	Admin:              {LOG, ROLE, ADDRESS},
	RestaurantOwner:    {RESTAURANT, CATEGORY, FOOD, ORDER_STATUS},
	RestaurantOperator: {CATEGORY, FOOD, ORDER_STATUS},
}
