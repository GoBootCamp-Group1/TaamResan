package storage

import (
	"TaamResan/internal/access"
	"TaamResan/internal/adapters/storage/entities"
	"context"
	"errors"
	"gorm.io/gorm"
)

type accessRepo struct {
	db *gorm.DB
}

func NewAccessRepo(db *gorm.DB) access.Repo {
	return &accessRepo{
		db: db,
	}
}

var (
	ErrNotOwner = errors.New("this user is not owner of the restaurant")
)

func (r *accessRepo) CheckRestaurantOwner(ctx context.Context, userId uint, restaurantId uint) error {
	var entity *entities.Restaurant
	err := r.db.WithContext(ctx).Model(&entities.Restaurant{}).Where("id = ? and owned_by = ?", restaurantId, userId).First(&entity).Error
	if err != nil {
		return ErrNotOwner
	}

	return nil
}