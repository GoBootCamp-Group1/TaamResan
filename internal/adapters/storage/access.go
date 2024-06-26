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
	ErrNotOwner           = errors.New("this user is not owner of the restaurant")
	ErrNotOwnerOrOperator = errors.New("this user is not owner/operator in the restaurant")
	ErrNotAllowed         = errors.New("this user is not allowed to do this action")
)

func (r *accessRepo) CheckRestaurantOwner(ctx context.Context, userId uint, restaurantId uint) error {
	var entity *entities.Restaurant
	err := r.db.WithContext(ctx).Model(&entities.Restaurant{}).Where("id = ? and owned_by = ?", restaurantId, userId).First(&entity).Error
	if err != nil {
		return ErrNotOwner
	}

	return nil
}

func (r *accessRepo) CheckRestaurantStaff(ctx context.Context, userId uint, restaurantId uint) error {
	var entity *entities.RestaurantStaff
	err := r.db.WithContext(ctx).Model(&entities.RestaurantStaff{}).
		Where("restaurant_id = ? and user_id = ?", restaurantId, userId).First(&entity).Error
	if err != nil {
		return ErrNotOwnerOrOperator
	}

	return nil
}

func (r *accessRepo) CheckAdminAccess(ctx context.Context, userId uint) error {
	var entity *entities.User
	err := r.db.WithContext(ctx).Model(&entities.User{}).Where("id = ?", userId).First(&entity).Error
	if err != nil {
		return ErrNotAllowed
	}

	return nil
}
