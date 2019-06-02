package reg

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
//go:generate mockgen -source=service.go -package=reg -destination=service.mock.go

var (
	// ErrUsernameExists returns when given username is already
	// present in database.
	ErrUsernameExists = errors.New("username already exists")
	// ErrEmailExists returns when given email is already
	// present in database.
	ErrEmailExists = errors.New("email already exists")
)

// Validater validates user fields.
type Validater interface {
	ValidateUser(context.Context, *Form) error
}

// Repository allows to work with database.
type Repository interface {
	CreateUser(ctx context.Context, f *user.NewUser, usr *user.User) (func() error, func() error, error)
	UniqueUsername(ctx context.Context, username string) error
	UniqueEmail(ctx context.Context, email string) error
}

// TokenGenerator is the behavior we need in our
// Registrate to generate tokens for registrate users.
type TokenGenerator interface {
	GenerateToken(ctx context.Context, claims jwt.Claims) (string, error)
}

// Form is a user form.
//easyjson:json
type Form struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

// Token holds token data.
//easyjson:json
type Token struct {
	Token string `json:"token"`
}

// Service holds everything required to registrate user.
type Service struct {
	Repository
	Validater
	TokenGenerator
	ExpireAfter time.Duration
}

// NewService factory prepares service for
// futher operations.
func NewService(r Repository, v Validater, tg TokenGenerator, exp time.Duration) *Service {
	s := Service{
		Repository:     r,
		Validater:      v,
		ExpireAfter:    exp,
		TokenGenerator: tg,
	}
	return &s
}

// Registrate registrates user.
func (s *Service) Registrate(ctx context.Context, f *Form, token *Token) error {
	if err := s.Validater.ValidateUser(ctx, f); err != nil {
		return errors.Wrap(err, "validater validate")
	}

	if err := s.Repository.UniqueUsername(ctx, f.Username); err != nil {
		return errors.Wrap(err, "unique username")
	}

	if err := s.Repository.UniqueEmail(ctx, f.Email); err != nil {
		return errors.Wrap(err, "uniqie email")
	}

	pw, err := bcrypt.GenerateFromPassword([]byte(f.Password), bcrypt.DefaultCost)
	if err != nil {
		return errors.Wrap(err, "generating password hash")
	}

	nu := user.NewUser{
		Email:        f.Email,
		PasswordHash: string(pw),
	}

	var user user.User
	commit, rollback, err := s.Repository.CreateUser(ctx, &nu, &user)
	if err != nil {
		return errors.Wrap(err, "create user")
	}

	//TODO: create email notifier service.

	claims := auth.NewClaims(user.Username, time.Now(), s.ExpireAfter)

	tknStr, err := s.TokenGenerator.GenerateToken(ctx, claims)
	if err != nil {
		rollback()
		return errors.Wrap(err, "generate token")
	}

	token.Token = tknStr

	return commit()
}
