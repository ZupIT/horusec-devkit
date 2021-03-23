package entities

import (
	"testing"

	"github.com/dgrijalva/jwt-go"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestValidate(t *testing.T) {
	t.Run("should return no error when valid claim", func(t *testing.T) {
		claims := &JWTClaims{
			Email:    "test@test.com",
			Username: "test",
			StandardClaims: jwt.StandardClaims{
				Subject: uuid.New().String(),
			},
		}

		assert.NoError(t, claims.Validate())
	})

	t.Run("should return error when invalid", func(t *testing.T) {
		claims := &JWTClaims{}

		assert.Error(t, claims.Validate())
	})
}
