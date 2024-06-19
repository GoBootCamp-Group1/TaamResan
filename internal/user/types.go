package user

import (
	"context"
	"crypto/sha256"
	"errors"
	"fmt"
	"time"
)

type Repo interface {
	Create(ctx context.Context, user *User) error
	GetByID(ctx context.Context, id uint) (*User, error)
	GetByMobile(ctx context.Context, mobile string) (*User, error)
	GetByEmail(ctx context.Context, email string) (*User, error)
}

var (
	ErrUserNotFound    = errors.New("user not found")
	ErrInvalidPassword = errors.New("invalid user password")
)

type Role uint8

func (ur Role) String() string {
	switch ur {
	case Customer:
		return "customer"
	case Admin:
		return "admin"
	case RestaurantOwner:
		return "restaurant owner"
	case RestaurantOperator:
		return "restaurant operator"
	default:
		return "unknown"
	}
}

const (
	Customer Role = iota + 1
	Admin
	RestaurantOwner
	RestaurantOperator
)

type User struct {
	ID        uint
	Name      string
	Email     string
	Mobile    string
	BirthDate time.Time
	Password  string
	Roles     []Role
}

func (u *User) ValidateMobile(mobile string) error {
	panic("implement me")
}

func (u *User) ValidateEmail(email string) error {
	panic("implement me")
}

func (u *User) ValidatePassword(pass string) error {
	panic("implement me")
}

func (u *User) PasswordIsValid(pass string) bool {
	h := sha256.New()
	h.Write([]byte(pass))
	passSha256 := h.Sum(nil)
	return fmt.Sprintf("%x", passSha256) == u.Password
}
