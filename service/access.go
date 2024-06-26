package service

import (
	"TaamResan/internal/access"
	"context"
)

type AccessService struct {
	accessOps *access.Ops
}

func NewAccessService(accessOps *access.Ops) *AccessService {
	return &AccessService{
		accessOps: accessOps,
	}
}

func (s *AccessService) CheckRestaurantOwner(ctx context.Context, userId uint, restaurantId uint) error {
	return s.accessOps.CheckRestaurantOwner(ctx, userId, restaurantId)
}

func (s *AccessService) CheckRestaurantStaff(ctx context.Context, userId uint, restaurantId uint) error {
	return s.accessOps.CheckRestaurantStaff(ctx, userId, restaurantId)
}

func (s *AccessService) CheckAdminAccess(ctx context.Context, userId uint) error {
	return s.accessOps.CheckAdminAccess(ctx, userId)
}
