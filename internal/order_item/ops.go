package order

type Ops struct {
	repo Repo
}

func NewOps(repo Repo) *Ops { return &Ops{repo: repo} }

//func (o Ops) Create(ctx context.Context, cartModel *cart.Cart) (*Order, error) {
//	return o.repo.Create(ctx, cartModel)
//}
