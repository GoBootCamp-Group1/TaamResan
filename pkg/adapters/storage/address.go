package storage

import (
	"TaamResan/internal/address"
	"TaamResan/pkg/adapters/storage/entities"
	"TaamResan/pkg/adapters/storage/mappers"
	"TaamResan/pkg/jwt"
	"context"
	"errors"
	"gorm.io/gorm"
	"time"
)

type addressRepo struct {
	db *gorm.DB
}

func NewAddressRepo(db *gorm.DB) address.Repo {
	return &addressRepo{
		db: db,
	}
}

var (
	ErrNotAllowed = errors.New("error not allowed")
)

func (r addressRepo) Create(ctx context.Context, address *address.Address) error {
	return r.db.Transaction(func(tx *gorm.DB) error {

		addressEntity := mappers.DomainToAddressEntity(address)

		// Assume we have the user ID of the logged-in user
		userID := ctx.Value(jwt.UserClaimKey).(*jwt.UserClaims).UserID

		// Fetch the existing user by ID
		var existingUser entities.User
		if err := tx.First(&existingUser, userID).Error; err != nil {
			return err
		}

		// Save the new address
		if err := tx.Create(&addressEntity).Error; err != nil {
			return err
		}

		//TODO: not working
		//// Associate the new address with the existing user
		//if err := tx.Debug().Model(&existingUser).Association("Addresses").Append(&addressEntity); err != nil {
		//	fmt.Println(err)
		//	return err
		//}

		//TODO: not filling created_at and updated_at
		//// Manually append address to user's address list
		//existingUser.Addresses = append(existingUser.Addresses, addressEntity)
		//
		//// Save user to update the join table
		//if err := tx.Save(&existingUser).Error; err != nil {
		//	return err
		//}

		// Create an entry in the custom join table
		userAddress := entities.UserAddress{
			UserID:    existingUser.ID,
			AddressID: addressEntity.ID,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		}
		if err := tx.Create(&userAddress).Error; err != nil {
			return err
		}

		return nil
	})
}

func (r addressRepo) Update(ctx context.Context, address *address.Address) error {
	addressEntity := mappers.DomainToAddressEntity(address)
	//update address entity
	return r.db.WithContext(ctx).Save(&addressEntity).Error
}

func (r addressRepo) Delete(ctx context.Context, address *address.Address) error {
	return r.db.Transaction(func(tx *gorm.DB) error {
		if err := r.checkUserAccess(ctx, address.ID); err != nil {
			return err
		}

		addressEntity := mappers.DomainToAddressEntity(address)

		//get user id from context
		userID := ctx.Value(jwt.UserClaimKey).(*jwt.UserClaims).UserID

		//delete user address
		err := tx.Delete(&entities.UserAddress{
			UserID:    userID,
			AddressID: address.ID,
		}).Error
		if err != nil {
			return err
		}

		//delete address entity
		err = tx.Delete(&addressEntity).Error
		if err != nil {
			return err
		}

		return nil
	})
}

func (r addressRepo) GetByID(ctx context.Context, id uint) (*address.Address, error) {
	var addressEntity entities.Address
	err := r.db.WithContext(ctx).Model(&entities.Address{}).Where("id = ?", id).First(&addressEntity).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, err
		}
		return nil, err
	}
	return mappers.AddressEntityToDomain(&addressEntity), nil
}

func (r addressRepo) GetAll(ctx context.Context) ([]*address.Address, error) {
	userID := ctx.Value(jwt.UserClaimKey).(*jwt.UserClaims).UserID
	var addressEntities []*address.Address
	err := r.db.WithContext(ctx).Model(&address.Address{}).
		Joins("join user_addresses on user_addresses.address_id = addresses.id and user_addresses.user_id = ?", userID).
		Find(&addressEntities).Error
	if err != nil {
		return nil, err
	}
	return addressEntities, nil
}

func (r addressRepo) checkUserAccess(ctx context.Context, addressId uint) error {
	userID := ctx.Value(jwt.UserClaimKey).(*jwt.UserClaims).UserID
	var userAddress *entities.UserAddress
	if err := r.db.WithContext(ctx).Model(&entities.UserAddress{}).
		Where("address_id = ? and user_id = ?", addressId, userID).First(&userAddress).Error; err != nil {
		return ErrNotAllowed
	}

	return nil
}
