package service

import (
	"TaamResan/internal/user"
	"context"
)

type UserService struct {
	userOps *user.Ops
}

func NewUserService(userOps *user.Ops) *UserService {
	return &UserService{
		userOps: userOps,
	}
}

func (s *UserService) CreateUser(ctx context.Context, user *user.User) error {
	return s.userOps.Create(ctx, user)
}
