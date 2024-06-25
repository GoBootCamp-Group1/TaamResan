package block_restaurant

import "context"

type Ops struct {
	repo Repo
}

func NewOps(repo Repo) *Ops {
	return &Ops{repo: repo}
}

func (o *Ops) Create(ctx context.Context, br *BlockRestaurant) (uint, error) {
	return o.repo.Create(ctx, br)
}

func (o *Ops) Delete(ctx context.Context, id uint) error {
	return o.repo.Delete(ctx, id)
}

func (o *Ops) GetAllByUserId(ctx context.Context, userId uint) ([]*BlockRestaurant, error) {
	return o.repo.GetAllByUserId(ctx, userId)
}
