package cart

import "context"

type Ops struct {
	repo Repo
}

func NewOps(repo Repo) *Ops {
	return &Ops{repo: repo}
}

func (o *Ops) Delete(ctx context.Context, id uint) error {
	return o.repo.Delete(ctx, id)
}

func (o *Ops) GetByUserId(ctx context.Context, userId uint) (*Cart, error) {
	return o.repo.GetByUserId(ctx, userId)
}

func (o *Ops) GetById(ctx context.Context, cartId uint) (*Cart, error) {
	return o.repo.GetById(ctx, cartId)
}

func (o *Ops) GetItemsFeeByID(ctx context.Context, cartId uint) (float64, error) {
	return o.repo.GetItemsFeeByID(ctx, cartId)
}
