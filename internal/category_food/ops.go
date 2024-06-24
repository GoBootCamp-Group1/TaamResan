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

func (o *Ops) Delete(ctx context.Context, id uint) error {
	return o.repo.Delete(ctx, id)
}

func (o *Ops) GetById(ctx context.Context, id uint) (*CategoryFood, error) {
	return o.repo.GetById(ctx, id)
}

func (o *Ops) GetAllByFoodId(ctx context.Context, foodId uint) ([]*CategoryFood, error) {
	return o.repo.GetAllByFoodId(ctx, foodId)
}

func (o *Ops) GetAllByCategoryId(ctx context.Context, categoryId uint) ([]*CategoryFood, error) {
	return o.repo.GetAllByCategoryId(ctx, categoryId)
}
