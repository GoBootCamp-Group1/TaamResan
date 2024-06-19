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
	ErrUserExists   = errors.New("user already exists")
	ErrCreatingUser = errors.New("can not create user")
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
