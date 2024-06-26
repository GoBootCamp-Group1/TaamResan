package service

import (
	"TaamResan/internal/food"
	"TaamResan/internal/restaurant"
	"context"
)

type SearchService struct {
	RestaurantRepo *restaurant.Ops
	FoodRepo       *food.Ops
}

func NewSearchService(restaurantRepo *restaurant.Ops, foodRepo *food.Ops) *SearchService {
	return &SearchService{
		RestaurantRepo: restaurantRepo,
		FoodRepo:       foodRepo,
	}
}

func (s *SearchService) SearchRestaurants(ctx context.Context, name string, categoryID uint64, lat float64, lng float64) ([]*restaurant.Restaurant, error) {
	return s.RestaurantRepo.SearchRestaurants(ctx, name, categoryID, lat, lng)
}

func (s *SearchService) SearchFoods(ctx context.Context, name string, categoryID uint64, lat float64, lng float64) ([]*food.Food, error) {
	return s.FoodRepo.SearchFoods(ctx, name, categoryID, lat, lng)
}
