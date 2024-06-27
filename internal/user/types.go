package user

import (
	"TaamResan/internal/role"
	"context"
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"fmt"
	"github.com/google/uuid"
	"time"
)

type Repo interface {
	Create(ctx context.Context, user *User) error
	Update(ctx context.Context, user *User) error
	GetByMobile(ctx context.Context, mobile string) (*User, error)
	GetByEmail(ctx context.Context, email string) (*User, error)
	//GetUserActiveWallet(ctx context.Context, userId uint) (*wallet.Wallet, error)
	CreateAdmin(ctx context.Context) error
}

var (
	ErrUserNotFound    = errors.New("user not found")
	ErrInvalidPassword = errors.New("invalid user password")
)

type User struct {
	ID        uint
	Uuid      string
	Name      string
	Email     string
	Mobile    string
	BirthDate time.Time
	Password  string
	Roles     []role.Role
}

func HashPassword(password string) string {
	h := sha256.New()
	h.Write([]byte(password))
	return hex.EncodeToString(h.Sum(nil))
}

func (u *User) PasswordIsValid(pass string) bool {
	h := sha256.New()
	h.Write([]byte(pass))
	passSha256 := h.Sum(nil)
	return fmt.Sprintf("%x", passSha256) == u.Password
}

var DefaultAdminUser = User{
	Uuid:      uuid.NewString(),
	Name:      "admin",
	Email:     "admin@admin.com",
	Mobile:    "091122233441",
	BirthDate: time.Date(2000, 2, 2, 0, 0, 0, 0, &time.Location{}),
	Password:  "Adminadmin1@",
}
