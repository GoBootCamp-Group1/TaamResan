package service

import (
	"TaamResan/internal/action_log"
	"TaamResan/internal/restaurant"
	"context"
)

type RestaurantService struct {
	restaurantOps *restaurant.Ops
	logOps        *action_log.Ops
}

func NewRestaurantService(restaurantOps *restaurant.Ops, logOps *action_log.Ops) *RestaurantService {
	return &RestaurantService{
		restaurantOps: restaurantOps,
		logOps:        logOps,
	}
}

func (s *RestaurantService) Create(ctx context.Context, res *restaurant.Restaurant) (uint, error) {
	id, err := s.restaurantOps.Create(ctx, res)
	if err != nil {
		return 0, err
	}

	if err = s.updateLog(ctx, id); err != nil {
		return 0, err
	}

	return id, nil
}

func (s *RestaurantService) Update(ctx context.Context, res *restaurant.Restaurant) error {
	err := s.restaurantOps.Update(ctx, res)
	if err != nil {
		return err
	}

	if err = s.updateLog(ctx, res.ID); err != nil {
		return err
	}

	return nil
}

func (s *RestaurantService) Delete(ctx context.Context, id uint) error {
	err := s.restaurantOps.Delete(ctx, id)
	if err != nil {
		return err
	}

	if err = s.updateLog(ctx, id); err != nil {
		return err
	}

	return nil
}

func (s *RestaurantService) GetById(ctx context.Context, id uint) (*restaurant.Restaurant, error) {
	res, err := s.restaurantOps.GetById(ctx, id)
	if err != nil {
		return nil, err
	}

	if err = s.updateLog(ctx, id); err != nil {
		return nil, err
	}

	return res, nil
}

func (s *RestaurantService) GetAll(ctx context.Context) ([]*restaurant.Restaurant, error) {
	return s.restaurantOps.GetAll(ctx)
}

func (s *RestaurantService) Approve(ctx context.Context, id uint) error {
	err := s.restaurantOps.Approve(ctx, id)
	if err != nil {
		return err
	}

	if err = s.updateLog(ctx, id); err != nil {
		return err
	}

	return nil
}

func (s *RestaurantService) DelegateOwnership(ctx context.Context, id uint, newOwnerId uint) error {
	err := s.restaurantOps.DelegateOwnership(ctx, id, newOwnerId)
	if err != nil {
		return err
	}

	if err = s.updateLog(ctx, id); err != nil {
		return err
	}

	return nil
}

func (s *RestaurantService) updateLog(ctx context.Context, id uint) error {
	log := ctx.Value(action_log.LogCtxKey).(*action_log.ActionLog)
	log.EntityType = action_log.RestaurantEntityType
	log.EntityID = id

	if err := s.logOps.Update(ctx, log); err != nil {
		return err
	}

	return nil
}
