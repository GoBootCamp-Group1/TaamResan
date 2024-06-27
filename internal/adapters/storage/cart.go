package storage

import (
	"TaamResan/internal/adapters/storage/entities"
	"TaamResan/internal/adapters/storage/mappers"
	"TaamResan/internal/cart"
	"context"
	"errors"
	"gorm.io/gorm"
)

type cartRepo struct {
	db *gorm.DB
}

var (
	ErrCartNotFound = errors.New("error cart not found")
	ErrDeletingCart = errors.New("error deleting cart")
)

func NewCartRepo(db *gorm.DB) cart.Repo {
	return &cartRepo{db: db}
}

func (r *cartRepo) Delete(ctx context.Context, id uint) error {
	return r.db.Transaction(func(tx *gorm.DB) error {
		// check existence
		var cartEntity *entities.Cart
		if err := r.db.WithContext(ctx).Model(&entities.Cart{}).Where("id = ?", id).Find(&cartEntity).Error; err != nil {
			return ErrCartNotFound
		}

		// delete items
		var cartItems *entities.CartItem
		if err := r.db.WithContext(ctx).Model(&entities.CartItem{}).Where("cart_id = ?", id).Delete(&cartItems).Error; err != nil {
			return ErrDeletingCart
		}

		// delete it
		if err := r.db.WithContext(ctx).Model(&entities.Cart{}).Where("id = ?", id).Delete(&cartEntity).Error; err != nil {
			return ErrDeletingCart
		}

		return nil
	})
}

func (r *cartRepo) GetByUserId(ctx context.Context, userId uint) (*cart.Cart, error) {
	var cartEntity *entities.Cart
	if err := r.db.WithContext(ctx).Model(&entities.Cart{}).
		Preload("Items.Food").
		Where("user_id = ?", userId).Find(&cartEntity).Error; err != nil {
		return nil, ErrCartNotFound
	}
	return mappers.CartEntityToDomain(cartEntity), nil
}

func (r *cartRepo) GetById(ctx context.Context, cartId uint) (*cart.Cart, error) {
	var cartEntity *entities.Cart
	if err := r.db.WithContext(ctx).Model(&entities.Cart{}).
		Preload("Items.Food").
		Where("id = ?", cartId).Find(&cartEntity).Error; err != nil {
		return nil, ErrCartNotFound
	}
	return mappers.CartEntityToDomain(cartEntity), nil
}
