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

func (s *SearchService) SearchRestaurants(ctx context.Context, searchData *restaurant.RestaurantSearch) ([]*restaurant.Restaurant, error) {
	return s.RestaurantRepo.SearchRestaurants(ctx, searchData)
}

func (s *SearchService) SearchFoods(ctx context.Context, searchData *food.FoodSearch) ([]*food.Food, error) {
	return s.FoodRepo.SearchFoods(ctx, searchData)
}
