package service

import (
	"TaamResan/internal/role"
	"TaamResan/internal/user"
	"TaamResan/pkg/jwt"
	"context"
	"errors"
	"time"

	jwt2 "github.com/golang-jwt/jwt/v5"
)

type AuthService struct {
	userOps                *user.Ops
	secret                 []byte
	tokenExpiration        uint
	refreshTokenExpiration uint
}

func NewAuthService(userOps *user.Ops, secret []byte,
	tokenExpiration uint, refreshTokenExpiration uint) *AuthService {
	return &AuthService{
		userOps:                userOps,
		secret:                 secret,
		tokenExpiration:        tokenExpiration,
		refreshTokenExpiration: refreshTokenExpiration,
	}
}

type UserToken struct {
	AuthorizationToken string
	RefreshToken       string
	ExpiresAt          int64
}

var (
	ErrCreatingAuthToken    = errors.New("can not create authentication token")
	ErrCreatingRefreshToken = errors.New("can not create refresh token")
)

func (s *AuthService) LoginWithMobile(ctx context.Context, mobile, pass string) (*UserToken, error) {
	user, err := s.userOps.GetUserByMobileAndPassword(ctx, mobile, pass)
	if err != nil {
		return nil, err
	}

	authExp, refreshExp := s.calculateTokenExpirationTime()

	authToken, err := jwt.CreateToken(s.secret, s.userClaims(user, authExp))
	if err != nil {
		return nil, ErrCreatingAuthToken
	}

	refreshToken, err := jwt.CreateToken(s.secret, s.userClaims(user, refreshExp))
	if err != nil {
		return nil, ErrCreatingRefreshToken
	}

	return &UserToken{
		AuthorizationToken: authToken,
		RefreshToken:       refreshToken,
		ExpiresAt:          authExp.Unix(),
	}, nil
}

func (s *AuthService) LoginWithEmail(ctx context.Context, email, pass string) (*UserToken, error) {
	user, err := s.userOps.GetUserByEmailAndPassword(ctx, email, pass)
	if err != nil {
		return nil, err
	}

	authExp, refreshExp := s.calculateTokenExpirationTime()

	authToken, err := jwt.CreateToken(s.secret, s.userClaims(user, authExp))
	if err != nil {
		return nil, ErrCreatingAuthToken
	}

	refreshToken, err := jwt.CreateToken(s.secret, s.userClaims(user, refreshExp))
	if err != nil {
		return nil, ErrCreatingRefreshToken
	}

	return &UserToken{
		AuthorizationToken: authToken,
		RefreshToken:       refreshToken,
		ExpiresAt:          authExp.Unix(),
	}, nil
}

func (s *AuthService) calculateTokenExpirationTime() (time.Time, time.Time) {
	// calc expiration time values
	var (
		authExp    = time.Now().Add(time.Minute * time.Duration(s.tokenExpiration))
		refreshExp = time.Now().Add(time.Minute * time.Duration(s.refreshTokenExpiration))
	)
	return authExp, refreshExp
}

func (s *AuthService) userClaims(user *user.User, exp time.Time) *jwt.UserClaims {
	return &jwt.UserClaims{
		RegisteredClaims: jwt2.RegisteredClaims{
			ExpiresAt: &jwt2.NumericDate{
				Time: exp,
			},
		},
		UserID: user.ID,
		Roles:  getRoleToString(user.Roles),
	}
}

func getRoleToString(roles []role.Role) []string {
	var rolesStr []string
	for _, r := range roles {
		rolesStr = append(rolesStr, r.String())
	}
	return rolesStr
}
