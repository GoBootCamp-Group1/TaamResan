package role

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

func (o *Ops) Create(ctx context.Context, role *Role) error {
	return o.repo.Create(ctx, role)
}

func (o *Ops) Update(ctx context.Context, role *Role) error {
	return o.repo.Update(ctx, role)
}

func (o *Ops) Delete(ctx context.Context, id uint) error {
	return o.repo.Delete(ctx, id)
}

func (o *Ops) Get(ctx context.Context, id uint) (*Role, error) {
	return o.repo.Get(ctx, id)
}

func (o *Ops) GetAll(ctx context.Context) ([]*Role, error) {
	return o.repo.GetAll(ctx)
}