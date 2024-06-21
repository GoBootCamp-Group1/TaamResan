package address

import (
	"context"
)

type Ops struct {
	repo Repo
}

func NewOps(repo Repo) *Ops {
	return &Ops{
		repo: repo,
	}
}

func (o *Ops) Create(ctx context.Context, address *Address) error {
	return o.repo.Create(ctx, address)
}

func (o *Ops) Update(ctx context.Context, address *Address) error {
	return o.repo.Update(ctx, address)
}

func (o *Ops) Delete(ctx context.Context, address *Address) error {
	return o.repo.Delete(ctx, address)
}

func (o *Ops) GetAddressByID(ctx context.Context, id uint) (*Address, error) {
	return o.repo.GetByID(ctx, id)
}
func (o *Ops) GetAll(ctx context.Context) ([]*Address, error) {
	return o.repo.GetAll(ctx)
}
