package restaurant_staff

import "context"

type Repo interface {
	Create(ctx context.Context, rStaff *RestaurantStaff) (uint, error)
	Delete(ctx context.Context, id uint) error
	GetAllByRestaurantId(ctx context.Context, restaurantId uint) ([]*RestaurantStaff, error)
}

type RestaurantStaff struct {
	ID           uint
	UserId       uint
	RestaurantId uint
	Position     uint
}

// Positions
const (
	Manager uint = iota + 1
	Operator
	Unknown
)

const (
	MANAGER  = "manager"
	OPERATOR = "operator"
)

func GetPosition(p string) uint {
	switch p {
	case MANAGER:
		return Manager
	case OPERATOR:
		return Operator
	default:
		return Unknown
	}
}
