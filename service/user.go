package service

import (
	"TaamResan/internal/user"
	"context"
	"errors"
)

type UserService struct {
	userOps *user.Ops
}

func NewUserService(userOps *user.Ops) *UserService {
	return &UserService{
		userOps: userOps,
	}
}

var (
	ErrUserExists            = errors.New("user already exists")
	ErrCreatingUser          = errors.New("can not create user")
	ErrNoActiveWalletForUser = errors.New("there is no active wallet for this user")
)

func (s *UserService) CreateUser(ctx context.Context, user *user.User) error {
	_, err := s.userOps.FindUserByMobile(ctx, user.Mobile)
	if err == nil {
		return ErrUserExists
	}

	if user.Email != "" {
		_, err = s.userOps.FindUserByEmail(ctx, user.Mobile)
		if err == nil {
			return ErrUserExists
		}
	}

	err = s.userOps.Create(ctx, user)
	if err != nil {
		return ErrCreatingUser
	}

	return nil
}

func (s *UserService) UpdateUserProfile(ctx context.Context, u *user.User) error {
	foundUserByMobile, err := s.userOps.FindUserByMobile(ctx, u.Mobile)
	if err == nil && foundUserByMobile.ID != u.ID {
		return ErrUserExists
	}

	if u.Email != "" {
		foundUserByEmail, err := s.userOps.FindUserByEmail(ctx, u.Email)
		if err == nil && foundUserByEmail.ID != u.ID {
			return ErrUserExists
		}
	}

	err = s.userOps.Update(ctx, u)
	if err != nil {
		return ErrCreatingUser
	}

	return nil
}

//func (s *UserService) GetUserActiveWallet(ctx context.Context, userId uint) (*wallet.Wallet, error) {
//	w, err := s.userOps.GetUserActiveWallet(ctx, userId)
//	if err == nil {
//		return nil, ErrNoActiveWalletForUser
//	}
//
//	return w, nil
//}

func (s *UserService) CreateAdmin(ctx context.Context) error {
	return s.userOps.CreateAdmin(ctx)
}

func (s *UserService) GetUserRoles(ctx context.Context, userId uint) ([]uint, error) {
	return s.userOps.GetUserRoles(ctx, userId)
}
