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
