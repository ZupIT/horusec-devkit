package entities

import "github.com/google/uuid"

type TokenData struct {
	Email     string    `json:"email"`
	Username  string    `json:"username"`
	AccountID uuid.UUID `json:"accountID"`
}
