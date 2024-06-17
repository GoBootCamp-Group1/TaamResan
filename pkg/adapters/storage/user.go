package storage

import (
	"TaamResan/internal/user"
	"TaamResan/pkg/adapters/storage/entities"
	"TaamResan/pkg/adapters/storage/mappers"
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
	panic("not implemented")
}

func (r *userRepo) GetByID(ctx context.Context, id uint) (*user.User, error) {
	panic("not implemented")
}

func (r *userRepo) GetByEmail(ctx context.Context, email string) (*user.User, error) {
	var user entities.User
	err := r.db.WithContext(ctx).Model(&entities.User{}).Where("email = ?", email).First(&user).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return mappers.UserEntityToDomain(&user), nil
}
