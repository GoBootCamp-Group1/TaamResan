package storage

import (
	"TaamResan/internal/cart"
	"TaamResan/pkg/adapters/storage/entities"
	"TaamResan/pkg/adapters/storage/mappers"
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

func (r *cartRepo) GetItemsFeeByID(ctx context.Context, id uint) (float64, error) {
	sql := `
		SELECT SUM(f.price * ct.amount) AS "total_amount"
		FROM carts c
				 INNER JOIN cart_items ct ON c.id = ct.cart_id
				 INNER JOIN foods f ON ct.food_id = f.id AND f.deleted_at IS NULL
		WHERE c.id = ?
		`
	var amount float64
	err := r.db.WithContext(ctx).Raw(sql, id).Scan(&amount).Error

	if err != nil {
		return 0, err
	}

	return amount, nil
}
