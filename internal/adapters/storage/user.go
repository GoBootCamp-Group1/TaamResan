package storage

import (
	"TaamResan/internal/adapters/storage/entities"
	"TaamResan/internal/adapters/storage/mappers"
	"TaamResan/internal/user"
	"context"
	"errors"
	"gorm.io/gorm"
)

type userRepo struct {
	db *gorm.DB
}

func NewUserRepo(db *gorm.DB) user.Repo {
	return &userRepo{
		db: db,
	}
}

func (r *userRepo) Create(ctx context.Context, user *user.User) error {
	entity := mappers.DomainToUserEntity(user)
	return r.db.WithContext(ctx).Create(&entity).Error
}

func (r *userRepo) Update(ctx context.Context, user *user.User) error {
	entity := mappers.DomainToUserEntity(user)
	return r.db.WithContext(ctx).Save(&entity).Error
}

func (r *userRepo) GetByID(ctx context.Context, id uint) (*user.User, error) {
	panic("not implemented")
}

func (r *userRepo) GetByMobile(ctx context.Context, mobile string) (*user.User, error) {
	var user entities.User
	err := r.db.WithContext(ctx).Model(&entities.User{}).Where("mobile = ?", mobile).First(&user).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, err
		}
		return nil, err
	}
	return mappers.UserEntityToDomain(&user), nil
}

func (r *userRepo) GetByEmail(ctx context.Context, email string) (*user.User, error) {
	var user entities.User
	err := r.db.WithContext(ctx).Model(&entities.User{}).Where("email = ?", email).First(&user).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, err
		}
		return nil, err
	}
	return mappers.UserEntityToDomain(&user), nil
}
