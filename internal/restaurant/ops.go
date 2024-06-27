package restaurant

import (
	"context"
)

type Ops struct {
	repo Repo
}

func NewOps(repo Repo) *Ops {
	return &Ops{repo: repo}
}

func (o *Ops) Create(ctx context.Context, restaurant *Restaurant) (uint, error) {
	return o.repo.Create(ctx, restaurant)
}

func (o *Ops) Update(ctx context.Context, restaurant *Restaurant) error {
	return o.repo.Update(ctx, restaurant)
}

func (o *Ops) Delete(ctx context.Context, id uint) error {
	return o.repo.Delete(ctx, id)
}

func (o *Ops) GetById(ctx context.Context, id uint) (*Restaurant, error) {
	return o.repo.GetById(ctx, id)
}

func (o *Ops) GetAll(ctx context.Context) ([]*Restaurant, error) {
	return o.repo.GetAll(ctx)
}

func (o *Ops) Approve(ctx context.Context, id uint) error {
	return o.repo.Approve(ctx, id)
}

func (o *Ops) DelegateOwnership(ctx context.Context, id uint, newOwnerId uint) error {
	return o.repo.DelegateOwnership(ctx, id, newOwnerId)
}

func (o *Ops) SearchRestaurants(ctx context.Context, searchData *RestaurantSearch) ([]*Restaurant, error) {
	return o.repo.SearchRestaurants(ctx, searchData)
}
