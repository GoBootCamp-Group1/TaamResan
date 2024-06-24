package storage

import (
	"TaamResan/internal/adapters/storage/entities"
	"TaamResan/internal/adapters/storage/mappers"
	"TaamResan/internal/category_food"
	"context"
	"errors"
	"gorm.io/gorm"
)

type categoryFoodRepo struct {
	db *gorm.DB
}

var (
	ErrCategoryFoodExists   = errors.New("error category food exists")
	ErrCreatingCategoryFood = errors.New("error creating category food")
	ErrCategoryFoodNotFound = errors.New("error category food not found")
	ErrDeletingCategoryFood = errors.New("error deleting category food")
)

func NewCategoryFoodRepo(db *gorm.DB) category_food.Repo { return &categoryFoodRepo{db: db} }

func (r *categoryFoodRepo) Create(ctx context.Context, categoryFood *category_food.CategoryFood) (id uint, err error) {
	var existingCategoryFood *entities.CategoryFood
	err = r.db.WithContext(ctx).Model(&entities.CategoryFood{}).
		Where("category_id = ? and food_id = ?", categoryFood.CategoryId, categoryFood.FoodId).
		First(&existingCategoryFood).Error

	if err == nil {
		return 0, ErrCategoryFoodExists
	}

	if !errors.Is(err, gorm.ErrRecordNotFound) {
		return 0, ErrCreatingCategoryFood
	}

	err = r.db.Transaction(func(tx *gorm.DB) error {
		entity := mappers.DomainToCategoryFoodEntity(categoryFood)
		err = tx.WithContext(ctx).Model(&entities.CategoryFood{}).Create(&entity).Error
		if err != nil {
			return ErrCreatingCategoryFood
		}
		id = entity.ID

		return nil
	})

	return id, err
}

func (r *categoryFoodRepo) Delete(ctx context.Context, id uint) error {
	var existingCategoryFood *entities.CategoryFood
	err := r.db.WithContext(ctx).Model(&entities.CategoryFood{}).Where("id = ?", id).First(&existingCategoryFood).Error
	if err != nil {
		return ErrCategoryFoodNotFound
	}
	return r.db.Transaction(func(tx *gorm.DB) error {
		err = tx.WithContext(ctx).Model(&entities.CategoryFood{}).Where("id = ?", id).Delete(&existingCategoryFood).Error
		if err != nil {
			return ErrDeletingCategoryFood
		}

		return nil
	})
}

func (r *categoryFoodRepo) GetById(ctx context.Context, id uint) (*category_food.CategoryFood, error) {
	var existingCategoryFood *entities.CategoryFood
	err := r.db.WithContext(ctx).Model(&entities.CategoryFood{}).Where("id = ?", id).
		First(&existingCategoryFood).Error
	if err != nil {
		return nil, ErrCategoryFoodNotFound
	}
	return mappers.CategoryFoodEntityToDomain(existingCategoryFood), nil
}

func (r *categoryFoodRepo) GetAllByFoodId(ctx context.Context, foodId uint) ([]*category_food.CategoryFood, error) {
	var categoryFoods []*entities.CategoryFood
	err := r.db.WithContext(ctx).Model(&entities.CategoryFood{}).Where("food_id = ?", foodId).Find(&categoryFoods).Error
	if err != nil {
		return nil, ErrCategoryFoodNotFound
	}
	var models []*category_food.CategoryFood
	for _, e := range categoryFoods {
		model := mappers.CategoryFoodEntityToDomain(e)
		models = append(models, model)
	}
	return models, nil
}

func (r *categoryFoodRepo) GetAllByCategoryId(ctx context.Context, categoryId uint) ([]*category_food.CategoryFood, error) {
	var categoryFoods []*entities.CategoryFood
	err := r.db.WithContext(ctx).Model(&entities.CategoryFood{}).Where("category_id = ?", categoryId).Find(&categoryFoods).Error
	if err != nil {
		return nil, ErrCategoryNotFound
	}
	var models []*category_food.CategoryFood
	for _, e := range categoryFoods {
		model := mappers.CategoryFoodEntityToDomain(e)
		models = append(models, model)
	}
	return models, nil
}
