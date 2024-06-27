package block_restaurant

import "context"

type Repo interface {
	Create(ctx context.Context, br *BlockRestaurant) (uint, error)
	Delete(ctx context.Context, id uint) error
	GetAllByUserId(ctx context.Context, userId uint) ([]*BlockRestaurant, error)
}

type BlockRestaurant struct {
	ID           uint
	UserId       uint
	RestaurantId uint
}
