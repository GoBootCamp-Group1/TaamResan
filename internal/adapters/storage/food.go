package storage

import (
	"TaamResan/internal/adapters/storage/entities"
	"TaamResan/internal/adapters/storage/mappers"
	"TaamResan/internal/food"
	"context"
	"errors"
	"gorm.io/gorm"
)

type foodRepo struct {
	db *gorm.DB
}

func NewFoodRepo(db *gorm.DB) food.Repo {
	return &foodRepo{db: db}
}

var (
	ErrCreatingFood = errors.New("error creating food")
	ErrFoodExists   = errors.New("error food already exists")
	ErrUpdatingFood = errors.New("error updating food")
	ErrDeletingFood = errors.New("error deleting food")
	ErrFoodNotFound = errors.New("error food not found")
)

func (r *foodRepo) Create(ctx context.Context, f *food.Food) (id uint, err error) {
	var existingFood *food.Food
	err = r.db.WithContext(ctx).Model(&food.Food{}).
		Where("name = ? and restaurant_id = ?", f.Name, f.RestaurantId).First(&existingFood).Error

	if err == nil {
		return 0, ErrFoodExists
	}

	if err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			return 0, ErrCreatingFood
		}
	}

	err = r.db.Transaction(func(tx *gorm.DB) error {
		entity := mappers.DomainToFoodEntity(f)
		err = tx.WithContext(ctx).Model(&entities.Food{}).Create(&entity).Error
		if err != nil {
			return ErrCreatingFood
		}
		id = entity.ID

		return nil
	})

	return id, err
}

func (r *foodRepo) Update(ctx context.Context, f *food.Food) error {
	var existingFood *entities.Food
	err := r.db.WithContext(ctx).Model(&entities.Food{}).Where("id = ?", f.ID).First(&existingFood).Error
	if err != nil {
		return ErrFoodNotFound
	}

	return r.db.Transaction(func(tx *gorm.DB) error {
		if existingFood.Name != f.Name {
			if err = tx.WithContext(ctx).Model(&food.Food{}).
				Where("name = ? and restaurant_id = ?", f.Name, f.RestaurantId).First(&existingFood).Error; err == nil {
				return ErrFoodExists
			}
			existingFood.Name = f.Name
		}

		existingFood.Price = f.Price
		existingFood.CancelRate = f.CancelRate
		existingFood.PreparationMinutes = f.PreparationMinutes

		err = tx.WithContext(ctx).Model(&entities.Food{}).Where("id = ?", f.ID).Save(&existingFood).Error
		if err != nil {
			return ErrUpdatingFood
		}

		return nil
	})
}

func (r *foodRepo) Delete(ctx context.Context, id uint) error {
	var existingFood *entities.Food
	err := r.db.WithContext(ctx).Model(&entities.Food{}).Where("id = ?", id).First(&existingFood).Error
	if err != nil {
		return ErrFoodNotFound
	}
	return r.db.Transaction(func(tx *gorm.DB) error {
		err = tx.WithContext(ctx).Model(&entities.Food{}).Where("id = ?", id).Delete(&existingFood).Error
		if err != nil {
			return ErrDeletingFood
		}

		return nil
	})
}

func (r *foodRepo) GetById(ctx context.Context, id uint) (*food.Food, error) {
	var existingFood *food.Food
	err := r.db.WithContext(ctx).Model(&food.Food{}).Where("id = ?", id).First(&existingFood).Error
	if err != nil {
		return nil, ErrFoodNotFound
	}
	return existingFood, nil
}

func (r *foodRepo) GetAll(ctx context.Context, restaurantId uint) ([]*food.Food, error) {
	var foods []*food.Food
	err := r.db.WithContext(ctx).Model(&food.Food{}).Where("restaurant_id = ?", restaurantId).Find(&foods).Error
	if err != nil {
		return nil, ErrFoodNotFound
	}
	return foods, nil
}
