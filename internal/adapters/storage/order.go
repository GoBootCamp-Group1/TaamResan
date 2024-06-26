package storage

import (
	"TaamResan/internal/adapters/storage/entities"
	"TaamResan/internal/adapters/storage/mappers"
	"TaamResan/internal/cart"
	"TaamResan/internal/cart_item"
	"TaamResan/internal/order"
	"context"
	"fmt"
	"gorm.io/gorm"
)

type orderRepo struct {
	db *gorm.DB
}

func NewOrderRepo(db *gorm.DB) order.Repo { return &orderRepo{db: db} }

func (o *orderRepo) Create(ctx context.Context, data *order.InputData, cart *cart.Cart) (*order.Order, error) {

	orderEntity := &entities.Order{
		RestaurantID: cart.RestaurantId,
		UserID:       cart.UserId,
		AddressID:    data.AddressID,
		Status:       order.STATUS_PAID,
	}

	tErr := o.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {

		err := tx.Model(&entities.Order{}).Create(orderEntity).Error

		if err != nil {
			return err
		}

		// Preload the relationships
		if err := tx.Model(&entities.Order{}).
			Preload("Restaurant").
			Preload("Address").
			Find(orderEntity).Error; err != nil {
			return err
		}

		for _, item := range cart.Items {

			orderItemEntity := &entities.OrderItem{
				OrderId: orderEntity.ID,
				FoodId:  item.FoodId,
				Amount:  item.Amount,
				Note:    "",
			}

			err := tx.Model(&entities.OrderItem{}).Create(&orderItemEntity).Error
			if err != nil {
				return err
			}

			err = tx.Delete(&entities.CartItem{}, item.ID).Error
			if err != nil {
				return err
			}

		}
		return nil
	})

	if tErr != nil {
		return nil, tErr
	}

	return mappers.OrderEntityToDomain(orderEntity), nil
}

func (o *orderRepo) AddCartItemToOrder(ctx context.Context, order *order.Order, item *cart_item.CartItem) error {

	fmt.Println("ORDER: ", order)
	fmt.Println("ORDER ID: ", order.ID)

	orderItemEntity := &entities.OrderItem{
		OrderId: order.ID,
		FoodId:  item.FoodId,
		Amount:  item.Amount,
		Note:    "",
	}

	return o.db.Transaction(func(tx *gorm.DB) error {
		err := o.db.WithContext(ctx).Model(&entities.OrderItem{}).Create(&orderItemEntity).Error
		if err != nil {
			return err
		}

		err = o.db.WithContext(ctx).Model(&entities.CartItem{}).Delete(item).Error
		if err != nil {
			return err
		}

		return nil
	})
}
