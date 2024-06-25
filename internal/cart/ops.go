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