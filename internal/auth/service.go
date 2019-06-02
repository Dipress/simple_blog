package auth

import (
	"context"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/dipress/blog/internal/user"
	"github.com/dipress/blog/kit/auth"
	"github.com/pkg/errors"
	"golang.org/x/crypto/bcrypt"
)

// easyjson service.go

var (
	// ErrNotFound returns when given email is not
	// found in database.
	ErrNotFound = errors.New("email not found")
	// ErrWrongPassword returns when given password
	// is not equal to it's hash in database.
	ErrWrongPassword = errors.New("wrong password")
)

// Repository allows to work with database.
type Repository interface {
	FindByEmail(ctx context.Context, email string, user *user.User) error
}

// TokenGenerator is the behavior we need in our
// Authenticate to generate tokens for authenticated users.
type TokenGenerator interface {
	GenerateToken(ctx context.Context, claims jwt.Claims) (string, error)
}

// Service holds required data for user
// authentication.
type Service struct {
	Repository
	TokenGenerator
	ExpireAfter time.Duration
}

// Form is a user auth form.
//easyjson:json
type Form struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

// Token holds token data.
//easyjson:json
type Token struct {
	Token string `json:"token"`
}

// NewService factory created ready to service.
func NewService(r Repository, t TokenGenerator, exp time.Duration) *Service {
	s := Service{
		Repository:     r,
		TokenGenerator: t,
		ExpireAfter:    exp,
	}

	return &s
}

// Authenticate allows authenticating user by given email and password
// and set t Token value as generated token.
func (s *Service) Authenticate(ctx context.Context, email, password string, t *Token) error {
	var user user.User
	if err := s.Repository.FindByEmail(ctx, email, &user); err != nil {
		return errors.Wrap(err, "find user by email")
	}

	// Compare the provided password with the saved hash. Use the bcrypt
	// comparison function so it is cryptographically secure.
	if err := bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(password)); err != nil {
		return ErrWrongPassword
	}

	// If we are this far the request is valid. Create some claims for the user
	// and generate their token.
	claims := auth.NewClaims(user.Username, time.Now(), s.ExpireAfter)

	tknStr, err := s.GenerateToken(ctx, claims)
	if err != nil {
		return errors.Wrap(err, "generate token")
	}
	t.Token = tknStr

	return nil
}
