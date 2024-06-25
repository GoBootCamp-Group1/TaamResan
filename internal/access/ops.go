package access

import "context"

type Ops struct {
	repo Repo
}

func NewOps(repo Repo) *Ops {
	return &Ops{
		repo: repo,
	}
}

func (o *Ops) CheckRestaurantOwner(ctx context.Context, userId uint, restaurantId uint) error {
	return o.repo.CheckRestaurantOwner(ctx, userId, restaurantId)
}

func (o *Ops) CheckAdminAccess(ctx context.Context, userId uint) error {
	return o.repo.CheckAdminAccess(ctx, userId)
}
