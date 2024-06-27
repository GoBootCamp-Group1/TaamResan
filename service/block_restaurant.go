package service

import (
	"TaamResan/internal/block_restaurant"
	"TaamResan/internal/restaurant"
	"context"
)

type BlockRestaurantService struct {
	blockRestaurantOps *block_restaurant.Ops
	restaurantOps      *restaurant.Ops
}

func NewBlockRestaurantService(blockRestaurantOps *block_restaurant.Ops, restaurantOps *restaurant.Ops) *BlockRestaurantService {
	return &BlockRestaurantService{
		blockRestaurantOps: blockRestaurantOps,
		restaurantOps:      restaurantOps,
	}
}

func (s *BlockRestaurantService) Create(ctx context.Context, br *block_restaurant.BlockRestaurant) (uint, error) {
	// check restaurant existence
	if _, err := s.restaurantOps.GetById(ctx, br.RestaurantId); err != nil {
		return 0, err
	}

	return s.blockRestaurantOps.Create(ctx, br)
}

func (s *BlockRestaurantService) Delete(ctx context.Context, id uint) error {
	return s.blockRestaurantOps.Delete(ctx, id)
}

func (s *BlockRestaurantService) GetAllByUserId(ctx context.Context, userId uint) ([]*block_restaurant.BlockRestaurant, error) {
	return s.blockRestaurantOps.GetAllByUserId(ctx, userId)
}
