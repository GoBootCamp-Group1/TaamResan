package storage

import (
	"TaamResan/internal/adapters/storage/entities"
	"TaamResan/internal/adapters/storage/mappers"
	"TaamResan/internal/block_restaurant"
	"context"
	"errors"
	"gorm.io/gorm"
)

type blockRestaurantRepo struct {
	db *gorm.DB
}

func NewBlockRestaurantRepo(db *gorm.DB) block_restaurant.Repo {
	return &blockRestaurantRepo{db: db}
}

var (
	ErrBlockRestaurantExists   = errors.New("error restaurant is already blocked")
	ErrCreatingBlockRestaurant = errors.New("error blocking restaurant")
	ErrBlockRestaurantNotFound = errors.New("blocked restaurant not found")
	ErrDeletingBlockRestaurant = errors.New("error un-blocking restaurant")
	ErrFetchingBlockRestaurant = errors.New("error fetching blocked restaurant")
	ErrNoBlockRestaurantFound  = errors.New("error no blocked restaurant found")
)

func (r *blockRestaurantRepo) Create(ctx context.Context, br *block_restaurant.BlockRestaurant) (id uint, err error) {
	err = r.db.Transaction(func(tx *gorm.DB) error {
		// check existence
		var existingBlockRestaurant *entities.BlockRestaurant
		if err = tx.WithContext(ctx).Model(&entities.BlockRestaurant{}).
			Where("user_id = ? and restaurant_id = ?", br.UserId, br.RestaurantId).
			First(&existingBlockRestaurant).Error; err == nil {
			return ErrBlockRestaurantExists
		}

		if !errors.Is(err, gorm.ErrRecordNotFound) {
			return ErrCreatingBlockRestaurant
		}

		// create existence
		entity := mappers.DomainToBlockRestaurantEntity(br)
		if err = tx.WithContext(ctx).Model(&entities.BlockRestaurant{}).Create(&entity).Error; err != nil {
			return ErrCreatingBlockRestaurant
		}
		id = entity.ID

		return nil
	})

	return id, err
}

func (r *blockRestaurantRepo) Delete(ctx context.Context, id uint) error {
	return r.db.Transaction(func(tx *gorm.DB) error {
		// check existence
		var entity *entities.BlockRestaurant
		err := tx.WithContext(ctx).Model(&entities.BlockRestaurant{}).Where("id = ?", id).First(&entity).Error
		if err != nil {
			return ErrBlockRestaurantNotFound
		}

		// create existence
		if err = tx.WithContext(ctx).Model(&entities.BlockRestaurant{}).Where("id = ?", id).Delete(&entity).Error; err != nil {
			return ErrDeletingBlockRestaurant
		}

		return nil
	})
}

func (r *blockRestaurantRepo) GetAllByUserId(ctx context.Context, userId uint) ([]*block_restaurant.BlockRestaurant, error) {
	var restaurants []*entities.BlockRestaurant
	if err := r.db.WithContext(ctx).Model(&entities.BlockRestaurant{}).Where("user_id = ?", userId).Find(&restaurants).Error; err != nil {
		return nil, ErrFetchingBlockRestaurant
	}

	if len(restaurants) == 0 {
		return nil, ErrNoBlockRestaurantFound
	}

	var models []*block_restaurant.BlockRestaurant
	for _, restaurant := range restaurants {
		models = append(models, mappers.BlockRestaurantEntityToDomain(restaurant))
	}

	return models, nil
}
