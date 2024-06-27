package storage

import (
	"TaamResan/internal/adapters/storage/entities"
	"TaamResan/internal/adapters/storage/mappers"
	"TaamResan/internal/address"
	"TaamResan/internal/restaurant"
	"TaamResan/internal/restaurant_staff"
	"TaamResan/internal/role"
	"TaamResan/pkg/jwt"
	"context"
	"errors"
	"gorm.io/gorm"
	"math"
)

type restaurantRepo struct {
	db        *gorm.DB
	addrRepo  address.Repo
	staffRepo restaurant_staff.Repo
}

func NewRestaurantRepo(db *gorm.DB) restaurant.Repo {
	return &restaurantRepo{
		db:        db,
		addrRepo:  NewAddressRepo(db),
		staffRepo: NewRestaurantStaffRepo(db),
	}
}

var (
	ErrCreatingAddr        = errors.New("error creating address")
	ErrCreatingRestaurant  = errors.New("error creating restaurant")
	ErrRestaurantExists    = errors.New("error restaurant already exists")
	ErrRestaurantNotFound  = errors.New("error restaurant not found")
	ErrAddrNotFound        = errors.New("error address not found")
	ErrUpdatingRestaurant  = errors.New("error updating restaurant")
	ErrUpdatingAddr        = errors.New("error updating address")
	ErrDeletingRestaurant  = errors.New("error deleting restaurant")
	ErrFetchingRestaurants = errors.New("error fetching restaurant")
)

func (r *restaurantRepo) Create(ctx context.Context, restaurant *restaurant.Restaurant) (id uint, err error) {
	// check addr existence
	var addrEntity *entities.Address
	err = r.db.WithContext(ctx).Model(&entities.Address{}).
		Where("title = ? and lat = ? and lng = ?", restaurant.Address.Title, restaurant.Address.Lat, restaurant.Address.Lng).
		Find(&addrEntity).Error
	if err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			return 0, err
		}
	}

	err = r.db.Transaction(func(tx *gorm.DB) error {
		if addrEntity.ID == 0 {
			// create addr
			addrEntity = mappers.DomainToAddressEntity(&restaurant.Address)
			if err1 := tx.WithContext(ctx).Model(&entities.Address{}).Create(&addrEntity).Error; err1 != nil {
				return ErrCreatingAddr
			}
		}

		// check restaurant existence
		var existingRestaurant *entities.Restaurant
		err = r.db.WithContext(ctx).Model(&entities.Restaurant{}).
			Where("name = ? and address_id = ? and owned_by = ?", restaurant.Name, addrEntity.ID, restaurant.OwnedBy).
			Find(&existingRestaurant).Error
		if existingRestaurant.ID != 0 {
			return ErrRestaurantExists
		}

		// create restaurant
		restaurant.Address.ID = addrEntity.ID
		restaurantEntity := mappers.DomainToRestaurantEntity(restaurant)
		if err1 := tx.WithContext(ctx).Model(&entities.Restaurant{}).Create(&restaurantEntity).Error; err1 != nil {
			return ErrCreatingRestaurant
		}
		id = restaurantEntity.ID

		// check restaurant staff existence
		var existingRestaurantStaff *entities.RestaurantStaff
		err = r.db.WithContext(ctx).Model(&entities.RestaurantStaff{}).
			Where("restaurant_id = ? and user_id = ?", restaurantEntity.ID, restaurant.OwnedBy).First(&existingRestaurantStaff).Error
		if err == nil {
			return ErrRestaurantStaffExists
		}

		if !errors.Is(err, gorm.ErrRecordNotFound) {
			return ErrCreatingRestaurantStaff
		}

		// create restaurant staff
		restaurantStaff := restaurant_staff.RestaurantStaff{
			UserId:       restaurant.OwnedBy,
			RestaurantId: restaurantEntity.ID,
			Position:     restaurant_staff.Manager,
		}
		entity := mappers.DomainToRestaurantStaffEntity(&restaurantStaff)
		err = tx.WithContext(ctx).Model(&entities.RestaurantStaff{}).Create(&entity).Error
		if err != nil {
			return ErrCreatingRestaurantStaff
		}

		// add to roles
		ownerRole := entities.UserRoles{
			UserId: restaurant.OwnedBy,
			RoleId: role.RestaurantOwner,
		}
		if err = tx.WithContext(ctx).Model(&entities.UserRoles{}).Save(&ownerRole).Error; err != nil {
			return ErrCreatingRestaurant
		}

		return nil
	})
	return id, err
}

func (r *restaurantRepo) Update(ctx context.Context, restaurant *restaurant.Restaurant) error {
	// check restaurant existence
	var existingRestaurant entities.Restaurant
	err := r.db.WithContext(ctx).Model(&entities.Restaurant{}).Where("id = ?", restaurant.ID).
		First(&existingRestaurant).Error
	if err != nil {
		return ErrRestaurantNotFound
	}
	return r.db.Transaction(func(tx *gorm.DB) error {
		// find addr
		addr, err1 := r.addrRepo.GetByID(ctx, existingRestaurant.AddressId)
		if err1 != nil {
			return ErrAddrNotFound
		}
		if isAddressChanged(restaurant.Address, *addr) {
			// update addr
			addr.Title = restaurant.Address.Title
			addr.Lat = restaurant.Address.Lat
			addr.Lng = restaurant.Address.Lng
			if err1 = r.addrRepo.Update(ctx, addr); err1 != nil {
				return ErrUpdatingAddr
			}
		}

		epsilon := 1e-9
		if math.Abs(restaurant.CourierSpeed-existingRestaurant.CourierSpeed) > epsilon {
			existingRestaurant.CourierSpeed = restaurant.CourierSpeed
		}

		if restaurant.Name != existingRestaurant.Name {
			existingRestaurant.Name = restaurant.Name
		}

		// update restaurant
		if err = tx.WithContext(ctx).Model(&entities.Restaurant{}).Where("id = ?", existingRestaurant.ID).Save(&existingRestaurant).Error; err != nil {
			return ErrUpdatingRestaurant
		}

		return nil
	})
}

func isAddressChanged(addr address.Address, loadedAddr address.Address) bool {
	return addr.Title != loadedAddr.Title || addr.Lat != loadedAddr.Lat || addr.Lng != loadedAddr.Lng
}

func (r *restaurantRepo) Delete(ctx context.Context, id uint) error {
	// check restaurant existence
	var existingRestaurant *entities.Restaurant
	err := r.db.WithContext(ctx).Model(&entities.Restaurant{}).Where("id = ?", id).First(&existingRestaurant).Error
	if err != nil {
		return ErrRestaurantNotFound
	}

	// find staffs
	staffs, err := r.staffRepo.GetAllByRestaurantId(ctx, id)
	if err != nil {
		return err
	}

	return r.db.Transaction(func(tx *gorm.DB) error {
		// delete staffs
		if len(staffs) > 0 {
			for _, staff := range staffs {
				if err1 := r.staffRepo.Delete(ctx, staff.ID); err1 != nil {
					return err1
				}
			}
		}

		// delete restaurant
		err = tx.WithContext(ctx).Model(&entities.Restaurant{}).Where("id = ?", id).Delete(&existingRestaurant).Error
		if err != nil {
			return ErrDeletingRestaurant
		}

		return nil
	})
}

func (r *restaurantRepo) GetById(ctx context.Context, id uint) (*restaurant.Restaurant, error) {
	// check restaurant existence
	var existingRestaurant entities.Restaurant
	err := r.db.WithContext(ctx).Model(&entities.Restaurant{}).Where("id = ?", id).First(&existingRestaurant).Error
	if err != nil {
		return nil, ErrRestaurantNotFound
	}
	restaurantModel := mappers.RestaurantEntityToDomain(&existingRestaurant)

	// get address
	addr, err := r.addrRepo.GetByID(ctx, existingRestaurant.AddressId)
	if err != nil {
		return nil, ErrAddrNotFound
	}
	restaurantModel.Address.Title = addr.Title
	restaurantModel.Address.Lat = addr.Lat
	restaurantModel.Address.Lng = addr.Lng

	return restaurantModel, nil
}

func (r *restaurantRepo) GetAll(ctx context.Context) ([]*restaurant.Restaurant, error) {
	userId := ctx.Value(jwt.UserClaimKey).(*jwt.UserClaims).UserID
	var restaurantsEntities []*entities.Restaurant
	err := r.db.WithContext(ctx).Model(&entities.Restaurant{}).Where("owned_by = ?", userId).Find(&restaurantsEntities).Error
	if err != nil {
		return nil, ErrFetchingRestaurants
	}
	var restaurants []*restaurant.Restaurant
	for _, res := range restaurantsEntities {
		model := mappers.RestaurantEntityToDomain(res)
		addr, err := r.addrRepo.GetByID(ctx, res.AddressId)
		if err != nil {
			return nil, ErrAddrNotFound
		}
		model.Address.Title = addr.Title
		model.Address.Lat = addr.Lat
		model.Address.Lng = addr.Lng
		restaurants = append(restaurants, model)
	}

	return restaurants, nil
}

func (r *restaurantRepo) Approve(ctx context.Context, id uint) error {
	// check restaurant existence
	var existingRestaurant entities.Restaurant
	err := r.db.WithContext(ctx).Model(&entities.Restaurant{}).Where("id = ?", id).First(&existingRestaurant).Error
	if err != nil {
		return ErrRestaurantNotFound
	}
	return r.db.Transaction(func(tx *gorm.DB) error {
		// approve restaurant and update
		existingRestaurant.ApprovalStatus = restaurant.Approved
		if err = tx.WithContext(ctx).Model(&entities.Restaurant{}).Where("id = ?", id).Save(&existingRestaurant).Error; err != nil {
			return ErrUpdatingRestaurant
		}

		return nil
	})
}

func (r *restaurantRepo) DelegateOwnership(ctx context.Context, id uint, newOwnerId uint) error {
	// check restaurant existence
	var existingRestaurant entities.Restaurant
	err := r.db.WithContext(ctx).Model(&entities.Restaurant{}).Where("id = ?", id).First(&existingRestaurant).Error
	if err != nil {
		return ErrRestaurantNotFound
	}

	// check restaurant staff existence
	var existingResStaff entities.RestaurantStaff
	err = r.db.WithContext(ctx).Model(&entities.RestaurantStaff{}).
		Where("restaurant_id = ? and position = ?", id, restaurant_staff.Manager).First(&existingResStaff).Error
	if err != nil {
		return ErrRestaurantNotFound
	}

	return r.db.Transaction(func(tx *gorm.DB) error {
		// update restaurant owner
		existingRestaurant.OwnedBy = newOwnerId
		err = tx.WithContext(ctx).Model(&entities.Restaurant{}).Where("id = ?", id).Save(&existingRestaurant).Error
		if err != nil {
			return ErrUpdatingRestaurant
		}

		// soft delete previous
		// create new restaurant staff (owner)
		err = r.staffRepo.Delete(ctx, existingResStaff.ID)
		if err != nil {
			return ErrDeletingRestaurantStaff
		}

		newResStaff := &restaurant_staff.RestaurantStaff{
			UserId:       newOwnerId,
			RestaurantId: id,
			Position:     restaurant_staff.Manager,
		}

		_, err = r.staffRepo.Create(ctx, newResStaff)
		if err != nil {
			return err
		}

		existingResStaff.UserId = newOwnerId
		err = tx.WithContext(ctx).Model(&entities.RestaurantStaff{}).Where("id = ?", existingResStaff.ID).Save(&existingResStaff).Error
		if err != nil {
			return err
		}

		return nil
	})
}

func (r *restaurantRepo) SearchRestaurants(ctx context.Context, searchData *restaurant.RestaurantSearch) ([]*restaurant.Restaurant, error) {
	var restaurantsEntities []*entities.Restaurant
	query := r.db.Table("restaurants").Select("restaurants.*").
		Joins("JOIN categories ON categories.restaurant_id = restaurants.id").
		Where("restaurants.name LIKE ?", "%"+searchData.Name+"%")

	if searchData.CategoryID != nil {
		query = query.Where("categories.id = ?", *searchData.CategoryID)
	}

	userID := ctx.Value(jwt.UserClaimKey).(*jwt.UserClaims).UserID
	blockedRestaurantSubQuery := r.db.Model(&entities.BlockRestaurant{}).
		Select("block_restaurants.restaurant_id").
		Where("block_restaurants.user_id = ?", userID)

	query = query.Not("restaurants.id IN (?)", blockedRestaurantSubQuery)

	if searchData.Lat != nil && searchData.Lng != nil {
		query = query.Where("ST_DistanceSphere(ST_MakePoint(restaurants.lng, restaurants.lat), ST_MakePoint(?, ?)) < ?", *searchData.Lng, *searchData.Lat, 5000)
	}

	err := query.Find(&restaurantsEntities).Error

	if err != nil {
		return nil, ErrFetchingRestaurants
	}

	var restaurants []*restaurant.Restaurant
	for _, res := range restaurantsEntities {
		model := mappers.RestaurantEntityToDomain(res)
		restaurants = append(restaurants, model)
	}

	return restaurants, nil
}
