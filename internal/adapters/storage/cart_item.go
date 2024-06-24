package storage

import (
	"TaamResan/internal/adapters/storage/entities"
	"TaamResan/internal/adapters/storage/mappers"
	"TaamResan/internal/cart"
	"TaamResan/internal/cart_item"
	"context"
	"errors"
	"gorm.io/gorm"
)

type cartItemRepo struct {
	db *gorm.DB
}

func NewCartItemRepo(db *gorm.DB) cart_item.Repo {
	return &cartItemRepo{db: db}
}

var (
	ErrCartItemExists   = errors.New("error cart item already exists")
	ErrCreatingCartItem = errors.New("error creating cartItem")
	ErrCartItemNotFound = errors.New("error cart item not found")
	ErrUpdatingCartItem = errors.New("error updating cart item")
	ErrDeletingCartItem = errors.New("error deleting cart item")
	ErrNoCartItemFound  = errors.New("error no cart item found")
)

func (r *cartItemRepo) Create(ctx context.Context, cartItem *cart_item.CartItem) (id uint, err error) {
	err = r.db.Transaction(func(tx *gorm.DB) error {
		var entity *entities.CartItem
		if err = tx.WithContext(ctx).Model(&entities.CartItem{}).
			Where("cart_id = ? and food_id = ?", cartItem.CartId, cartItem.FoodId).
			First(&entity).Error; err == nil {
			return ErrCartItemExists
		}

		if !errors.Is(err, gorm.ErrRecordNotFound) {
			return ErrCreatingCartItem
		}

		entity = mappers.DomainToCartItemEntity(cartItem)
		if err = tx.WithContext(ctx).Model(&entities.CartItem{}).Create(&entity).Error; err != nil {
			return ErrCreatingCartItem
		}
		id = entity.ID

		// update cart's restaurantId if is nil
		var cartEntity *cart.Cart
		if err = tx.WithContext(ctx).Model(&entities.Cart{}).Debug().Where("id = ?", cartItem.CartId).Find(&cartEntity).Error; err != nil {
			return err
		}
		if cartEntity.RestaurantId == nil {
			var foodEntity *entities.Food
			if err = tx.WithContext(ctx).Model(&entities.Food{}).Where("id = ?", cartItem.FoodId).Find(&foodEntity).Error; err != nil {
				return err
			}
			cartEntity.RestaurantId = &foodEntity.RestaurantId
			if err = tx.WithContext(ctx).Model(&entities.Cart{}).Debug().Where("id = ?", cartItem.CartId).Save(&cartEntity).Error; err != nil {
				return err
			}
		}

		return nil
	})

	return id, err
}

func (r *cartItemRepo) Update(ctx context.Context, cartItem *cart_item.CartItem) error {
	return r.db.Transaction(func(tx *gorm.DB) error {
		// check for existence
		var entity *entities.CartItem
		if err := tx.WithContext(ctx).Model(&entities.CartItem{}).Where("id = ?", cartItem.ID).Find(&entity).Error; err != nil {
			return ErrCartItemNotFound
		}

		// check for update
		if cartItem.Amount != entity.Amount {
			entity.Amount = cartItem.Amount
		}

		if cartItem.Note != entity.Note {
			entity.Note = cartItem.Note
		}

		if err := tx.WithContext(ctx).Model(&entities.CartItem{}).Where("id = ?", cartItem.ID).Save(&entity).Error; err != nil {
			return ErrUpdatingCartItem
		}

		return nil
	})
}

func (r *cartItemRepo) Delete(ctx context.Context, id uint) error {
	return r.db.Transaction(func(tx *gorm.DB) error {
		var entity *entities.CartItem
		if err := tx.WithContext(ctx).Model(&entities.CartItem{}).Where("id = ?", id).Find(&entity).Error; err != nil {
			return ErrCartItemNotFound
		}

		if err := tx.WithContext(ctx).Model(&entities.CartItem{}).Where("id = ?", id).Delete(&entity).Error; err != nil {
			return ErrDeletingCartItem
		}

		return nil
	})
}

func (r *cartItemRepo) GetAllByCartId(ctx context.Context, cartId uint) ([]*cart_item.CartItem, error) {
	// check cart exists
	var cartEntity *entities.Cart
	err := r.db.WithContext(ctx).Model(&entities.Cart{}).Where("id = ?", cartId).Find(&cartEntity).Error
	if err != nil || cartEntity.ID == 0 {
		return nil, ErrCartNotFound
	}

	// get all
	var items []*entities.CartItem
	if err := r.db.WithContext(ctx).Model(&entities.CartItem{}).Where("cart_id = ?", cartId).Find(&items).Error; err != nil {
		return nil, err
	}

	if len(items) == 0 {
		return nil, ErrNoCartItemFound
	}

	var models []*cart_item.CartItem
	if len(items) > 0 {
		for _, item := range items {
			models = append(models, mappers.CartItemEntityToDomain(item))
		}
	}
	return models, nil
}
