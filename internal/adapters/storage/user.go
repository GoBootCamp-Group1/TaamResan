package storage

import (
	"TaamResan/internal/adapters/storage/entities"
	"TaamResan/internal/adapters/storage/mappers"
	"TaamResan/internal/role"
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
	return r.db.Transaction(func(tx *gorm.DB) error {
		entity := mappers.DomainToUserEntity(user)
		err := tx.WithContext(ctx).Create(&entity).Error
		if err != nil {
			return err
		}

		var userRoles []entities.UserRoles
		if len(user.Roles) > 0 {
			for _, role := range user.Roles {
				ur := entities.UserRoles{
					UserId: user.ID,
					RoleId: role.ID,
				}
				userRoles = append(userRoles, ur)
			}
		} else {
			userRoles = append(userRoles, entities.UserRoles{
				UserId: entity.ID,
				RoleId: role.DefaultRole.ID,
			})
		}
		if err = tx.Create(&userRoles).Error; err != nil {
			return err
		}

		//Create Wallet
		walletEntity := entities.Wallet{
			UserID: entity.ID,
			Credit: 0.0,
		}
		if err = tx.Create(&walletEntity).Error; err != nil {
			return err
		}

		// create cart
		cartEntity := entities.Cart{
			UserId: entity.ID,
		}
		if err = tx.WithContext(ctx).Model(&entities.Cart{}).Create(&cartEntity).Error; err != nil {
			return err
		}

		return nil
	})
}

func (r *userRepo) Update(ctx context.Context, user *user.User) error {
	entity := mappers.DomainToUserEntity(user)
	return r.db.WithContext(ctx).Save(&entity).Error
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

//// TODO:REVIEW
//func (r *userRepo) GetUserActiveWallet(ctx context.Context, userId uint) (*wallet.Wallet, error) {
//	var w entities.Wallet
//	err := r.db.WithContext(ctx).Model(&entities.Wallet{}).Where("user_id = ?", userId).First(&w).Error
//	if err != nil {
//		if errors.Is(err, gorm.ErrRecordNotFound) {
//			return nil, err
//		}
//		return nil, err
//	}
//	return mappers.WalletEntityToDomain(&w), nil
//}

func (r *userRepo) CreateAdmin(ctx context.Context) error {
	// check that if have user with admin role
	var userRole *entities.UserRoles
	if err := r.db.WithContext(ctx).Model(&entities.UserRoles{}).
		Where("role_id = ?", role.Admin).Find(&userRole).Error; err != nil {
		return err
	}

	if userRole.ID != 0 {
		return nil // admin already created, do nothing
	}

	// if not, create one
	return r.db.Transaction(func(tx *gorm.DB) error {
		entity := mappers.DomainToUserEntity(&user.DefaultAdminUser)
		if err := tx.WithContext(ctx).Model(&entities.User{}).Create(entity).Error; err != nil {
			return err
		}

		ur := entities.UserRoles{UserId: entity.ID, RoleId: role.Admin}
		if err := tx.WithContext(ctx).Model(&entities.UserRoles{}).Create(&ur).Error; err != nil {
			return err
		}

		return nil
	})
}
