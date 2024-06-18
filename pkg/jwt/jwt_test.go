package jwt

import (
	"github.com/golang-jwt/jwt/v5"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestCreateAndParseToken(t *testing.T) {
	secret := []byte("secret")
	claims := &UserClaims{
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    "test",
			Subject:   "subject",
			Audience:  []string{"audience"},
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
		UserID:   1,
		Roles:    []string{"admin"},
		Sections: []string{"section1", "section2"},
	}

	tokenString, err := CreateToken(secret, claims)
	assert.NoError(t, err)
	assert.NotEmpty(t, tokenString)

	parsedClaims, err := ParseToken(tokenString, secret)
	assert.NoError(t, err)
	assert.NotNil(t, parsedClaims)

	assert.Equal(t, claims.UserID, parsedClaims.UserID)
	assert.Equal(t, claims.Roles, parsedClaims.Roles)
	assert.ElementsMatch(t, claims.Sections, parsedClaims.Sections)
	assert.Equal(t, claims.RegisteredClaims.Issuer, parsedClaims.RegisteredClaims.Issuer)
	assert.Equal(t, claims.RegisteredClaims.Subject, parsedClaims.RegisteredClaims.Subject)
	assert.Equal(t, claims.RegisteredClaims.Audience, parsedClaims.RegisteredClaims.Audience)

	if claims.RegisteredClaims.ExpiresAt != nil && parsedClaims.RegisteredClaims.ExpiresAt != nil {
		assert.True(t, claims.RegisteredClaims.ExpiresAt.Time.Equal(parsedClaims.RegisteredClaims.ExpiresAt.Time))
	}
	if claims.RegisteredClaims.IssuedAt != nil && parsedClaims.RegisteredClaims.IssuedAt != nil {
		assert.True(t, claims.RegisteredClaims.IssuedAt.Time.Equal(parsedClaims.RegisteredClaims.IssuedAt.Time))
	}
}
