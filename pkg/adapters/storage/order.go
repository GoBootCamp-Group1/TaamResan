package storage

import (
	"TaamResan/internal/cart"
	"TaamResan/internal/cart_item"
	"TaamResan/internal/order"
	"TaamResan/pkg/adapters/storage/entities"
	"TaamResan/pkg/adapters/storage/mappers"
	"context"
	"errors"
	"gorm.io/gorm"
)

type orderRepo struct {
	db *gorm.DB
}

var (
	ErrOrderNotFound = errors.New("error order not found")
)

func NewOrderRepo(db *gorm.DB) order.Repo { return &orderRepo{db: db} }

func (o *orderRepo) Create(ctx context.Context, data *order.InputData, cart *cart.Cart) (*order.Order, error) {

	orderEntity := &entities.Order{
		RestaurantID: *cart.RestaurantId,
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

func (o *orderRepo) Update(ctx context.Context, order *order.Order) (*order.Order, error) {

	orderEntity := mappers.DomainToOrderEntity(order)

	tErr := o.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		err := tx.WithContext(ctx).Model(&entities.Order{}).
			Where("id = ?", orderEntity.ID).Save(&orderEntity).Error
		if err != nil {
			return ErrUpdatingFood
		}

		return nil
	})

	if tErr != nil {
		return nil, tErr
	}

	return mappers.OrderEntityToDomain(orderEntity), nil
}

func (o *orderRepo) GetItemsCancellationFee(ctx context.Context, order *order.Order) (float64, error) {

	sql := `
		SELECT SUM((f.cancel_rate::numeric / 100) * f.price * ot.amount) AS "total_cancellation_amount"
		FROM orders o
				 INNER JOIN order_items ot ON o.id = ot.order_id
				 INNER JOIN foods f ON ot.food_id = f.id AND f.deleted_at IS NULL
		WHERE o.id = ?
		`
	var amount float64
	err := o.db.WithContext(ctx).Raw(sql, order.ID).Scan(&amount).Error

	if err != nil {
		return 0, err
	}

	return amount, nil
}

func (o *orderRepo) GetItemsFee(ctx context.Context, order *order.Order) (float64, error) {

	sql := `
		SELECT SUM(f.price * ot.amount) AS "total_amount"
		FROM orders o
				 INNER JOIN order_items ot ON o.id = ot.order_id
				 INNER JOIN foods f ON ot.food_id = f.id AND f.deleted_at IS NULL
		WHERE o.id = ?
		`
	var amount float64
	err := o.db.WithContext(ctx).Raw(sql, order.ID).Scan(&amount).Error

	if err != nil {
		return 0, err
	}

	return amount, nil
}

func (o *orderRepo) GetOrderByID(ctx context.Context, id uint) (*order.Order, error) {

	var orderEntity *entities.Order

	err := o.db.WithContext(ctx).Model(&entities.Order{}).
		Preload("Restaurant").
		Preload("Address").
		Where("id = ?", id).Find(&orderEntity).Error

	if err != nil {
		return nil, err
	}

	return mappers.OrderEntityToDomain(orderEntity), nil
}

func (o *orderRepo) AddCartItemToOrder(ctx context.Context, order *order.Order, item *cart_item.CartItem) error {

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

func (o *orderRepo) ChangeStatusByRestaurant(ctx context.Context, order *order.Order) error {
	// check if order exists
	var existingOrder *entities.Order
	if err := o.db.WithContext(ctx).Model(&entities.Order{}).Where("id = ?", order.ID).Find(&existingOrder).Error; err != nil {
		return err
	}

	if existingOrder.ID == 0 {
		return ErrOrderNotFound
	}

	// save
	existingOrder.Status = order.Status
	if err := o.db.WithContext(ctx).Model(&entities.Order{}).Where("id = ?", order.ID).Save(&existingOrder).Error; err != nil {
		return err
	}

	return nil
}
