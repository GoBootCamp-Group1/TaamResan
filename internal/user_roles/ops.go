package user_roles

import "context"

type Ops struct {
	repo Repo
}

func NewOps(repo Repo) *Ops {
	return &Ops{
		repo: repo,
	}
}

func (o *Ops) Create(ctx context.Context, ur *UserRoles) error {
	return o.repo.Create(ctx, ur)
}

func (o *Ops) Update(ctx context.Context, ur *UserRoles) error {
	return o.repo.Update(ctx, ur)
}

func (o *Ops) Delete(ctx context.Context, ur *UserRoles) error {
	return o.repo.Delete(ctx, ur)
}

func (o *Ops) Get(ctx context.Context, id uint) (*UserRoles, error) {
	return o.repo.Get(ctx, id)
}
