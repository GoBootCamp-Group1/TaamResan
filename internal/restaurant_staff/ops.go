package restaurant_staff

import "context"

type Ops struct {
	repo Repo
}

func NewOps(repo Repo) *Ops {
	return &Ops{repo: repo}
}

func (o *Ops) Create(ctx context.Context, rStaff *RestaurantStaff) (uint, error) {
	return o.repo.Create(ctx, rStaff)
}
