package service

import (
	"TaamResan/internal/restaurant"
	"context"
)

type RestaurantService struct {
	restaurantOps *restaurant.Ops
}

func NewRestaurantService(restaurantOps *restaurant.Ops) *RestaurantService {
	return &RestaurantService{
		restaurantOps: restaurantOps,
	}
}

func (s *RestaurantService) Create(ctx context.Context, res *restaurant.Restaurant) (uint, error) {
	return s.restaurantOps.Create(ctx, res)
}

func (s *RestaurantService) Update(ctx context.Context, res *restaurant.Restaurant) error {
	return s.restaurantOps.Update(ctx, res)
}

func (s *RestaurantService) Delete(ctx context.Context, id uint) error {
	return s.restaurantOps.Delete(ctx, id)
}

func (s *RestaurantService) GetById(ctx context.Context, id uint) (*restaurant.Restaurant, error) {
	return s.restaurantOps.GetById(ctx, id)
}

func (s *RestaurantService) GetAll(ctx context.Context) ([]*restaurant.Restaurant, error) {
	return s.restaurantOps.GetAll(ctx)
}

func (s *RestaurantService) Approve(ctx context.Context, id uint) error {
	return s.restaurantOps.Approve(ctx, id)
}

func (s *RestaurantService) DelegateOwnership(ctx context.Context, id uint, newOwnerId uint) error {
	return s.restaurantOps.DelegateOwnership(ctx, id, newOwnerId)
}
