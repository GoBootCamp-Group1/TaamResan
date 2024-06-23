package category_food

import "context"

type Ops struct {
	repo Repo
}

func NewOps(repo Repo) *Ops {
	return &Ops{repo: repo}
}

func (o *Ops) Create(ctx context.Context, categoryFood *CategoryFood) (uint, error) {
	return o.repo.Create(ctx, categoryFood)
}

func (o *Ops) Update(ctx context.Context, categoryFood *CategoryFood) error {
	return o.repo.Update(ctx, categoryFood)
}

func (o *Ops) Delete(ctx context.Context, id uint) error {
	return o.repo.Delete(ctx, id)
}

func (o *Ops) GetById(ctx context.Context, id uint) (*CategoryFood, error) {
	return o.repo.GetById(ctx, id)
}

func (o *Ops) GetAll(ctx context.Context, restaurantId uint) ([]*CategoryFood, error) {
	return o.repo.GetAll(ctx, restaurantId)
}
