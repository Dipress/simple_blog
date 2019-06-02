package user

import (
	"errors"
	"time"
)

// easyjson -all model.go

var (
	// ErrNotFound raises when user not found in the database.
	ErrNotFound = errors.New("user not found")
)

// User contains all user field.
type User struct {
	ID           int       `json:"id"`
	Username     string    `json:"username"`
	Email        string    `json:"email"`
	PasswordHash string    `json:"password_hash"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

// NewUser contains the information which needs to create a new User.
type NewUser struct {
	Username     string
	Email        string
	PasswordHash string
}
