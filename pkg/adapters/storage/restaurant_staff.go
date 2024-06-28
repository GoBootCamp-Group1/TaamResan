package storage

import (
	"TaamResan/internal/restaurant_staff"
	"TaamResan/pkg/adapters/storage/entities"
	"TaamResan/pkg/adapters/storage/mappers"
	"context"
	"errors"
	"gorm.io/gorm"
)

type restaurantStaffRepo struct {
	db *gorm.DB
}

func NewRestaurantStaffRepo(db *gorm.DB) restaurant_staff.Repo {
	return &restaurantStaffRepo{db: db}
}

var (
	ErrCreatingRestaurantStaff = errors.New("error creating restaurant staff")
	ErrRestaurantStaffExists   = errors.New("error restaurant staff already exists")
	ErrRestaurantStaffNotFound = errors.New("error restaurant staff not found")
	ErrDeletingRestaurantStaff = errors.New("error deleting restaurant staff")
)

func (r *restaurantStaffRepo) Create(ctx context.Context, rStaff *restaurant_staff.RestaurantStaff) (id uint, err error) {
	// check existence
	var existingRestaurantStaff *entities.RestaurantStaff
	err = r.db.WithContext(ctx).Model(&entities.RestaurantStaff{}).
		Where("restaurant_id = ? and user_id = ?", rStaff.RestaurantId, rStaff.UserId).First(&existingRestaurantStaff).Error
	if err == nil {
		return 0, ErrRestaurantStaffExists
	}

	if !errors.Is(err, gorm.ErrRecordNotFound) {
		return 0, ErrCreatingRestaurantStaff
	}

	err = r.db.Transaction(func(tx *gorm.DB) error {
		entity := mappers.DomainToRestaurantStaffEntity(rStaff)
		err = tx.WithContext(ctx).Model(&entities.RestaurantStaff{}).Create(&entity).Error
		if err != nil {
			return ErrCreatingRestaurantStaff
		}
		id = entity.ID

		return nil
	})
	return id, err
}

func (r *restaurantStaffRepo) Delete(ctx context.Context, id uint) error {
	return r.db.Transaction(func(tx *gorm.DB) error {
		var entity *entities.RestaurantStaff
		if err := tx.WithContext(ctx).Model(&entities.RestaurantStaff{}).Where("id = ?", id).First(&entity).Error; err != nil {
			return ErrRestaurantStaffNotFound
		}
		if err := tx.WithContext(ctx).Model(&entities.RestaurantStaff{}).Where("id = ?", id).Delete(&entity).Error; err != nil {
			return ErrDeletingRestaurantStaff
		}
		return nil
	})
}

func (r *restaurantStaffRepo) GetAllByRestaurantId(ctx context.Context, restaurantId uint) ([]*restaurant_staff.RestaurantStaff, error) {
	var staffs []*restaurant_staff.RestaurantStaff
	if err := r.db.WithContext(ctx).Model(&entities.RestaurantStaff{}).Where("restaurant_id = ?", restaurantId).Find(&staffs).Error; err != nil {
		return nil, ErrRestaurantStaffNotFound
	}
	return staffs, nil
}

func (r *restaurantStaffRepo) GetById(ctx context.Context, id uint) (*restaurant_staff.RestaurantStaff, error) {
	var entity *entities.RestaurantStaff
	if err := r.db.WithContext(ctx).Model(&entities.RestaurantStaff{}).Where("id = ?", id).First(&entity).Error; err != nil {
		return nil, ErrRestaurantStaffNotFound
	}

	return mappers.RestaurantStaffEntityToDomain(entity), nil
}

func (r *restaurantStaffRepo) GetOwnerByRestaurantId(ctx context.Context, restaurantId uint) (*restaurant_staff.RestaurantStaff, error) {
	var entity *entities.RestaurantStaff
	if err := r.db.WithContext(ctx).Model(&entities.RestaurantStaff{}).
		Where("restaurant_id = ?", restaurantId).
		Where("position = ?", restaurant_staff.Manager).
		First(&entity).Error; err != nil {
		return nil, ErrRestaurantStaffNotFound
	}

	return mappers.RestaurantStaffEntityToDomain(entity), nil
}
