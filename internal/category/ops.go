package category

import "context"

type Ops struct {
	repo Repo
}

func NewOps(repo Repo) *Ops {
	return &Ops{repo: repo}
}

func (o *Ops) Create(ctx context.Context, category *Category) (uint, error) {
	return o.repo.Create(ctx, category)
}

func (o *Ops) Update(ctx context.Context, category *Category) error {
	return o.repo.Update(ctx, category)
}

func (o *Ops) Delete(ctx context.Context, id uint) error {
	return o.repo.Delete(ctx, id)
}

func (o *Ops) GetById(ctx context.Context, id uint) (*Category, error) {
	return o.repo.GetById(ctx, id)
}

func (o *Ops) GetByName(ctx context.Context, name string, restaurantId uint) (*Category, error) {
	return o.repo.GetByName(ctx, name, restaurantId)
}

func (o *Ops) GetAll(ctx context.Context, restaurantId uint) ([]*Category, error) {
	return o.repo.GetAll(ctx, restaurantId)
}
