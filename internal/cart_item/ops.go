package cart_item

import "context"

type Ops struct {
	repo Repo
}

func NewOps(repo Repo) *Ops {
	return &Ops{repo: repo}
}

func (o *Ops) Create(ctx context.Context, cartItem *CartItem) (uint, error) {
	return o.repo.Create(ctx, cartItem)
}

func (o *Ops) Update(ctx context.Context, cartItem *CartItem) error {
	return o.repo.Update(ctx, cartItem)
}

func (o *Ops) Delete(ctx context.Context, id uint) error {
	return o.repo.Delete(ctx, id)
}

func (o *Ops) GetAllByCartId(ctx context.Context, cartId uint) ([]*CartItem, error) {
	return o.repo.GetAllByCartId(ctx, cartId)
}
