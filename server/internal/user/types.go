package user

import (
	"context"
	"crypto/sha256"
	"errors"
	"fmt"
)

type Repo interface {
	Create(ctx context.Context, user *User) error
	GetByID(ctx context.Context, id uint) (*User, error)
	GetByEmail(ctx context.Context, email string) (*User, error)
}

var (
	ErrUserNotFound    = errors.New("user not found")
	ErrInvalidPassword = errors.New("invalid user password")
)

type Role uint8

func (ur Role) String() string {
	switch ur {
	case RoleUser:
		return "user"
	case RoleAdmin:
		return "admin"
	default:
		return "unknown"
	}
}

const (
	RoleUser Role = iota + 1
	RoleAdmin
	RoleManager
	RoleOperator
)

type User struct {
	ID        uint
	FirstName string
	LastName  string
	Email     string
	Password  string
	Role      Role
}

func (u *User) PasswordIsValid(pass string) bool {
	h := sha256.New()
	h.Write([]byte(pass))
	passSha256 := h.Sum(nil)
	return fmt.Sprintf("%x", passSha256) == u.Password
}
