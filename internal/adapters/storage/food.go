package storage

import (
	"TaamResan/internal/adapters/storage/entities"
	"TaamResan/internal/adapters/storage/mappers"
	"TaamResan/internal/category_food"
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
	// check food existence
	var existingFood *entities.Food
	err = r.db.WithContext(ctx).Model(&entities.Food{}).
		Where("name = ? and restaurant_id = ?", f.Name, f.RestaurantId).First(&existingFood).Error

	if err == nil {
		return 0, ErrFoodExists
	}

	if !errors.Is(err, gorm.ErrRecordNotFound) {
		return 0, ErrCreatingFood
	}

	err = r.db.Transaction(func(tx *gorm.DB) error {
		// create food
		entity := mappers.DomainToFoodEntity(f)
		err = tx.WithContext(ctx).Model(&entities.Food{}).Create(&entity).Error
		if err != nil {
			return ErrCreatingFood
		}
		id = entity.ID

		// create category food if it has category
		if len(f.Categories) > 0 {
			for _, c := range f.Categories {
				cf := category_food.CategoryFood{
					CategoryId: c.ID,
					FoodId:     id,
				}
				cfEntity := mappers.DomainToCategoryFoodEntity(&cf)
				err = tx.WithContext(ctx).Model(&entities.CategoryFood{}).Create(&cfEntity).Error
				if err != nil {
					return ErrCreatingCategoryFood
				}
			}
		}

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
	// check food existence
	var existingFood *entities.Food
	err := r.db.WithContext(ctx).Model(&entities.Food{}).Where("id = ?", id).First(&existingFood).Error
	if err != nil {
		return ErrFoodNotFound
	}
	return r.db.Transaction(func(tx *gorm.DB) error {
		// check category-food existence
		var cfEntity *entities.CategoryFood
		if err = tx.WithContext(ctx).Model(&entities.CategoryFood{}).Where("food_id = ?", id).First(&cfEntity).Error; err != nil {
			return ErrCategoryFoodNotFound
		}

		// delete category-food
		if err = tx.WithContext(ctx).Model(&entities.CategoryFood{}).Where("id = ?", cfEntity.ID).Delete(&cfEntity).Error; err != nil {
			return ErrDeletingCategoryFood
		}

		// delete food
		err = tx.WithContext(ctx).Model(&entities.Food{}).Where("id = ?", id).Delete(&existingFood).Error
		if err != nil {
			return ErrDeletingFood
		}

		return nil
	})
}

func (r *foodRepo) GetById(ctx context.Context, id uint) (*food.Food, error) {
	var existingFood *entities.Food
	err := r.db.WithContext(ctx).Model(&entities.Food{}).Where("id = ?", id).First(&existingFood).Error
	if err != nil {
		return nil, ErrFoodNotFound
	}
	return mappers.FoodEntityToDomain(existingFood), nil
}

func (r *foodRepo) GetAll(ctx context.Context, restaurantId uint) ([]*food.Food, error) {
	var foods []*entities.Food
	err := r.db.WithContext(ctx).Model(&entities.Food{}).Where("restaurant_id = ?", restaurantId).Find(&foods).Error
	if err != nil {
		return nil, ErrFoodNotFound
	}
	var models []*food.Food
	for _, f := range foods {
		model := mappers.FoodEntityToDomain(f)
		models = append(models, model)
	}
	return models, nil
}

func (r *foodRepo) SearchFoods(ctx context.Context, searchData *food.FoodSearch) ([]*food.Food, error) {

	var foods []*entities.Food
	query := r.db.Debug().WithContext(ctx).Model(&entities.Food{}).
		Preload("Categories").
		Preload("Restaurant").
		Joins("JOIN category_foods ON category_foods.food_id = foods.id").
		Joins("JOIN categories ON categories.id = category_foods.category_id")

	if searchData.Name != "" {
		query = query.Where("foods.name LIKE ?", "%"+searchData.Name+"%")
	}

	if searchData.CategoryID != nil {
		query = query.Where("categories.id = ?", *searchData.CategoryID)
	}

	if searchData.Lat != nil && searchData.Lng != nil {
		//TODO: using database related method, for range of 5000 meters in radius
		query = query.Joins("JOIN restaurants ON restaurants.id = foods.restaurant_id").
			Where("ST_DistanceSphere(ST_MakePoint(restaurants.lng, restaurants.lat), ST_MakePoint(?, ?)) < ?", searchData.Lng, searchData.Lat, 5000)
	}

	err := query.Find(&foods).Error

	if err != nil {
		return nil, ErrFoodNotFound
	}

	var models []*food.Food
	for _, f := range foods {
		model := mappers.FoodEntityToDomain(f)
		models = append(models, model)
	}
	return models, nil
}
