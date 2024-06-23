package storage

import (
	"TaamResan/internal/adapters/storage/entities"
	"TaamResan/internal/adapters/storage/mappers"
	"TaamResan/internal/category"
	"context"
	"errors"
	"gorm.io/gorm"
)

type categoryRepo struct {
	db *gorm.DB
}

func NewCategoryRepo(db *gorm.DB) category.Repo {
	return &categoryRepo{db: db}
}

var (
	ErrCreatingCategory = errors.New("error creating category")
	ErrCategoryExists   = errors.New("error category already exists")
	ErrUpdatingCategory = errors.New("error updating category")
	ErrDeletingCategory = errors.New("error deleting category")
	ErrCategoryNotFound = errors.New("error category not found")
)

func (r *categoryRepo) Create(ctx context.Context, cat *category.Category) (id uint, err error) {
	var existingCategory *category.Category
	err = r.db.WithContext(ctx).Model(&category.Category{}).
		Where("name = ? and restaurant_id = ?", cat.Name, cat.RestaurantId).First(&existingCategory).Error

	if err == nil {
		return 0, ErrCategoryExists
	}

	if err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			return 0, ErrCreatingCategory
		}
	}

	err = r.db.Transaction(func(tx *gorm.DB) error {
		entity := mappers.DomainToCategoryEntity(cat)
		err = tx.WithContext(ctx).Model(&entities.Category{}).Create(&entity).Error
		if err != nil {
			return ErrCreatingCategory
		}
		id = entity.ID

		return nil
	})

	return id, err
}

func (r *categoryRepo) Update(ctx context.Context, cat *category.Category) error {
	var existingCategory *entities.Category
	err := r.db.WithContext(ctx).Model(&entities.Category{}).Where("id = ?", cat.ID).First(&existingCategory).Error
	if err != nil {
		return ErrCategoryNotFound
	}

	return r.db.Transaction(func(tx *gorm.DB) error {
		if existingCategory.Name != cat.Name {
			if err = tx.WithContext(ctx).Model(&category.Category{}).
				Where("name = ? and restaurant_id = ?", cat.Name, cat.RestaurantId).First(&existingCategory).Error; err == nil {
				return ErrCategoryExists
			}
			existingCategory.Name = cat.Name
		}

		// check if parentId changed and is new parent id not equal to itself
		if existingCategory.ParentId != cat.ParentId && existingCategory.ID != *cat.ParentId {
			existingCategory.ParentId = cat.ParentId
		}

		err = tx.WithContext(ctx).Model(&entities.Category{}).Where("id = ?", cat.ID).Save(&existingCategory).Error
		if err != nil {
			return ErrUpdatingCategory
		}

		return nil
	})
}

func (r *categoryRepo) Delete(ctx context.Context, id uint) error {
	var existingCategory *category.Category
	err := r.db.WithContext(ctx).Model(&category.Category{}).Where("id = ?", id).First(&existingCategory).Error
	if err != nil {
		return ErrCategoryNotFound
	}
	return r.db.Transaction(func(tx *gorm.DB) error {
		err = tx.WithContext(ctx).Model(&category.Category{}).Where("id = ?", id).Delete(&existingCategory).Error
		if err != nil {
			return ErrDeletingCategory
		}

		return nil
	})
}

func (r *categoryRepo) GetById(ctx context.Context, id uint) (*category.Category, error) {
	var existingCategory *category.Category
	err := r.db.WithContext(ctx).Model(&category.Category{}).Where("id = ?", id).First(&existingCategory).Error
	if err != nil {
		return nil, ErrCategoryNotFound
	}
	return existingCategory, nil
}

func (r *categoryRepo) GetAll(ctx context.Context, restaurantId uint) ([]*category.Category, error) {
	var categories []*category.Category
	err := r.db.WithContext(ctx).Model(&category.Category{}).Where("restaurant_id = ?", restaurantId).Find(&categories).Error
	if err != nil {
		return nil, ErrCategoryNotFound
	}
	return categories, nil
}
