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
