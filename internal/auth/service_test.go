package auth

import (
	"context"
	"testing"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	user "github.com/dipress/blog/internal/user"
	"github.com/golang/mock/gomock"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
	"golang.org/x/crypto/bcrypt"
)

func Test_Service(t *testing.T) {
	pw, err := bcrypt.GenerateFromPassword([]byte("password123"), bcrypt.DefaultCost)
	if err != nil {
		t.Fatalf("failed to generate password: %v", err)
	}

	tests := []struct {
		name               string
		repositoryFunc     func(ctx context.Context, email string, user *user.User) error
		tokenGeneratorFunc func(ctx context.Context, claims jwt.Claims) (string, error)
		wantErr            bool
		expect             Token
	}{
		{
			name: "ok",
			repositoryFunc: func(ctx context.Context, email string, user *user.User) error {
				user.PasswordHash = string(pw)
				return nil
			},
			tokenGeneratorFunc: func(ctx context.Context, claims jwt.Claims) (string, error) {
				return "token", nil
			},
			expect: Token{
				Token: "token",
			},
		},
		{
			name: "not found",
			repositoryFunc: func(ctx context.Context, email string, user *user.User) error {
				return ErrNotFound
			},
			wantErr: true,
		},
		{
			name: "wrong password",
			repositoryFunc: func(ctx context.Context, email string, user *user.User) error {
				return nil
			},
			wantErr: true,
		},
		{
			name: "token gen",
			repositoryFunc: func(ctx context.Context, email string, user *user.User) error {
				user.PasswordHash = string(pw)
				return nil
			},
			tokenGeneratorFunc: func(ctx context.Context, claims jwt.Claims) (string, error) {
				return "", errors.New("mock error")
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			s := NewService(repositoryFunc(tt.repositoryFunc), tokenGeneratorFunc(tt.tokenGeneratorFunc), time.Hour)

			ctx, cancel := context.WithCancel(context.Background())
			defer cancel()

			var got Token
			email := "username@example.com"
			password := "password123"
			err := s.Authenticate(ctx, email, password, &got)

			if tt.wantErr {
				assert.Error(t, err)
				return
			}
			assert.Nil(t, err)
			assert.Equal(t, got, tt.expect)
		})
	}
}

type repositoryFunc func(ctx context.Context, email string, user *user.User) error

func (r repositoryFunc) FindByEmail(ctx context.Context, email string, user *user.User) error {
	return r(ctx, email, user)
}

type tokenGeneratorFunc func(ctx context.Context, claims jwt.Claims) (string, error)

func (t tokenGeneratorFunc) GenerateToken(ctx context.Context, claims jwt.Claims) (string, error) {
	return t(ctx, claims)
}
