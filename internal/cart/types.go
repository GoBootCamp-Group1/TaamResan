package cart

import "context"

type Repo interface {
	Delete(ctx context.Context, id uint) error
	GetByUserId(ctx context.Context, userId uint) (*Cart, error)
}

type Cart struct {
	ID           uint
	UserId       uint
	RestaurantId *uint
}
