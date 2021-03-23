package entities

import (
	"github.com/dgrijalva/jwt-go"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
	"github.com/google/uuid"
)

type JWTClaims struct {
	Email       string   `json:"email"`
	Username    string   `json:"username"`
	Permissions []string `json:"permissions"`
	jwt.StandardClaims
}

func (j *JWTClaims) Validate() error {
	return validation.ValidateStruct(j,
		validation.Field(&j.Username, validation.Required, validation.Length(1, 255)),
		validation.Field(&j.Email, validation.Required, validation.Length(1, 255)),
		validation.Field(&j.Subject, validation.Required, is.UUID, validation.NotIn(uuid.Nil)),
	)
}
