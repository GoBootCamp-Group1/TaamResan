package role

import "context"

type Repo interface {
	Create(ctx context.Context, role *Role) error
	Update(ctx context.Context, role *Role) error
	Delete(ctx context.Context, id uint) error
	GetByName(ctx context.Context, name string) (*Role, error)
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
