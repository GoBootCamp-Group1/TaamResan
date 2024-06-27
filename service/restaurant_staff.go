package service

import (
	"TaamResan/internal/restaurant_staff"
	"context"
)

type RestaurantStaffService struct {
	restaurantStaffOps *restaurant_staff.Ops
}

func NewRestaurantStaffService(restaurantStaffOps *restaurant_staff.Ops) *RestaurantStaffService {
	return &RestaurantStaffService{
		restaurantStaffOps: restaurantStaffOps,
	}
}

func (s *RestaurantStaffService) Create(ctx context.Context, res *restaurant_staff.RestaurantStaff) (uint, error) {
	return s.restaurantStaffOps.Create(ctx, res)
}

func (s *RestaurantStaffService) Delete(ctx context.Context, id uint) error {
	return s.restaurantStaffOps.Delete(ctx, id)
}

func (s *RestaurantStaffService) GetAllByRestaurantId(ctx context.Context, restaurantId uint) ([]*restaurant_staff.RestaurantStaff, error) {
	return s.restaurantStaffOps.GetAllByRestaurantId(ctx, restaurantId)
}

func (s *RestaurantStaffService) GetById(ctx context.Context, id uint) (*restaurant_staff.RestaurantStaff, error) {
	return s.restaurantStaffOps.GetById(ctx, id)
}
