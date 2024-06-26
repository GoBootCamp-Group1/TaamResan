package access

import "context"

type Repo interface {
	CheckRestaurantOwner(ctx context.Context, userId uint, restaurantId uint) error
	CheckRestaurantStaff(ctx context.Context, userId uint, restaurantId uint) error
	CheckAdminAccess(ctx context.Context, userId uint) error
}
