package user

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

func (o *Ops) FindUserByMobile(ctx context.Context, mobile string) (*User, error) {
	user, err := o.repo.GetByMobile(ctx, mobile)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (o *Ops) FindUserByEmail(ctx context.Context, email string) (*User, error) {
	user, err := o.repo.GetByEmail(ctx, email)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (o *Ops) Create(ctx context.Context, user *User) error {
	return o.repo.Create(ctx, user)
}

func (o *Ops) GetUserByID(ctx context.Context, id uint) (*User, error) {
	return o.repo.GetByID(ctx, id)
}

func (o *Ops) GetUserByMobileAndPassword(ctx context.Context, mobile, password string) (*User, error) {
	user, err := o.repo.GetByMobile(ctx, mobile)
	if err != nil {
		return nil, err
	}

	if user == nil {
		return nil, ErrUserNotFound
	}

	if !user.PasswordIsValid(password) {
		return nil, ErrInvalidPassword
	}

	return user, nil
}

func (o *Ops) GetUserByEmailAndPassword(ctx context.Context, email, password string) (*User, error) {
	user, err := o.repo.GetByEmail(ctx, email)
	if err != nil {
		return nil, err
	}

	if user == nil {
		return nil, ErrUserNotFound
	}

	if !user.PasswordIsValid(password) {
		return nil, ErrInvalidPassword
	}

	return user, nil
}
