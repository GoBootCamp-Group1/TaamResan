package food

import "context"

type Ops struct {
	repo Repo
}

func NewOps(repo Repo) *Ops {
	return &Ops{repo: repo}
}

func (o *Ops) Create(ctx context.Context, food *Food) (uint, error) {
	return o.repo.Create(ctx, food)
}

func (o *Ops) Update(ctx context.Context, food *Food) error {
	return o.repo.Update(ctx, food)
}

func (o *Ops) Delete(ctx context.Context, id uint) error {
	return o.repo.Delete(ctx, id)
}

func (o *Ops) GetById(ctx context.Context, id uint) (*Food, error) {
	return o.repo.GetById(ctx, id)
}

func (o *Ops) GetAll(ctx context.Context, restaurantId uint) ([]*Food, error) {
	return o.repo.GetAll(ctx, restaurantId)
}

func (o *Ops) SearchFoods(ctx context.Context, searchData *FoodSearch) ([]*Food, error) {
	return o.repo.SearchFoods(ctx, searchData)
}
